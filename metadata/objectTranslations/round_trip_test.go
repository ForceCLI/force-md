package objectTranslations

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/metadata"
)

// Splitting a translation into source format keeps every element, including
// untranslated values, which Salesforce writes as comments.
func TestDecomposeKeepsAllTranslationElements(t *testing.T) {
	path := filepath.Join(t.TempDir(), "Van__c-fr.objectTranslation")
	content := `<?xml version="1.0" encoding="UTF-8"?>
<CustomObjectTranslation xmlns="http://soap.sforce.com/2006/04/metadata">
    <fields>
        <help>Le chef d&apos;équipe</help>
        <label><!-- Primary Team Lead --></label>
        <lookupFilter>
            <errorMessage><!-- Staff must be a Team Lead. --></errorMessage>
        </lookupFilter>
        <name>Primary_Team_Lead__c</name>
        <relationshipLabel>Camionnettes</relationshipLabel>
    </fields>
    <layouts>
        <layout>Van Layout</layout>
        <layoutType>CaseClose</layoutType>
        <sections>
            <label><!-- Fields --></label>
            <section>Fields</section>
        </sections>
    </layouts>
    <recordTypes>
        <label>Camionnette</label>
        <name>Van</name>
    </recordTypes>
</CustomObjectTranslation>
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

	wantObject := `<?xml version="1.0" encoding="UTF-8"?>
<CustomObjectTranslation xmlns="http://soap.sforce.com/2006/04/metadata">
    <layouts>
        <layout>Van Layout</layout>
        <layoutType>CaseClose</layoutType>
        <sections>
            <label><!-- Fields --></label>
            <section>Fields</section>
        </sections>
    </layouts>
    <recordTypes>
        <label>Camionnette</label>
        <name>Van</name>
    </recordTypes>
</CustomObjectTranslation>
`
	if got := string(files["objectTranslations/Van__c-fr/Van__c-fr.objectTranslation-meta.xml"]); got != wantObject {
		t.Errorf("object translation =\n%s\nwant\n%s", got, wantObject)
	}

	wantField := `<?xml version="1.0" encoding="UTF-8"?>
<CustomFieldTranslation xmlns="http://soap.sforce.com/2006/04/metadata">
    <help>Le chef d&apos;équipe</help>
    <label><!-- Primary Team Lead --></label>
    <lookupFilter>
        <errorMessage><!-- Staff must be a Team Lead. --></errorMessage>
    </lookupFilter>
    <name>Primary_Team_Lead__c</name>
    <relationshipLabel>Camionnettes</relationshipLabel>
</CustomFieldTranslation>
`
	if got := string(files["objectTranslations/Van__c-fr/Primary_Team_Lead__c.fieldTranslation-meta.xml"]); got != wantField {
		t.Errorf("field translation =\n%s\nwant\n%s", got, wantField)
	}
}
