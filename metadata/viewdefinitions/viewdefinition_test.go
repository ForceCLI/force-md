package viewdefinitions

import (
	"os"
	"path/filepath"
	"testing"
)

const sampleMetadataXML = `<?xml version="1.0" encoding="UTF-8"?>
<ViewDefinition xmlns="http://soap.sforce.com/2006/04/metadata">
    <apiVersion>54.0</apiVersion>
    <isProtected>false</isProtected>
    <masterLabel>App Home</masterLabel>
    <targetType>slack</targetType>
</ViewDefinition>
`

func TestOpen_ParsesViewDefinitionMetadataXml(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "app_home.view-meta.xml")
	if err := os.WriteFile(path, []byte(sampleMetadataXML), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	v, err := Open(path)
	if err != nil {
		t.Fatalf("Open returned error: %v", err)
	}

	if v.MasterLabel.Text != "App Home" {
		t.Errorf("MasterLabel = %q, want %q", v.MasterLabel.Text, "App Home")
	}
	if v.TargetType == nil || v.TargetType.Text != "slack" {
		t.Errorf("TargetType = %+v, want non-nil 'slack'", v.TargetType)
	}
	if v.ApiVersion == nil || v.ApiVersion.Text != "54.0" {
		t.Errorf("ApiVersion = %+v, want non-nil '54.0'", v.ApiVersion)
	}
	if v.IsProtected == nil || v.IsProtected.Text != "false" {
		t.Errorf("IsProtected = %+v, want non-nil 'false'", v.IsProtected)
	}
	if got, want := v.Type(), "ViewDefinition"; string(got) != want {
		t.Errorf("Type() = %q, want %q", got, want)
	}
}
