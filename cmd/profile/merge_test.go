package profile

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ForceCLI/force-md/metadata/permissionGranter"
	"github.com/ForceCLI/force-md/metadata/profile"
)

const emptyProfile = `<?xml version="1.0" encoding="UTF-8"?>
<Profile xmlns="http://soap.sforce.com/2006/04/metadata">
</Profile>
`

const sourcePermissionSet = `<?xml version="1.0" encoding="UTF-8"?>
<PermissionSet xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Source</label>
    <fieldPermissions>
        <editable>true</editable>
        <field>Account.Editable_Field__c</field>
        <readable>true</readable>
    </fieldPermissions>
    <fieldPermissions>
        <editable>false</editable>
        <field>Account.ReadOnly_Field__c</field>
        <readable>true</readable>
    </fieldPermissions>
    <objectPermissions>
        <allowCreate>true</allowCreate>
        <allowDelete>false</allowDelete>
        <allowEdit>true</allowEdit>
        <allowRead>true</allowRead>
        <modifyAllRecords>false</modifyAllRecords>
        <object>Custom_Object__c</object>
        <viewAllRecords>true</viewAllRecords>
    </objectPermissions>
    <userPermissions>
        <enabled>true</enabled>
        <name>ViewSetup</name>
    </userPermissions>
</PermissionSet>
`

const upgradePermissionSet = `<?xml version="1.0" encoding="UTF-8"?>
<PermissionSet xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Upgrade</label>
    <fieldPermissions>
        <editable>true</editable>
        <field>Account.ReadOnly_Field__c</field>
        <readable>true</readable>
    </fieldPermissions>
</PermissionSet>
`

func writeTempFile(t *testing.T, dir, name, contents string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestMergeGrantedPermissionsFromPermissionSet(t *testing.T) {
	tempDir := t.TempDir()
	profilePath := writeTempFile(t, tempDir, "Admin.profile-meta.xml", emptyProfile)
	sourcePath := writeTempFile(t, tempDir, "Source.permissionset-meta.xml", sourcePermissionSet)

	grant, err := permissionGranter.Open(sourcePath)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := grant.(*profile.Profile); ok {
		t.Fatal("expected permission set source, got profile")
	}
	mergeGrantedPermissions(profilePath, grant)

	p, err := profile.Open(profilePath)
	if err != nil {
		t.Fatal(err)
	}

	fields := make(map[string]struct{ editable, readable bool })
	for _, f := range p.FieldPermissions {
		fields[f.Field] = struct{ editable, readable bool }{f.Editable.ToBool(), f.Readable.ToBool()}
	}
	if got := fields["Account.Editable_Field__c"]; !got.editable || !got.readable {
		t.Errorf("Account.Editable_Field__c should be editable and readable, got %+v", got)
	}
	if got := fields["Account.ReadOnly_Field__c"]; got.editable || !got.readable {
		t.Errorf("Account.ReadOnly_Field__c should be readable only, got %+v", got)
	}

	if len(p.ObjectPermissions) != 1 || p.ObjectPermissions[0].Object != "Custom_Object__c" {
		t.Fatalf("expected object permissions for Custom_Object__c, got %+v", p.ObjectPermissions)
	}
	o := p.ObjectPermissions[0]
	if !o.AllowCreate.ToBool() || !o.AllowEdit.ToBool() || !o.AllowRead.ToBool() || !o.ViewAllRecords.ToBool() {
		t.Errorf("granted object permissions not applied: %+v", o)
	}
	if o.AllowDelete.ToBool() || o.ModifyAllRecords.ToBool() {
		t.Errorf("ungranted object permissions should stay false: %+v", o)
	}

	foundUserPerm := false
	for _, u := range p.UserPermissions {
		if u.Name == "ViewSetup" {
			foundUserPerm = true
		}
	}
	if !foundUserPerm {
		t.Error("ViewSetup user permission should be merged")
	}
}

func TestMergeGrantedPermissionsUpgradesWithoutDowngrading(t *testing.T) {
	tempDir := t.TempDir()
	profilePath := writeTempFile(t, tempDir, "Admin.profile-meta.xml", emptyProfile)
	sourcePath := writeTempFile(t, tempDir, "Source.permissionset-meta.xml", sourcePermissionSet)
	upgradePath := writeTempFile(t, tempDir, "Upgrade.permissionset-meta.xml", upgradePermissionSet)

	for _, path := range []string{sourcePath, upgradePath} {
		grant, err := permissionGranter.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		mergeGrantedPermissions(profilePath, grant)
	}
	// Re-applying the read-only source must not downgrade the upgraded field.
	grant, err := permissionGranter.Open(sourcePath)
	if err != nil {
		t.Fatal(err)
	}
	mergeGrantedPermissions(profilePath, grant)

	p, err := profile.Open(profilePath)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range p.FieldPermissions {
		if f.Field == "Account.ReadOnly_Field__c" {
			if !f.Editable.ToBool() {
				t.Error("Account.ReadOnly_Field__c should stay editable after re-merging the read-only source")
			}
			return
		}
	}
	t.Fatal("Account.ReadOnly_Field__c not found in merged profile")
}

func TestMergeCommandAcceptsProfileSource(t *testing.T) {
	tempDir := t.TempDir()
	profilePath := writeTempFile(t, tempDir, "Admin.profile-meta.xml", emptyProfile)
	sourceProfile := strings.Replace(strings.Replace(sourcePermissionSet, "PermissionSet", "Profile", 2), "    <label>Source</label>\n", "", 1)
	sourcePath := writeTempFile(t, tempDir, "Source.profile-meta.xml", sourceProfile)

	grant, err := permissionGranter.Open(sourcePath)
	if err != nil {
		t.Fatal(err)
	}
	apply, ok := grant.(*profile.Profile)
	if !ok {
		t.Fatal("expected profile source")
	}
	mergePermissions(profilePath, *apply)

	p, err := profile.Open(profilePath)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.FieldPermissions) != 2 {
		t.Errorf("expected 2 field permissions from profile source, got %d", len(p.FieldPermissions))
	}
}
