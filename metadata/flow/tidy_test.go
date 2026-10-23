package flow

import (
	"fmt"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
)

const unbalancedFormulaFlow = `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
    <formulas>
        <name>Unbalanced</name>
        <dataType>String</dataType>
        <expression>{!$Record.Account__c.Name + &quot;-&quot;+TEXT(YEAR(DATEVALUE({!$Record.CreatedDate})))</expression>
    </formulas>
    <formulas>
        <name>Balanced</name>
        <dataType>Date</dataType>
        <expression>{!$Record.Referral_Date__c}+14</expression>
    </formulas>
    <processType>AutoLaunchedFlow</processType>
    <status>%s</status>
</Flow>
`

func formatErrorEntries(hook *test.Hook) []*log.Entry {
	var entries []*log.Entry
	for _, e := range hook.AllEntries() {
		if strings.HasPrefix(e.Message, "error formatting") {
			entries = append(entries, e)
		}
	}
	return entries
}

func tidyWithHook(t *testing.T, status string) (*Flow, *test.Hook) {
	t.Helper()
	f := openFlow(t, fmt.Sprintf(unbalancedFormulaFlow, status))
	hook := test.NewGlobal()
	t.Cleanup(hook.Reset)
	level := log.GetLevel()
	log.SetLevel(log.DebugLevel)
	t.Cleanup(func() { log.SetLevel(level) })
	f.Tidy()
	return f, hook
}

func TestTidyWarnsOnUnparseableFormulaInActiveFlow(t *testing.T) {
	_, hook := tidyWithHook(t, "Active")
	entries := formatErrorEntries(hook)
	if len(entries) != 1 || entries[0].Level != log.WarnLevel {
		t.Fatalf("expected one warning, got %v", entries)
	}
}

func TestTidyLogsUnparseableFormulaInInvalidDraftAtDebug(t *testing.T) {
	f, hook := tidyWithHook(t, "InvalidDraft")
	entries := formatErrorEntries(hook)
	if len(entries) != 1 || entries[0].Level != log.DebugLevel {
		t.Fatalf("expected one debug entry, got %v", entries)
	}
	if got, want := f.Formulas[0].Expression.Text, `{!$Record.Account__c.Name + &quot;-&quot;+TEXT(YEAR(DATEVALUE({!$Record.CreatedDate})))`; got != want {
		t.Errorf("unparseable formula changed: got %q, want %q", got, want)
	}
	if got, want := f.Formulas[1].Expression.Text, `{!$Record.Referral_Date__c} + 14`; got != want {
		t.Errorf("parseable formula not formatted: got %q, want %q", got, want)
	}
}
