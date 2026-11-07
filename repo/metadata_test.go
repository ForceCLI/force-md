package repo

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ForceCLI/force-md/metadata/permissionset"
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

const permissionSetWithByteOrderMark = "\xef\xbb\xbf<?xml version=\"1.0\" encoding=\"utf-8\"?>\n" +
	"<PermissionSet xmlns=\"http://soap.sforce.com/2006/04/metadata\">\n" +
	"\t<description>Standard permission for managers</description>\n" +
	"\t<label>Manager</label>\n" +
	"</PermissionSet>\n"

func TestRootElementNameSkipsLeadingByteOrderMark(t *testing.T) {
	SetAllowMissingXMLDeclaration(false)

	name, err := RootElementName([]byte(permissionSetWithByteOrderMark))
	if err != nil {
		t.Fatalf("expected a file starting with a byte order mark to be XML, got error: %v", err)
	}
	if name != "PermissionSet" {
		t.Fatalf("expected PermissionSet, got %q", name)
	}
}

func TestRootElementNameRejectsByteOrderMarkAfterStart(t *testing.T) {
	SetAllowMissingXMLDeclaration(true)
	t.Cleanup(func() {
		SetAllowMissingXMLDeclaration(false)
	})

	xmlData := []byte("\n\xef\xbb\xbf<PermissionSet></PermissionSet>")
	if _, err := RootElementName(xmlData); !errors.Is(err, NotXMLError) {
		t.Fatalf("expected NotXMLError for a byte order mark after the start of the file, got %v", err)
	}
}

func TestMetadataFromPathLoadsFileStartingWithByteOrderMark(t *testing.T) {
	SetAllowMissingXMLDeclaration(false)

	path := filepath.Join(t.TempDir(), "Manager.permissionset-meta.xml")
	if err := os.WriteFile(path, []byte(permissionSetWithByteOrderMark), 0o644); err != nil {
		t.Fatal(err)
	}

	m, err := MetadataFromPath(path)
	if err != nil {
		t.Fatalf("expected the permission set to load, got error: %v", err)
	}
	ps, ok := m.(*permissionset.PermissionSet)
	if !ok {
		t.Fatalf("expected *permissionset.PermissionSet, got %T", m)
	}
	if ps.Label != "Manager" {
		t.Fatalf("expected label Manager, got %q", ps.Label)
	}
}

func TestRootElementNameRejectsTextBeforeRootElement(t *testing.T) {
	SetAllowMissingXMLDeclaration(true)
	t.Cleanup(func() {
		SetAllowMissingXMLDeclaration(false)
	})

	apex := []byte("public class NamespaceUtil {\n\tList<ApexClass> classes = [SELECT Id FROM ApexClass];\n}\n")
	if _, err := RootElementName(apex); !errors.Is(err, NotXMLError) {
		t.Fatalf("expected NotXMLError for Apex source, got %v", err)
	}
}
