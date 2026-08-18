package testSuites

import (
	"os"
	"path/filepath"
	"testing"
)

const sampleTestSuite = `<?xml version="1.0" encoding="UTF-8"?>
<ApexTestSuite xmlns="http://soap.sforce.com/2006/04/metadata">
    <testClassName>AccountServiceTest</testClassName>
    <testClassName>ContactServiceTest</testClassName>
</ApexTestSuite>
`

func TestOpen(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Services.testSuite-meta.xml")
	if err := os.WriteFile(path, []byte(sampleTestSuite), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	s, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if len(s.TestClassNames) != 2 {
		t.Fatalf("TestClassNames count = %d, want 2", len(s.TestClassNames))
	}
	if got := s.TestClassNames[0].String(); got != "AccountServiceTest" {
		t.Errorf("TestClassNames[0] = %q, want AccountServiceTest", got)
	}
	if got := s.TestClassNames[1].String(); got != "ContactServiceTest" {
		t.Errorf("TestClassNames[1] = %q, want ContactServiceTest", got)
	}
	if got := string(s.Type()); got != "ApexTestSuite" {
		t.Errorf("Type = %q, want ApexTestSuite", got)
	}
}

func TestOpenEmptySuite(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Empty.testSuite-meta.xml")
	minimal := `<?xml version="1.0" encoding="UTF-8"?>
<ApexTestSuite xmlns="http://soap.sforce.com/2006/04/metadata"/>
`
	if err := os.WriteFile(path, []byte(minimal), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	s, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if len(s.TestClassNames) != 0 {
		t.Errorf("TestClassNames count = %d, want 0", len(s.TestClassNames))
	}
}
