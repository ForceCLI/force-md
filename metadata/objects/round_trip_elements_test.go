package objects

import (
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

func roundTrip(t *testing.T, name string, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	o, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(o)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	return string(out)
}

// Elements that only older API versions or less common objects use are
// written back unchanged.
func TestObjectElementsRoundTrip(t *testing.T) {
	content := `<?xml version="1.0" encoding="UTF-8"?>
<CustomObject xmlns="http://soap.sforce.com/2006/04/metadata">
    <articleTypeChannelDisplay>
        <articleTypeTemplates>
            <channel>App</channel>
            <template>Tab</template>
        </articleTypeTemplates>
    </articleTypeChannelDisplay>
    <customSettingsVisibility>Public</customSettingsVisibility>
    <enableChangeDataCapture>false</enableChangeDataCapture>
    <fields>
        <fullName>Kind__c</fullName>
        <complianceGroup>None</complianceGroup>
        <globalPicklist>Kinds</globalPicklist>
        <isAIPredictionField>true</isAIPredictionField>
        <label>Kind</label>
        <picklist>
            <controllingField>Type__c</controllingField>
            <picklistValues>
                <fullName>Closed Won</fullName>
                <closed>true</closed>
                <color>#00FF00</color>
                <controllingFieldValues>Sale</controllingFieldValues>
                <controllingFieldValues>Renewal</controllingFieldValues>
                <default>false</default>
                <forecastCategory>Closed</forecastCategory>
                <probability>100</probability>
                <won>true</won>
            </picklistValues>
            <restrictedPicklist>true</restrictedPicklist>
            <sorted>false</sorted>
        </picklist>
        <type>Picklist</type>
    </fields>
    <indexes>
        <fullName>Time_Metric</fullName>
        <fields>
            <name>Timestamp__c</name>
            <sortDirection>DESC</sortDirection>
        </fields>
        <label>Time/Metric</label>
    </indexes>
    <label>Widget</label>
    <listViews>
        <fullName>Team</fullName>
        <filterScope>Everything</filterScope>
        <label>Team</label>
        <sharedTo>
            <allPartnerUsers></allPartnerUsers>
            <roleAndSubordinatesInternal>Sales_Director</roleAndSubordinatesInternal>
        </sharedTo>
    </listViews>
    <nameField>
        <encryptionScheme>None</encryptionScheme>
        <label>Widget Name</label>
        <type>Text</type>
    </nameField>
    <namedFilters>
        <fullName>nf_Distributor</fullName>
        <active>true</active>
        <errorMessage>Select a distributor</errorMessage>
        <field>Widget__c.Account__c</field>
        <filterItems>
            <field>Account.Type</field>
            <operation>equals</operation>
            <value>Distributor</value>
        </filterItems>
        <isOptional>false</isOptional>
        <name>Widget_Distributor</name>
        <sourceObject>Account</sourceObject>
    </namedFilters>
    <searchLayouts>
        <massQuickActions>Add_Referral_Source</massQuickActions>
    </searchLayouts>
    <webLinks>
        <fullName>Selector</fullName>
        <availability>online</availability>
        <displayType>button</displayType>
        <linkType>sControl</linkType>
        <masterLabel>Selector</masterLabel>
        <openType>replace</openType>
        <protected>false</protected>
        <scontrol>ProductSelector</scontrol>
    </webLinks>
</CustomObject>
`
	// Empty elements are written self-closing.
	want := strings.Replace(content, "<allPartnerUsers></allPartnerUsers>", "<allPartnerUsers/>", 1)
	if got := roundTrip(t, "Widget__c.object", content); got != want {
		t.Errorf("round trip changed the object:\n%s", got)
	}
}

// A field set member without isFieldManaged is written without one.
func TestFieldSetMemberWithoutIsFieldManagedRoundTrips(t *testing.T) {
	content := `<?xml version="1.0" encoding="UTF-8"?>
<CustomObject xmlns="http://soap.sforce.com/2006/04/metadata">
    <fieldSets>
        <fullName>Summary</fullName>
        <description>Summary fields</description>
        <displayedFields>
            <field>Name</field>
            <isRequired>false</isRequired>
        </displayedFields>
        <label>Summary</label>
    </fieldSets>
    <label>Widget</label>
</CustomObject>
`
	if got := roundTrip(t, "Widget__c.object", content); got != content {
		t.Errorf("round trip changed the object:\n%s", got)
	}
}

// Source format names an index file after the index's fullName, since its
// label can contain characters, such as /, that cannot appear in a file name.
func TestDecomposeNamesIndexFileAfterFullName(t *testing.T) {
	path := filepath.Join(t.TempDir(), "Metric__b.object")
	content := `<?xml version="1.0" encoding="UTF-8"?>
<CustomObject xmlns="http://soap.sforce.com/2006/04/metadata">
    <indexes>
        <fullName>Time_Metric</fullName>
        <fields>
            <name>Timestamp__c</name>
            <sortDirection>DESC</sortDirection>
        </fields>
        <label>Time/Metric</label>
    </indexes>
    <label>Metric</label>
</CustomObject>
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	o, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	files, err := o.Files(metadata.SourceFormat)
	if err != nil {
		t.Fatalf("Files: %v", err)
	}
	if _, ok := files["objects/Metric__b/indexes/Time_Metric.index-meta.xml"]; !ok {
		t.Errorf("index file not named after fullName: %v", slices.Collect(maps.Keys(files)))
	}
}
