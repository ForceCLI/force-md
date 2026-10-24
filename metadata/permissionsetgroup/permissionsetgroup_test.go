package permissionsetgroup

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// A permission set group without a description is written back without an
// empty description element.
func TestPermissionSetGroupWithoutDescriptionRoundTrip(t *testing.T) {
	content := `<?xml version="1.0" encoding="UTF-8"?>
<PermissionSetGroup xmlns="http://soap.sforce.com/2006/04/metadata">
    <hasActivationRequired>false</hasActivationRequired>
    <label>Sales Workspace</label>
    <permissionSets>Sales_Workspace</permissionSets>
    <status>Updated</status>
</PermissionSetGroup>
`
	path := filepath.Join(t.TempDir(), "Sales_Workspace.permissionsetgroup")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	p, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(p)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != content {
		t.Errorf("round trip changed the permission set group:\n%s", out)
	}
}
