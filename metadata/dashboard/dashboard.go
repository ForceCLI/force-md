package dashboard

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "Dashboard"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Dashboard struct {
	metadata.MetadataInfo
	XMLName                 xml.Name     `xml:"Dashboard"`
	Xmlns                   string       `xml:"xmlns,attr"`
	Xsi                     string       `xml:"xmlns:xsi,attr,omitempty"`
	BackgroundEndColor      TextLiteral  `xml:"backgroundEndColor"`
	BackgroundFadeDirection TextLiteral  `xml:"backgroundFadeDirection"`
	BackgroundStartColor    TextLiteral  `xml:"backgroundStartColor"`
	ChartTheme              *TextLiteral `xml:"chartTheme"`
	ColorPalette            *TextLiteral `xml:"colorPalette"`
	DashboardChartTheme     *TextLiteral `xml:"dashboardChartTheme"`
	DashboardColorPalette   *TextLiteral `xml:"dashboardColorPalette"`
	DashboardFilters        []struct {
		DashboardFilterOptions []struct {
			Operator TextLiteral   `xml:"operator"`
			Values   []TextLiteral `xml:"values"`
		} `xml:"dashboardFilterOptions"`
		Name TextLiteral `xml:"name"`
	} `xml:"dashboardFilters"`
	DashboardGridLayout *struct {
		DashboardGridComponents []struct {
			ColSpan            TextLiteral        `xml:"colSpan"`
			ColumnIndex        TextLiteral        `xml:"columnIndex"`
			DashboardComponent DashboardComponent `xml:"dashboardComponent"`
			RowIndex           TextLiteral        `xml:"rowIndex"`
			RowSpan            TextLiteral        `xml:"rowSpan"`
		} `xml:"dashboardGridComponents"`
		NumberOfColumns TextLiteral `xml:"numberOfColumns"`
		RowHeight       TextLiteral `xml:"rowHeight"`
	} `xml:"dashboardGridLayout"`
	DashboardType *TextLiteral      `xml:"dashboardType"`
	Description   *TextLiteral      `xml:"description"`
	IsGridLayout  *TextLiteral      `xml:"isGridLayout"`
	LeftSection   *DashboardSection `xml:"leftSection"`
	MiddleSection *DashboardSection `xml:"middleSection"`
	Owner         *TextLiteral      `xml:"owner"`
	RightSection  *DashboardSection `xml:"rightSection"`
	RunningUser   *TextLiteral      `xml:"runningUser"`
	TextColor     TextLiteral       `xml:"textColor"`
	Title         TextLiteral       `xml:"title"`
	TitleColor    TextLiteral       `xml:"titleColor"`
	TitleSize     TextLiteral       `xml:"titleSize"`
}

// DashboardSection is a column of components in a dashboard that does not
// use a grid layout
type DashboardSection struct {
	ColumnSize TextLiteral          `xml:"columnSize"`
	Components []DashboardComponent `xml:"components"`
}

// DashboardComponent is a component of a dashboard, in either a grid layout
// or a section
type DashboardComponent struct {
	AutoselectColumnsFromReport *TextLiteral `xml:"autoselectColumnsFromReport"`
	ChartAxisRange              *TextLiteral `xml:"chartAxisRange"`
	ChartAxisRangeMax           *TextLiteral `xml:"chartAxisRangeMax"`
	ChartAxisRangeMin           *TextLiteral `xml:"chartAxisRangeMin"`
	ChartSummary                []struct {
		Aggregate   *TextLiteral `xml:"aggregate"`
		AxisBinding *TextLiteral `xml:"axisBinding"`
		Column      TextLiteral  `xml:"column"`
	} `xml:"chartSummary"`
	ComponentChartTheme        *TextLiteral `xml:"componentChartTheme"`
	ComponentType              TextLiteral  `xml:"componentType"`
	DashboardComponentContents *struct {
		RichTextContent TextLiteral `xml:"richTextContent"`
	} `xml:"dashboardComponentContents"`
	DashboardFilterColumns []struct {
		Column TextLiteral `xml:"column"`
	} `xml:"dashboardFilterColumns"`
	DashboardTableColumn []struct {
		AggregateType    *TextLiteral `xml:"aggregateType"`
		CalculatePercent *TextLiteral `xml:"calculatePercent"`
		Column           TextLiteral  `xml:"column"`
		ShowSubTotal     *TextLiteral `xml:"showSubTotal"`
		ShowTotal        *TextLiteral `xml:"showTotal"`
		SortBy           *TextLiteral `xml:"sortBy"`
	} `xml:"dashboardTableColumn"`
	DecimalPrecision        *TextLiteral `xml:"decimalPrecision"`
	DisplayUnits            *TextLiteral `xml:"displayUnits"`
	DrillDownUrl            *TextLiteral `xml:"drillDownUrl"`
	DrillEnabled            *TextLiteral `xml:"drillEnabled"`
	DrillToDetailEnabled    *TextLiteral `xml:"drillToDetailEnabled"`
	EnableHover             *TextLiteral `xml:"enableHover"`
	ExpandOthers            *TextLiteral `xml:"expandOthers"`
	FlexComponentProperties *struct {
		DecimalPrecision *TextLiteral `xml:"decimalPrecision"`
		FlexTableColumn  []struct {
			BreakPoint1     *TextLiteral `xml:"breakPoint1"`
			BreakPoint2     *TextLiteral `xml:"breakPoint2"`
			BreakPointOrder *TextLiteral `xml:"breakPointOrder"`
			HighRangeColor  *TextLiteral `xml:"highRangeColor"`
			LowRangeColor   *TextLiteral `xml:"lowRangeColor"`
			MidRangeColor   *TextLiteral `xml:"midRangeColor"`
			ReportColumn    NillableText `xml:"reportColumn"`
			ShowSubTotal    TextLiteral  `xml:"showSubTotal"`
			ShowTotal       TextLiteral  `xml:"showTotal"`
			Type            TextLiteral  `xml:"type"`
		} `xml:"flexTableColumn"`
		FlexTableSortInfo *struct {
			SortColumn *TextLiteral `xml:"sortColumn"`
			SortOrder  TextLiteral  `xml:"sortOrder"`
		} `xml:"flexTableSortInfo"`
		HideChatterPhotos     *TextLiteral `xml:"hideChatterPhotos"`
		UseReportTableSetting *TextLiteral `xml:"useReportTableSetting"`
	} `xml:"flexComponentProperties"`
	Footer                 *TextLiteral   `xml:"footer"`
	GaugeMax               *TextLiteral   `xml:"gaugeMax"`
	GaugeMin               *TextLiteral   `xml:"gaugeMin"`
	GroupingColumn         []NillableText `xml:"groupingColumn"`
	GroupingSortProperties *struct {
		GroupingSorts []struct {
			GroupingLevel               TextLiteral  `xml:"groupingLevel"`
			InheritedReportGroupingSort *TextLiteral `xml:"inheritedReportGroupingSort"`
			SortColumn                  *TextLiteral `xml:"sortColumn"`
			SortOrder                   *TextLiteral `xml:"sortOrder"`
		} `xml:"groupingSorts"`
	} `xml:"groupingSortProperties"`
	Header               *TextLiteral `xml:"header"`
	IndicatorBreakpoint1 *TextLiteral `xml:"indicatorBreakpoint1"`
	IndicatorBreakpoint2 *TextLiteral `xml:"indicatorBreakpoint2"`
	IndicatorHighColor   *TextLiteral `xml:"indicatorHighColor"`
	IndicatorLowColor    *TextLiteral `xml:"indicatorLowColor"`
	IndicatorMiddleColor *TextLiteral `xml:"indicatorMiddleColor"`
	LegendPosition       *TextLiteral `xml:"legendPosition"`
	MaxValuesDisplayed   *TextLiteral `xml:"maxValuesDisplayed"`
	MetricLabel          *TextLiteral `xml:"metricLabel"`
	Report               *string      `xml:"report"`
	ShowPercentage       *TextLiteral `xml:"showPercentage"`
	ShowPicturesOnCharts *TextLiteral `xml:"showPicturesOnCharts"`
	ShowPicturesOnTables *TextLiteral `xml:"showPicturesOnTables"`
	ShowRange            *TextLiteral `xml:"showRange"`
	ShowTotal            *TextLiteral `xml:"showTotal"`
	ShowValues           *TextLiteral `xml:"showValues"`
	SortBy               *TextLiteral `xml:"sortBy"`
	SortLegendValues     *TextLiteral `xml:"sortLegendValues"`
	Title                *TextLiteral `xml:"title"`
	UseReportChart       *TextLiteral `xml:"useReportChart"`
}

func (c *Dashboard) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Dashboard) Type() metadata.MetadataType {
	return NAME
}

func (c *Dashboard) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the original path from metadata info
	originalPath := string(c.MetadataInfo.Path())

	// Extract the folder structure from the original path
	// e.g., dashboards/HomePageMaster/Business_Development_Individual21.dashboard-meta.xml
	// Should preserve: HomePageMaster/Business_Development_Individual21

	// Get the directory name for dashboards
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Get relative path within the dashboards directory
	var relativePath string
	if within, ok := metadata.PathWithinDirectory(originalPath, "dashboards"); ok {
		relativePath = within
	}

	if relativePath == "" {
		return nil, fmt.Errorf("could not extract dashboard path from %s", originalPath)
	}

	// Remove the file extension and -meta.xml suffix to get the clean relative path
	relativePath = strings.TrimSuffix(relativePath, "-meta.xml")
	relativePath = strings.TrimSuffix(relativePath, ".dashboard")

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal dashboard metadata: %w", err)
	}

	files := make(map[string][]byte)

	var fileName string
	switch format {
	case metadata.SourceFormat:
		// Source format: preserve folder structure and add -meta.xml suffix
		fileName = relativePath + ".dashboard-meta.xml"
	case metadata.MetadataFormat:
		// Metadata format: preserve folder structure, no -meta.xml suffix
		fileName = relativePath + ".dashboard"
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	files[filepath.Join(dirName, fileName)] = xmlContent

	return files, nil
}

func Open(path string) (*Dashboard, error) {
	p := &Dashboard{}
	return p, metadata.ParseMetadataXml(p, path)
}
