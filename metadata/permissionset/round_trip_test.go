package permissionset

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// Object permissions without viewAllFields are written without it, and
// external credential principal accesses stay before field permissions.
func TestPermissionSetRoundTrip(t *testing.T) {
	contents := `<?xml version="1.0" encoding="UTF-8"?>
<PermissionSet xmlns="http://soap.sforce.com/2006/04/metadata">
    <externalCredentialPrincipalAccesses>
        <enabled>true</enabled>
        <externalCredentialPrincipal>API-Principal</externalCredentialPrincipal>
    </externalCredentialPrincipalAccesses>
    <fieldPermissions>
        <editable>true</editable>
        <field>Zip_Code__c.Region__c</field>
        <readable>true</readable>
    </fieldPermissions>
    <hasActivationRequired>false</hasActivationRequired>
    <label>Access to Zip Code</label>
    <objectPermissions>
        <allowCreate>true</allowCreate>
        <allowDelete>true</allowDelete>
        <allowEdit>true</allowEdit>
        <allowRead>true</allowRead>
        <modifyAllRecords>true</modifyAllRecords>
        <object>Zip_Code__c</object>
        <viewAllRecords>true</viewAllRecords>
    </objectPermissions>
</PermissionSet>
`
	path := filepath.Join(t.TempDir(), "Access_to_Zip_Code.permissionset")
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
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
	if string(out) != contents {
		t.Errorf("round trip changed the permission set:\n%s", out)
	}
}

func TestAddObjectPermissionsSetsViewAllFields(t *testing.T) {
	p := &PermissionSet{}
	if err := p.AddObjectPermissions("Account"); err != nil {
		t.Fatal(err)
	}
	if v := p.ObjectPermissions[0].ViewAllFields; v == nil || v.Text != "false" {
		t.Errorf("expected viewAllFields false, got %v", v)
	}
}

// External data source accesses are kept and written between external
// credential principal accesses and field permissions.
func TestPermissionSetRoundTripKeepsExternalDataSourceAccesses(t *testing.T) {
	contents := `<?xml version="1.0" encoding="UTF-8"?>
<PermissionSet xmlns="http://soap.sforce.com/2006/04/metadata">
    <externalCredentialPrincipalAccesses>
        <enabled>true</enabled>
        <externalCredentialPrincipal>API-Principal</externalCredentialPrincipal>
    </externalCredentialPrincipalAccesses>
    <externalDataSourceAccesses>
        <enabled>true</enabled>
        <externalDataSource>Inventory</externalDataSource>
    </externalDataSourceAccesses>
    <fieldPermissions>
        <editable>true</editable>
        <field>Zip_Code__c.Region__c</field>
        <readable>true</readable>
    </fieldPermissions>
    <hasActivationRequired>false</hasActivationRequired>
    <label>Inventory Access</label>
</PermissionSet>
`
	path := filepath.Join(t.TempDir(), "Inventory_Access.permissionset")
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}
	p, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if len(p.ExternalDataSourceAccesses) != 1 || p.ExternalDataSourceAccesses[0].ExternalDataSource != "Inventory" {
		t.Fatalf("expected the Inventory external data source access, got %+v", p.ExternalDataSourceAccesses)
	}
	out, err := internal.Marshal(p)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != contents {
		t.Errorf("round trip changed the permission set:\n%s", out)
	}
}
