package profile

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// External data source accesses are kept and written between the description
// and field permissions.
func TestProfileRoundTripKeepsExternalDataSourceAccesses(t *testing.T) {
	contents := `<?xml version="1.0" encoding="UTF-8"?>
<Profile xmlns="http://soap.sforce.com/2006/04/metadata">
    <custom>true</custom>
    <description>Inventory users</description>
    <externalDataSourceAccesses>
        <enabled>true</enabled>
        <externalDataSource>Inventory</externalDataSource>
    </externalDataSourceAccesses>
    <fieldPermissions>
        <editable>true</editable>
        <field>Zip_Code__c.Region__c</field>
        <readable>true</readable>
    </fieldPermissions>
    <userLicense>Salesforce</userLicense>
</Profile>
`
	path := filepath.Join(t.TempDir(), "Inventory.profile")
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
		t.Errorf("round trip changed the profile:\n%s", out)
	}
}
