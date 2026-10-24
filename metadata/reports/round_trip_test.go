package report

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

func assertRoundTrip(t *testing.T, content string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "reports", "Folder", "Test.report")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	r, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(r)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != content {
		t.Errorf("round trip changed the report:\n%s", out)
	}
}

// A joined report's blocks, its top-level blockInfo with a nil blockId, and
// the description of a cross-block aggregate are written back unchanged.
func TestJoinedReportRoundTrip(t *testing.T) {
	assertRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<Report xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <aggregates>
        <calculatedFormula>B0#Campaign.Campaign_Point__c:SUM</calculatedFormula>
        <datatype>number</datatype>
        <description>Sum of points</description>
        <developerName>FORMULA1</developerName>
        <isActive>true</isActive>
        <isCrossBlock>true</isCrossBlock>
        <masterLabel>Point Total</masterLabel>
        <scale>0</scale>
    </aggregates>
    <block>
        <blockInfo>
            <aggregateReferences>
                <aggregate>FORMULA1</aggregate>
            </aggregateReferences>
            <blockId>B0</blockId>
            <joinTable>u</joinTable>
        </blockInfo>
        <columns>
            <aggregateTypes>Sum</aggregateTypes>
            <field>Campaign.Campaign_Point__c</field>
        </columns>
        <filter>
            <booleanFilter>1 OR 2</booleanFilter>
            <criteriaItems>
                <column>CAMPAIGN_NAME</column>
                <columnToColumn>false</columnToColumn>
                <isUnlocked>true</isUnlocked>
                <operator>notEqual</operator>
                <value></value>
            </criteriaItems>
            <criteriaItems>
                <column>TYPE</column>
                <columnToColumn>false</columnToColumn>
                <isUnlocked>true</isUnlocked>
                <operator>equals</operator>
                <value>Webinar</value>
            </criteriaItems>
        </filter>
        <format>Summary</format>
        <name>Campaigns block 1</name>
        <params>
            <name>scopeid</name>
            <value></value>
        </params>
        <reportType>CampaignList</reportType>
        <scope>orgAll</scope>
        <showGrandTotal>true</showGrandTotal>
        <showSubTotals>true</showSubTotals>
        <sortColumn>Campaign.Campaign_Point__c</sortColumn>
        <sortOrder>Desc</sortOrder>
        <timeFrameFilter>
            <dateColumn>START_DATE</dateColumn>
            <endDate>2024-12-31</endDate>
            <interval>INTERVAL_CUSTOM</interval>
            <startDate>2024-01-01</startDate>
        </timeFrameFilter>
    </block>
    <blockInfo>
        <blockId xsi:nil="true"/>
        <joinTable>u</joinTable>
    </blockInfo>
    <format>MultiBlock</format>
    <name>Joined</name>
    <params>
        <name>lsk</name>
        <value>1</value>
    </params>
    <reportType>CampaignList</reportType>
    <showDetails>true</showDetails>
</Report>
`)
}

// Numeric bucket ranges, cross filter criteria, a custom time frame's end
// date, more than one aggregate per column, and the aggregate of color
// ranges and formatting rules are written back unchanged.
func TestReportElementsRoundTrip(t *testing.T) {
	assertRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<Report xmlns="http://soap.sforce.com/2006/04/metadata">
    <buckets>
        <bucketType>number</bucketType>
        <developerName>BucketField_Age</developerName>
        <masterLabel>Age</masterLabel>
        <nullTreatment>z</nullTreatment>
        <sourceColumnName>AGE</sourceColumnName>
        <useOther>false</useOther>
        <values>
            <sourceValues>
                <to>548</to>
            </sourceValues>
            <value>18 Months</value>
        </values>
        <values>
            <sourceValues>
                <from>548</from>
                <to>1,095</to>
            </sourceValues>
            <value>36 Months</value>
        </values>
        <values>
            <sourceValues>
                <sourceValue>Other</sourceValue>
            </sourceValues>
            <value>Other</value>
        </values>
    </buckets>
    <chart>
        <backgroundColor1>#FFFFFF</backgroundColor1>
        <backgroundColor2>#FFFFFF</backgroundColor2>
        <backgroundFadeDir>Diagonal</backgroundFadeDir>
        <chartSummaries>
            <axisBinding>y</axisBinding>
            <column>RowCount</column>
        </chartSummaries>
        <chartSummaries>
            <aggregate>Sum</aggregate>
            <axisBinding>y2</axisBinding>
            <column>AMOUNT</column>
        </chartSummaries>
        <chartType>VerticalColumnLine</chartType>
        <enableHoverLabels>false</enableHoverLabels>
        <expandOthers>true</expandOthers>
        <groupingColumn>CREATED_DATE</groupingColumn>
        <location>CHART_BOTTOM</location>
        <showAxisLabels>true</showAxisLabels>
        <showPercentage>false</showPercentage>
        <showTotal>false</showTotal>
        <size>Medium</size>
        <summaryAxisRange>Auto</summaryAxisRange>
        <textColor>#000000</textColor>
        <textSize>12</textSize>
        <titleColor>#000000</titleColor>
        <titleSize>18</titleSize>
    </chart>
    <colorRanges>
        <aggregate>Sum</aggregate>
        <columnName>AMOUNT</columnName>
        <highColor>#FFFFFF</highColor>
        <lowBreakpoint>1.0</lowBreakpoint>
        <lowColor>#FFFFFF</lowColor>
        <midColor>#54C254</midColor>
    </colorRanges>
    <columns>
        <aggregateTypes>Sum</aggregateTypes>
        <aggregateTypes>Average</aggregateTypes>
        <field>AMOUNT</field>
    </columns>
    <crossFilters>
        <criteriaItems>
            <column>CreatedBy</column>
            <isUnlocked>false</isUnlocked>
            <operator>equals</operator>
            <value>someone</value>
        </criteriaItems>
        <operation>with</operation>
        <primaryTableColumn>CONTACT_ID</primaryTableColumn>
        <relatedTable>OCR-Opportunity</relatedTable>
        <relatedTableJoinColumn>Contact</relatedTableJoinColumn>
    </crossFilters>
    <format>Summary</format>
    <formattingRules>
        <aggregate>Sum</aggregate>
        <columnName>AMOUNT</columnName>
        <values>
            <backgroundColor>#FFFFFF</backgroundColor>
            <rangeUpperBound>1.0</rangeUpperBound>
        </values>
    </formattingRules>
    <groupingsDown>
        <dateGranularity>Month</dateGranularity>
        <field>CREATED_DATE</field>
        <sortOrder>Asc</sortOrder>
    </groupingsDown>
    <name>Elements</name>
    <reportType>Opportunity</reportType>
    <scope>organization</scope>
    <showDetails>false</showDetails>
    <timeFrameFilter>
        <dateColumn>CREATED_DATE</dateColumn>
        <endDate>2024-12-31</endDate>
        <interval>INTERVAL_CUSTOM</interval>
        <startDate>2024-01-01</startDate>
    </timeFrameFilter>
</Report>
`)
}

// A report in a folder whose name ends with "reports" keeps its folder path
// instead of being written to a file named after the folder.
func TestReportFilesKeepFolderEndingInReports(t *testing.T) {
	path := filepath.Join(t.TempDir(), "src", "reports", "Joe_File", "RegistrationsDBreports", "Total_Referrals.report")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Report xmlns="http://soap.sforce.com/2006/04/metadata">
    <format>Tabular</format>
    <name>Total Referrals</name>
    <reportType>Opportunity</reportType>
</Report>
`), 0644); err != nil {
		t.Fatal(err)
	}
	r, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	files, err := r.Files(metadata.SourceFormat)
	if err != nil {
		t.Fatalf("Files: %v", err)
	}
	want := filepath.Join("reports", "Joe_File", "RegistrationsDBreports", "Total_Referrals.report-meta.xml")
	if _, ok := files[want]; !ok || len(files) != 1 {
		t.Errorf("Files() = %v, want only %s", keys(files), want)
	}
}

func keys(m map[string][]byte) []string {
	var k []string
	for name := range m {
		k = append(k, name)
	}
	return k
}
