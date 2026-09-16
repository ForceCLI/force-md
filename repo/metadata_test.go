package repo

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestRootElementNameRequiresXMLDeclarationByDefault(t *testing.T) {
	SetAllowMissingXMLDeclaration(false)

	if _, err := RootElementName([]byte("<PermissionSet></PermissionSet>")); err == nil {
		t.Fatal("expected error when XML declaration is missing")
	}
}

func TestRootElementNameAllowsMissingDeclarationWhenEnabled(t *testing.T) {
	SetAllowMissingXMLDeclaration(true)
	t.Cleanup(func() {
		SetAllowMissingXMLDeclaration(false)
	})

	name, err := RootElementName([]byte("<PermissionSet></PermissionSet>"))
	if err != nil {
		t.Fatalf("expected missing declaration to be accepted, got error: %v", err)
	}
	if name != "PermissionSet" {
		t.Fatalf("expected PermissionSet, got %q", name)
	}
}

func TestMetadataFromPathReturnsNotFoundForUnidentifiableAbsolutePath(t *testing.T) {
	// An empty file has no root element, so every lookup falls through to the
	// walk up the parent directories, which must stop at the filesystem root.
	path := filepath.Join(t.TempDir(), "Empty.permissionset-meta.xml")
	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatal(err)
	}

	done := make(chan error, 1)
	go func() {
		_, err := MetadataFromPath(path)
		done <- err
	}()
	select {
	case err := <-done:
		if !errors.Is(err, MetadataFileNotFound) {
			t.Fatalf("expected MetadataFileNotFound, got %v", err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("MetadataFromPath did not return for an unidentifiable absolute path")
	}
}
