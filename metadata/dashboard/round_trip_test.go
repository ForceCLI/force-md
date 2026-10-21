package dashboard

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// A dashboard without the optional isGridLayout and displayUnits elements is
// written back without them, rather than with empty elements, and one with
// them, and with a section component's showPicturesOnCharts and
// showPicturesOnTables, keeps them. owner is written between the sections,
// in the order of the Dashboard type.
func TestDashboardOptionalElementsRoundTrip(t *testing.T) {
	content := `<?xml version="1.0" encoding="UTF-8"?>
<Dashboard xmlns="http://soap.sforce.com/2006/04/metadata">
    <backgroundEndColor>#FFFFFF</backgroundEndColor>
    <backgroundFadeDirection>Diagonal</backgroundFadeDirection>
    <backgroundStartColor>#FFFFFF</backgroundStartColor>
    <dashboardType>SpecifiedUser</dashboardType>
    <leftSection>
        <columnSize>Medium</columnSize>
        <components>
            <autoselectColumnsFromReport>true</autoselectColumnsFromReport>
            <componentType>Table</componentType>
            <groupingSortProperties/>
            <header>Accounts</header>
            <indicatorHighColor>#54C254</indicatorHighColor>
            <indicatorLowColor>#C25454</indicatorLowColor>
            <indicatorMiddleColor>#C2C254</indicatorMiddleColor>
            <report>Reports/Accounts</report>
        </components>
    </leftSection>
    <owner>owner@example.com</owner>
    <rightSection>
        <columnSize>Medium</columnSize>
    </rightSection>
    <runningUser>user@example.com</runningUser>
    <textColor>#000000</textColor>
    <title>Accounts</title>
    <titleColor>#000000</titleColor>
    <titleSize>12</titleSize>
</Dashboard>
`
	path := filepath.Join(t.TempDir(), "Accounts.dashboard")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	d, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != content {
		t.Errorf("round trip changed the dashboard:\n%s", out)
	}

	withValues := `<?xml version="1.0" encoding="UTF-8"?>
<Dashboard xmlns="http://soap.sforce.com/2006/04/metadata">
    <backgroundEndColor>#FFFFFF</backgroundEndColor>
    <backgroundFadeDirection>Diagonal</backgroundFadeDirection>
    <backgroundStartColor>#FFFFFF</backgroundStartColor>
    <isGridLayout>false</isGridLayout>
    <leftSection>
        <columnSize>Medium</columnSize>
        <components>
            <autoselectColumnsFromReport>true</autoselectColumnsFromReport>
            <componentType>Bar</componentType>
            <displayUnits>Auto</displayUnits>
            <groupingSortProperties/>
            <report>Reports/Accounts</report>
            <showPicturesOnCharts>false</showPicturesOnCharts>
            <showPicturesOnTables>true</showPicturesOnTables>
        </components>
    </leftSection>
    <textColor>#000000</textColor>
    <title>Accounts</title>
    <titleColor>#000000</titleColor>
    <titleSize>12</titleSize>
</Dashboard>
`
	if err := os.WriteFile(path, []byte(withValues), 0644); err != nil {
		t.Fatal(err)
	}
	d, err = Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err = internal.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != withValues {
		t.Errorf("round trip changed the dashboard:\n%s", out)
	}
}
