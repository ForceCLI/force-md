package cmd

import (
	"os"
	"path/filepath"
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
	if err := openTidyArgs(r, []string{dir}, true); err != nil {
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
	if err := openTidyArgs(r, []string{file}, true); err != nil {
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
	if err := openTidyArgs(r, []string{dir}, true); err != nil {
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
	if err := openTidyArgs(r, []string{dir}, false); err == nil {
		t.Error("expected an error for a directory without --recursive")
	}
}
