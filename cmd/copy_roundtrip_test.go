package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

// Copying metadata to source format and back to metadata format writes each
// file as it was read.
func TestCopyRoundTripPreservesMetadata(t *testing.T) {
	files := map[string]string{
		"brandingSets/LEXTHEMINGExample.brandingSet": `<?xml version="1.0" encoding="UTF-8"?>
<BrandingSet xmlns="http://soap.sforce.com/2006/04/metadata">
    <brandingSetProperty>
        <propertyName>BRAND_COLOR</propertyName>
        <propertyValue>#e99a04</propertyValue>
    </brandingSetProperty>
    <brandingSetProperty>
        <propertyName>BANNER_IMAGE</propertyName>
        <propertyValue></propertyValue>
    </brandingSetProperty>
    <masterLabel>Example</masterLabel>
    <type>lightningExperience</type>
</BrandingSet>
`,
		"lightningExperienceThemes/Example.lightningExperienceTheme": `<?xml version="1.0" encoding="UTF-8"?>
<LightningExperienceTheme xmlns="http://soap.sforce.com/2006/04/metadata">
    <defaultBrandingSet>LEXTHEMINGExample</defaultBrandingSet>
    <designSystemVersion>SLDS_v2</designSystemVersion>
    <isDarkModeEnabled>true</isDarkModeEnabled>
    <masterLabel>Example</masterLabel>
    <shouldOverrideLoadingImage>false</shouldOverrideLoadingImage>
</LightningExperienceTheme>
`,
		"campaignInfluenceModels/primaryModel.campaignInfluenceModel": `<?xml version="1.0" encoding="UTF-8"?>
<CampaignInfluenceModel xmlns="http://soap.sforce.com/2006/04/metadata">
    <isActive>true</isActive>
    <isDefaultModel>true</isDefaultModel>
    <isModelLocked>true</isModelLocked>
    <modelDescription>Primary campaign gets 100% of the revenue share</modelDescription>
    <name>Primary Campaign Source Model</name>
    <recordPreference>AllRecords</recordPreference>
</CampaignInfluenceModel>
`,
		"escalationRules/Case.escalationRules": `<?xml version="1.0" encoding="UTF-8"?>
<EscalationRules xmlns="http://soap.sforce.com/2006/04/metadata">
    <escalationRule>
        <fullName>Standard</fullName>
        <active>true</active>
        <ruleEntry>
            <businessHoursSource>None</businessHoursSource>
            <criteriaItems>
                <field>Case.Priority</field>
                <operation>equals</operation>
                <value>High</value>
            </criteriaItems>
            <escalationAction>
                <assignedTo>support@example.com</assignedTo>
                <assignedToType>User</assignedToType>
                <minutesToEscalation>60</minutesToEscalation>
                <notifyCaseOwner>true</notifyCaseOwner>
            </escalationAction>
            <escalationStartTime>CaseCreation</escalationStartTime>
        </ruleEntry>
    </escalationRule>
</EscalationRules>
`,
		"assignmentRules/Lead.assignmentRules": `<?xml version="1.0" encoding="UTF-8"?>
<AssignmentRules xmlns="http://soap.sforce.com/2006/04/metadata">
    <assignmentRule>
        <fullName>Standard</fullName>
        <active>true</active>
        <ruleEntry>
            <assignedTo>user@example.com</assignedTo>
            <assignedToType>User</assignedToType>
            <booleanFilter>1 or 2</booleanFilter>
            <criteriaItems>
                <field>Lead.Form__c</field>
                <operation>contains</operation>
                <value>Join</value>
            </criteriaItems>
            <criteriaItems>
                <field>Lead.LeadSource</field>
                <operation>equals</operation>
                <value>Web</value>
            </criteriaItems>
        </ruleEntry>
        <ruleEntry>
            <assignedTo>Admissions</assignedTo>
            <assignedToType>Queue</assignedToType>
        </ruleEntry>
    </assignmentRule>
</AssignmentRules>
`,
		"settings/AccountPlan.settings": `<?xml version="1.0" encoding="UTF-8"?>
<AccountPlanSettings xmlns="http://soap.sforce.com/2006/04/metadata">
    <enableAccountPlan>false</enableAccountPlan>
</AccountPlanSettings>
`,
		"settings/Voice.settings": `<?xml version="1.0" encoding="UTF-8"?>
<VoiceSettings xmlns="http://soap.sforce.com/2006/04/metadata"/>

`,
		"settings/Activities.settings": `<?xml version="1.0" encoding="UTF-8"?>
<ActivitiesSettings xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <allowUsersToRelateMultipleContactsToTasksAndEvents>true</allowUsersToRelateMultipleContactsToTasksAndEvents>
</ActivitiesSettings>
`,
		"standardValueSets/AccountType.standardValueSet": `<?xml version="1.0" encoding="UTF-8"?>
<StandardValueSet xmlns="http://soap.sforce.com/2006/04/metadata">
    <sorted>false</sorted>
    <standardValue>
        <fullName>Competitor</fullName>
        <default>false</default>
        <isActive>false</isActive>
        <label>Competitor</label>
    </standardValue>
</StandardValueSet>
`,
		"messageChannels/Refresh_View.messageChannel": `<?xml version="1.0" encoding="UTF-8"?>
<LightningMessageChannel xmlns="http://soap.sforce.com/2006/04/metadata">
    <description>Signals that the record page should refresh</description>
    <isExposed>true</isExposed>
    <masterLabel>Refresh View</masterLabel>
</LightningMessageChannel>
`,
		"externalServiceRegistrations/Example.externalServiceRegistration": `<?xml version="1.0" encoding="UTF-8"?>
<ExternalServiceRegistration xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Example</label>
    <namedCredentialReference>Example</namedCredentialReference>
    <operations>
        <active>true</active>
        <name>getBookings</name>
    </operations>
    <registrationProviderType>SchemaInferred</registrationProviderType>
    <schema>{&quot;openapi&quot;:&quot;3.0.1&quot;}</schema>
    <schemaType>OpenApi3</schemaType>
    <serviceBinding>{&quot;host&quot;:&quot;&quot;,&quot;basePath&quot;:&quot;/&quot;}</serviceBinding>
    <status>Complete</status>
    <systemVersion>5</systemVersion>
</ExternalServiceRegistration>
`,
		"reports/Joe_File-meta.xml": `<?xml version="1.0" encoding="UTF-8"?>
<ReportFolder xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>Joe File</name>
</ReportFolder>
`,
		"reports/Joe_File/RegistrationsDBreports-meta.xml": `<?xml version="1.0" encoding="UTF-8"?>
<ReportFolder xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>Registrations DB reports</name>
</ReportFolder>
`,
		"reports/Joe_File/RegistrationsDBreports/Total_Referrals.report": `<?xml version="1.0" encoding="UTF-8"?>
<Report xmlns="http://soap.sforce.com/2006/04/metadata">
    <format>Tabular</format>
    <name>Total Referrals</name>
    <reportType>Opportunity</reportType>
    <showDetails>true</showDetails>
</Report>
`,
		"reports/Joe_File/RegistrationsDBreports/Provider_Referrals.report": `<?xml version="1.0" encoding="UTF-8"?>
<Report xmlns="http://soap.sforce.com/2006/04/metadata">
    <format>Tabular</format>
    <name>Provider Referrals</name>
    <reportType>Opportunity</reportType>
    <showDetails>true</showDetails>
</Report>
`,
	}

	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "src")
	for name, content := range files {
		path := filepath.Join(sourceDir, name)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	sfdxDir := filepath.Join(tempDir, "sfdx")
	if err := CopyMetadata(sourceDir, sfdxDir, "source", ""); err != nil {
		t.Fatal(err)
	}
	backDir := filepath.Join(tempDir, "back")
	if err := CopyMetadata(sfdxDir, backDir, "metadata", ""); err != nil {
		t.Fatal(err)
	}

	for name, content := range files {
		got, err := os.ReadFile(filepath.Join(backDir, name))
		if err != nil {
			t.Errorf("%s not written back: %v", name, err)
			continue
		}
		if string(got) != content {
			t.Errorf("round trip changed %s:\n%s", name, got)
		}
	}
}
