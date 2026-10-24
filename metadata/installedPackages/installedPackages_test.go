package installedPackages

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// A nil activateRSS keeps the xsi namespace prefix on the attribute and on
// its declaration.
func TestInstalledPackageRoundTrip(t *testing.T) {
	content := `<?xml version="1.0" encoding="UTF-8"?>
<InstalledPackage xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <activateRSS xsi:nil="true"/>
    <versionNumber>1.110</versionNumber>
</InstalledPackage>
`
	path := filepath.Join(t.TempDir(), "Test.installedPackage")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
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
	if string(out) != content {
		t.Errorf("round trip changed the installed package:\n%s", out)
	}
}
