package validationrule

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenReadsFormulaInCDATASection(t *testing.T) {
	path := filepath.Join(t.TempDir(), "No_Duplicates.validationRule-meta.xml")
	content := `<?xml version="1.0" encoding="UTF-8"?>
<ValidationRule xmlns="http://soap.sforce.com/2006/04/metadata">
    <fullName>No_Duplicates</fullName>
    <active>true</active>
    <errorConditionFormula><![CDATA[AND(
    ISPICKVAL(Status__c, "R&D"),
    Count__c >= 1
)]]></errorConditionFormula>
    <errorMessage>Only one &amp; no more.</errorMessage>
</ValidationRule>
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	rule, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	want := "AND(\n    ISPICKVAL(Status__c, \"R&D\"),\n    Count__c >= 1\n)"
	if got := rule.ErrorConditionFormula.String(); got != want {
		t.Errorf("formula = %q, want %q", got, want)
	}
}
