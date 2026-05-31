package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	// Register the CustomPermission type so MetadataFromPath can parse the
	// fixtures created below.
	_ "github.com/ForceCLI/force-md/metadata/customPermissions"
)

func writeCustomPermission(t *testing.T, dir, name string) string {
	t.Helper()
	path := filepath.Join(dir, name+".customPermission-meta.xml")
	content := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<CustomPermission xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>%s</label>
</CustomPermission>
`, name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write fixture %s: %v", path, err)
	}
	return path
}

func TestOpenParallelRegistersAllFiles(t *testing.T) {
	dir := t.TempDir()

	const count = 25
	paths := make([]string, count)
	for i := range count {
		paths[i] = writeCustomPermission(t, dir, fmt.Sprintf("Perm%02d", i))
	}

	r := NewRepo()
	errs := r.OpenParallel(paths)

	if len(errs) != count {
		t.Fatalf("expected %d error slots, got %d", count, len(errs))
	}
	for i, err := range errs {
		if err != nil {
			t.Errorf("unexpected error for %s: %v", paths[i], err)
		}
	}

	items := r.Items("CustomPermission")
	if len(items) != count {
		t.Fatalf("expected %d registered items, got %d", count, len(items))
	}

	// Every fixture path should be registered exactly once.
	got := make(map[string]bool, len(items))
	for _, item := range items {
		got[string(item.GetMetadataInfo().Path())] = true
	}
	for _, p := range paths {
		if !got[p] {
			t.Errorf("path not registered: %s", p)
		}
	}
}

func TestOpenParallelReportsErrorsPerPath(t *testing.T) {
	dir := t.TempDir()

	good := writeCustomPermission(t, dir, "Good")
	missing := filepath.Join(dir, "DoesNotExist.customPermission-meta.xml")

	r := NewRepo()
	errs := r.OpenParallel([]string{good, missing})

	if errs[0] != nil {
		t.Errorf("expected nil error for valid file, got %v", errs[0])
	}
	if errs[1] == nil {
		t.Error("expected error for missing file, got nil")
	}

	// Only the valid file should have been registered.
	if items := r.Items("CustomPermission"); len(items) != 1 {
		t.Fatalf("expected 1 registered item, got %d", len(items))
	}
}

func TestOpenParallelEmpty(t *testing.T) {
	r := NewRepo()
	if errs := r.OpenParallel(nil); len(errs) != 0 {
		t.Fatalf("expected no errors for empty input, got %d", len(errs))
	}
}
