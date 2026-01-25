package repo

import "testing"

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
