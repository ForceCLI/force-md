package cmd

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/ForceCLI/force-md/repo"
)

const tidyTestFlow = `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
    <processType>AutoLaunchedFlow</processType>
    <status>Active</status>
</Flow>
`

func writeTidyFixture(t *testing.T, path, contents string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}
}

func openedPaths(r *repo.Repo) []string {
	var paths []string
	for _, typ := range r.Types() {
		for _, m := range r.Items(typ) {
			paths = append(paths, string(m.GetMetadataInfo().Path()))
		}
	}
	sort.Strings(paths)
	return paths
}

func TestOpenTidyArgsRecursiveWalksDirectories(t *testing.T) {
	dir := t.TempDir()
	top := filepath.Join(dir, "flows", "Top.flow-meta.xml")
	nested := filepath.Join(dir, "nested", "flows", "Nested.flow-meta.xml")
	writeTidyFixture(t, top, tidyTestFlow)
	writeTidyFixture(t, nested, tidyTestFlow)
	writeTidyFixture(t, filepath.Join(dir, "README.md"), "# not metadata\n")
	writeTidyFixture(t, filepath.Join(dir, ".git", "flows", "Hidden.flow-meta.xml"), tidyTestFlow)

	r := repo.NewRepo()
	if _, err := openTidyArgs(r, []string{dir}, true); err != nil {
		t.Fatal(err)
	}

	got := openedPaths(r)
	want := []string{top, nested}
	sort.Strings(want)
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Errorf("opened %v, want %v", got, want)
	}
}

func TestOpenTidyArgsRecursiveAcceptsFiles(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "flows", "Only.flow-meta.xml")
	writeTidyFixture(t, file, tidyTestFlow)

	r := repo.NewRepo()
	if _, err := openTidyArgs(r, []string{file}, true); err != nil {
		t.Fatal(err)
	}
	if got := openedPaths(r); len(got) != 1 || got[0] != file {
		t.Errorf("opened %v, want [%s]", got, file)
	}
}

func TestOpenTidyArgsRecursiveWarnsAndContinuesOnInvalidMetadata(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "flows", "Good.flow-meta.xml")
	writeTidyFixture(t, good, tidyTestFlow)
	writeTidyFixture(t, filepath.Join(dir, "flows", "Bad.flow-meta.xml"), `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
    <processType>AutoLaunchedFlow</status>
</Flow>
`)

	r := repo.NewRepo()
	if _, err := openTidyArgs(r, []string{dir}, true); err != nil {
		t.Fatal(err)
	}
	if got := openedPaths(r); len(got) != 1 || got[0] != good {
		t.Errorf("opened %v, want [%s]", got, good)
	}
}

func TestOpenTidyArgsWithoutRecursiveRejectsDirectories(t *testing.T) {
	dir := t.TempDir()
	writeTidyFixture(t, filepath.Join(dir, "flows", "Top.flow-meta.xml"), tidyTestFlow)

	r := repo.NewRepo()
	if _, err := openTidyArgs(r, []string{dir}, false); err == nil {
		t.Error("expected an error for a directory without --recursive")
	}
}

const unformattedApexClass = `public class Example { public Integer twice(Integer n) { return n*2; } }`

const formattedApexClass = `public class Example {
	public Integer twice(Integer n) {
		return n * 2;
	}
}
`

const unformattedApexTrigger = `trigger AccountTrigger on Account (before insert) { for (Account a : Trigger.new) { a.Name = a.Name.trim(); } }`

const formattedApexTrigger = `trigger AccountTrigger on Account (before insert) {
	for (Account a : Trigger.new) {
		a.Name = a.Name.trim();
	}
}
`

const apexClassMeta = `<?xml version="1.0" encoding="UTF-8"?>
<ApexClass xmlns="http://soap.sforce.com/2006/04/metadata">
    <apiVersion>60.0</apiVersion>
    <status>Active</status>
</ApexClass>
`

func readTidyFixture(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

// A recursive walk returns the Apex class and trigger sources below a
// directory, and those named as arguments, but not their -meta.xml files or
// the files in hidden directories.
func TestOpenTidyArgsRecursiveFindsApexSources(t *testing.T) {
	dir := t.TempDir()
	class := filepath.Join(dir, "classes", "Example.cls")
	trigger := filepath.Join(dir, "triggers", "AccountTrigger.trigger")
	named := filepath.Join(t.TempDir(), "Named.cls")
	writeTidyFixture(t, class, unformattedApexClass)
	writeTidyFixture(t, class+"-meta.xml", apexClassMeta)
	writeTidyFixture(t, trigger, unformattedApexTrigger)
	writeTidyFixture(t, filepath.Join(dir, ".sfdx", "classes", "Hidden.cls"), unformattedApexClass)
	writeTidyFixture(t, named, unformattedApexClass)

	r := repo.NewRepo()
	got, err := openTidyArgs(r, []string{dir, named}, true)
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(got)
	want := []string{class, trigger, named}
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("found Apex sources %v, want %v", got, want)
	}
	if opened := openedPaths(r); len(opened) != 1 || opened[0] != class+"-meta.xml" {
		t.Errorf("opened %v, want [%s]", opened, class+"-meta.xml")
	}
}

// Without recursive set, an Apex source named as an argument is opened as
// metadata, which fails, as it did before Apex formatting was added.
func TestOpenTidyArgsWithoutRecursiveRejectsApexSources(t *testing.T) {
	class := filepath.Join(t.TempDir(), "Example.cls")
	writeTidyFixture(t, class, unformattedApexClass)

	if _, err := openTidyArgs(repo.NewRepo(), []string{class}, false); err == nil {
		t.Error("expected an error for an Apex source without --recursive")
	}
}

func TestFormatApexWritesFormattedSources(t *testing.T) {
	dir := t.TempDir()
	class := filepath.Join(dir, "Example.cls")
	trigger := filepath.Join(dir, "AccountTrigger.trigger")
	writeTidyFixture(t, class, unformattedApexClass)
	writeTidyFixture(t, trigger, unformattedApexTrigger)

	formatApex([]string{class, trigger})

	if got := readTidyFixture(t, class); got != formattedApexClass {
		t.Errorf("formatted class:\n%s\nwant:\n%s", got, formattedApexClass)
	}
	if got := readTidyFixture(t, trigger); got != formattedApexTrigger {
		t.Errorf("formatted trigger:\n%s\nwant:\n%s", got, formattedApexTrigger)
	}
}

// A source with a syntax error is left unchanged, and does not stop the
// other sources from being formatted.
func TestFormatApexLeavesInvalidSourcesUnchanged(t *testing.T) {
	dir := t.TempDir()
	invalid := filepath.Join(dir, "Invalid.cls")
	valid := filepath.Join(dir, "Example.cls")
	const invalidSource = "public class Invalid { public void run() { "
	writeTidyFixture(t, invalid, invalidSource)
	writeTidyFixture(t, valid, unformattedApexClass)

	formatApex([]string{invalid, valid})

	if got := readTidyFixture(t, invalid); got != invalidSource {
		t.Errorf("invalid source changed to:\n%s", got)
	}
	if got := readTidyFixture(t, valid); got != formattedApexClass {
		t.Errorf("formatted class:\n%s\nwant:\n%s", got, formattedApexClass)
	}
}

func TestListUnformattedApexListsWithoutWriting(t *testing.T) {
	dir := t.TempDir()
	unformatted := filepath.Join(dir, "Example.cls")
	formatted := filepath.Join(dir, "Formatted.cls")
	writeTidyFixture(t, unformatted, unformattedApexClass)
	writeTidyFixture(t, formatted, formattedApexClass)

	got := listUnformattedApex([]string{formatted, unformatted})

	if !reflect.DeepEqual(got, []string{unformatted}) {
		t.Errorf("listed %v, want [%s]", got, unformatted)
	}
	if contents := readTidyFixture(t, unformatted); contents != unformattedApexClass {
		t.Errorf("listing changed the file to:\n%s", contents)
	}
}
