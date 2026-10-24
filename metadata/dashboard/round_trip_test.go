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

func assertDashboardRoundTrip(t *testing.T, content string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "Test.dashboard")
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
}

// A grid component's chart theme, legend sorting, table setting and more
// than one chart summary, and a filter option with more than one value, are
// written back unchanged.
func TestDashboardGridComponentRoundTrip(t *testing.T) {
	assertDashboardRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<Dashboard xmlns="http://soap.sforce.com/2006/04/metadata">
    <backgroundEndColor>#FFFFFF</backgroundEndColor>
    <backgroundFadeDirection>Diagonal</backgroundFadeDirection>
    <backgroundStartColor>#FFFFFF</backgroundStartColor>
    <dashboardFilters>
        <dashboardFilterOptions>
            <operator>between</operator>
            <values>1</values>
            <values>10</values>
        </dashboardFilterOptions>
        <name>Amount</name>
    </dashboardFilters>
    <dashboardGridLayout>
        <dashboardGridComponents>
            <colSpan>4</colSpan>
            <columnIndex>0</columnIndex>
            <dashboardComponent>
                <autoselectColumnsFromReport>true</autoselectColumnsFromReport>
                <chartSummary>
                    <axisBinding>y</axisBinding>
                    <column>RowCount</column>
                </chartSummary>
                <chartSummary>
                    <aggregate>Sum</aggregate>
                    <axisBinding>y2</axisBinding>
                    <column>AMOUNT</column>
                </chartSummary>
                <componentChartTheme>dark</componentChartTheme>
                <componentType>Bar</componentType>
                <flexComponentProperties>
                    <hideChatterPhotos>true</hideChatterPhotos>
                    <useReportTableSetting>false</useReportTableSetting>
                </flexComponentProperties>
                <groupingSortProperties>
                    <groupingSorts>
                        <groupingLevel>g1</groupingLevel>
                        <inheritedReportGroupingSort>true</inheritedReportGroupingSort>
                        <sortOrder>a</sortOrder>
                    </groupingSorts>
                </groupingSortProperties>
                <report>Reports/Accounts</report>
                <sortLegendValues>true</sortLegendValues>
            </dashboardComponent>
            <rowIndex>0</rowIndex>
            <rowSpan>6</rowSpan>
        </dashboardGridComponents>
        <numberOfColumns>12</numberOfColumns>
        <rowHeight>36</rowHeight>
    </dashboardGridLayout>
    <dashboardType>SpecifiedUser</dashboardType>
    <isGridLayout>true</isGridLayout>
    <runningUser>user@example.com</runningUser>
    <textColor>#000000</textColor>
    <title>Grid</title>
    <titleColor>#000000</titleColor>
    <titleSize>12</titleSize>
</Dashboard>
`)
}

// A section component's chart summaries, grouping columns, including a nil
// one, and legend sorting are written back unchanged.
func TestDashboardSectionComponentRoundTrip(t *testing.T) {
	assertDashboardRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<Dashboard xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <backgroundEndColor>#FFFFFF</backgroundEndColor>
    <backgroundFadeDirection>Diagonal</backgroundFadeDirection>
    <backgroundStartColor>#FFFFFF</backgroundStartColor>
    <leftSection>
        <columnSize>Medium</columnSize>
        <components>
            <autoselectColumnsFromReport>true</autoselectColumnsFromReport>
            <chartSummary>
                <aggregate>Sum</aggregate>
                <axisBinding>y</axisBinding>
                <column>AMOUNT</column>
            </chartSummary>
            <componentType>Column</componentType>
            <groupingColumn>STAGE_NAME</groupingColumn>
            <groupingColumn xsi:nil="true"/>
            <groupingSortProperties/>
            <report>Reports/Opportunities</report>
            <sortLegendValues>false</sortLegendValues>
        </components>
    </leftSection>
    <middleSection>
        <columnSize>Medium</columnSize>
        <components>
            <autoselectColumnsFromReport>true</autoselectColumnsFromReport>
            <componentType>Table</componentType>
            <groupingSortProperties>
                <groupingSorts>
                    <groupingLevel>g1</groupingLevel>
                    <sortColumn>RowCount</sortColumn>
                    <sortOrder>d</sortOrder>
                </groupingSorts>
            </groupingSortProperties>
            <report>Reports/Accounts</report>
        </components>
    </middleSection>
    <rightSection>
        <columnSize>Medium</columnSize>
    </rightSection>
    <textColor>#000000</textColor>
    <title>Sections</title>
    <titleColor>#000000</titleColor>
    <titleSize>12</titleSize>
</Dashboard>
`)
}
