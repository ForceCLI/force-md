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
	XMLName            xml.Name `xml:"Dashboard"`
	Xmlns              string   `xml:"xmlns,attr"`
	Xsi                *string  `xml:"xsi,attr"`
	BackgroundEndColor struct {
		Text string `xml:",chardata"`
	} `xml:"backgroundEndColor"`
	BackgroundFadeDirection struct {
		Text string `xml:",chardata"`
	} `xml:"backgroundFadeDirection"`
	BackgroundStartColor struct {
		Text string `xml:",chardata"`
	} `xml:"backgroundStartColor"`
	ChartTheme *struct {
		Text string `xml:",chardata"`
	} `xml:"chartTheme"`
	ColorPalette *struct {
		Text string `xml:",chardata"`
	} `xml:"colorPalette"`
	DashboardChartTheme *struct {
		Text string `xml:",chardata"`
	} `xml:"dashboardChartTheme"`
	DashboardColorPalette *struct {
		Text string `xml:",chardata"`
	} `xml:"dashboardColorPalette"`
	DashboardFilters []struct {
		DashboardFilterOptions []struct {
			Operator struct {
				Text string `xml:",chardata"`
			} `xml:"operator"`
			Values struct {
				Text string `xml:",chardata"`
			} `xml:"values"`
		} `xml:"dashboardFilterOptions"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"dashboardFilters"`
	DashboardGridLayout *struct {
		DashboardGridComponents []struct {
			ColSpan struct {
				Text string `xml:",chardata"`
			} `xml:"colSpan"`
			ColumnIndex struct {
				Text string `xml:",chardata"`
			} `xml:"columnIndex"`
			DashboardComponent struct {
				AutoselectColumnsFromReport struct {
					Text string `xml:",chardata"`
				} `xml:"autoselectColumnsFromReport"`
				ChartAxisRange *struct {
					Text string `xml:",chardata"`
				} `xml:"chartAxisRange"`
				ChartAxisRangeMax *struct {
					Text string `xml:",chardata"`
				} `xml:"chartAxisRangeMax"`
				ChartSummary *struct {
					Aggregate *struct {
						Text string `xml:",chardata"`
					} `xml:"aggregate"`
					AxisBinding *struct {
						Text string `xml:",chardata"`
					} `xml:"axisBinding"`
					Column struct {
						Text string `xml:",chardata"`
					} `xml:"column"`
				} `xml:"chartSummary"`
				ComponentType struct {
					Text string `xml:",chardata"`
				} `xml:"componentType"`
				DashboardComponentContents *struct {
					RichTextContent struct {
						Text string `xml:",chardata"`
					} `xml:"richTextContent"`
				} `xml:"dashboardComponentContents"`
				DashboardFilterColumns []struct {
					Column struct {
						Text string `xml:",chardata"`
					} `xml:"column"`
				} `xml:"dashboardFilterColumns"`
				DecimalPrecision *struct {
					Text string `xml:",chardata"`
				} `xml:"decimalPrecision"`
				DisplayUnits *struct {
					Text string `xml:",chardata"`
				} `xml:"displayUnits"`
				DrillEnabled *struct {
					Text string `xml:",chardata"`
				} `xml:"drillEnabled"`
				DrillToDetailEnabled *struct {
					Text string `xml:",chardata"`
				} `xml:"drillToDetailEnabled"`
				EnableHover *struct {
					Text string `xml:",chardata"`
				} `xml:"enableHover"`
				ExpandOthers *struct {
					Text string `xml:",chardata"`
				} `xml:"expandOthers"`
				FlexComponentProperties *struct {
					DecimalPrecision *struct {
						Text string `xml:",chardata"`
					} `xml:"decimalPrecision"`
					FlexTableColumn []struct {
						BreakPoint1 *struct {
							Text string `xml:",chardata"`
						} `xml:"breakPoint1"`
						BreakPoint2 *struct {
							Text string `xml:",chardata"`
						} `xml:"breakPoint2"`
						BreakPointOrder *struct {
							Text string `xml:",chardata"`
						} `xml:"breakPointOrder"`
						HighRangeColor *struct {
							Text string `xml:",chardata"`
						} `xml:"highRangeColor"`
						LowRangeColor *struct {
							Text string `xml:",chardata"`
						} `xml:"lowRangeColor"`
						MidRangeColor *struct {
							Text string `xml:",chardata"`
						} `xml:"midRangeColor"`
						ReportColumn struct {
							Text string `xml:",chardata"`
						} `xml:"reportColumn"`
						ShowSubTotal struct {
							Text string `xml:",chardata"`
						} `xml:"showSubTotal"`
						ShowTotal struct {
							Text string `xml:",chardata"`
						} `xml:"showTotal"`
						Type struct {
							Text string `xml:",chardata"`
						} `xml:"type"`
					} `xml:"flexTableColumn"`
					FlexTableSortInfo struct {
						SortColumn *struct {
							Text string `xml:",chardata"`
						} `xml:"sortColumn"`
						SortOrder struct {
							Text string `xml:",chardata"`
						} `xml:"sortOrder"`
					} `xml:"flexTableSortInfo"`
					HideChatterPhotos struct {
						Text string `xml:",chardata"`
					} `xml:"hideChatterPhotos"`
				} `xml:"flexComponentProperties"`
				Footer *struct {
					Text string `xml:",chardata"`
				} `xml:"footer"`
				GroupingColumn []struct {
					Text string `xml:",chardata"`
				} `xml:"groupingColumn"`
				GaugeMax *struct {
					Text string `xml:",chardata"`
				} `xml:"gaugeMax"`
				GaugeMin *struct {
					Text string `xml:",chardata"`
				} `xml:"gaugeMin"`
				GroupingSortProperties struct {
					GroupingSorts []struct {
						GroupingLevel struct {
							Text string `xml:",chardata"`
						} `xml:"groupingLevel"`
						SortColumn *struct {
							Text string `xml:",chardata"`
						} `xml:"sortColumn"`
						SortOrder *struct {
							Text string `xml:",chardata"`
						} `xml:"sortOrder"`
						InheritedReportGroupingSort *struct {
							Text string `xml:",chardata"`
						} `xml:"inheritedReportGroupingSort"`
					} `xml:"groupingSorts"`
				} `xml:"groupingSortProperties"`
				Header *struct {
					Text string `xml:",chardata"`
				} `xml:"header"`
				IndicatorBreakpoint1 *struct {
					Text string `xml:",chardata"`
				} `xml:"indicatorBreakpoint1"`
				IndicatorBreakpoint2 *struct {
					Text string `xml:",chardata"`
				} `xml:"indicatorBreakpoint2"`
				IndicatorHighColor *struct {
					Text string `xml:",chardata"`
				} `xml:"indicatorHighColor"`
				IndicatorLowColor *struct {
					Text string `xml:",chardata"`
				} `xml:"indicatorLowColor"`
				IndicatorMiddleColor *struct {
					Text string `xml:",chardata"`
				} `xml:"indicatorMiddleColor"`
				LegendPosition *struct {
					Text string `xml:",chardata"`
				} `xml:"legendPosition"`
				MaxValuesDisplayed *struct {
					Text string `xml:",chardata"`
				} `xml:"maxValuesDisplayed"`
				MetricLabel *struct {
					Text string `xml:",chardata"`
				} `xml:"metricLabel"`
				Report         *string `xml:"report"`
				ShowPercentage *struct {
					Text string `xml:",chardata"`
				} `xml:"showPercentage"`
				ShowPicturesOnCharts *struct {
					Text string `xml:",chardata"`
				} `xml:"showPicturesOnCharts"`
				ShowRange *struct {
					Text string `xml:",chardata"`
				} `xml:"showRange"`
				ShowTotal *struct {
					Text string `xml:",chardata"`
				} `xml:"showTotal"`
				ShowValues *struct {
					Text string `xml:",chardata"`
				} `xml:"showValues"`
				SortBy *struct {
					Text string `xml:",chardata"`
				} `xml:"sortBy"`
				Title *struct {
					Text string `xml:",chardata"`
				} `xml:"title"`
				UseReportChart *struct {
					Text string `xml:",chardata"`
				} `xml:"useReportChart"`
			} `xml:"dashboardComponent"`
			RowIndex struct {
				Text string `xml:",chardata"`
			} `xml:"rowIndex"`
			RowSpan struct {
				Text string `xml:",chardata"`
			} `xml:"rowSpan"`
		} `xml:"dashboardGridComponents"`
		NumberOfColumns struct {
			Text string `xml:",chardata"`
		} `xml:"numberOfColumns"`
		RowHeight struct {
			Text string `xml:",chardata"`
		} `xml:"rowHeight"`
	} `xml:"dashboardGridLayout"`
	DashboardType *TextLiteral `xml:"dashboardType"`
	Description   *TextLiteral `xml:"description"`
	IsGridLayout  struct {
		Text string `xml:",chardata"`
	} `xml:"isGridLayout"`
	LeftSection *struct {
		ColumnSize struct {
			Text string `xml:",chardata"`
		} `xml:"columnSize"`
		Components []struct {
			AutoselectColumnsFromReport struct {
				Text string `xml:",chardata"`
			} `xml:"autoselectColumnsFromReport"`
			ChartAxisRange *struct {
				Text string `xml:",chardata"`
			} `xml:"chartAxisRange"`
			ComponentType struct {
				Text string `xml:",chardata"`
			} `xml:"componentType"`
			DashboardFilterColumns []struct {
				Column struct {
					Text string `xml:",chardata"`
				} `xml:"column"`
			} `xml:"dashboardFilterColumns"`
			DashboardTableColumn []struct {
				AggregateType *struct {
					Text string `xml:",chardata"`
				} `xml:"aggregateType"`
				CalculatePercent *struct {
					Text string `xml:",chardata"`
				} `xml:"calculatePercent"`
				Column struct {
					Text string `xml:",chardata"`
				} `xml:"column"`
				ShowTotal *struct {
					Text string `xml:",chardata"`
				} `xml:"showTotal"`
				SortBy *struct {
					Text string `xml:",chardata"`
				} `xml:"sortBy"`
			} `xml:"dashboardTableColumn"`
			DisplayUnits struct {
				Text string `xml:",chardata"`
			} `xml:"displayUnits"`
			DrillEnabled *struct {
				Text string `xml:",chardata"`
			} `xml:"drillEnabled"`
			DrillToDetailEnabled *struct {
				Text string `xml:",chardata"`
			} `xml:"drillToDetailEnabled"`
			EnableHover *struct {
				Text string `xml:",chardata"`
			} `xml:"enableHover"`
			ExpandOthers *struct {
				Text string `xml:",chardata"`
			} `xml:"expandOthers"`
			Footer *struct {
				Text string `xml:",chardata"`
			} `xml:"footer"`
			GaugeMax *struct {
				Text string `xml:",chardata"`
			} `xml:"gaugeMax"`
			GaugeMin *struct {
				Text string `xml:",chardata"`
			} `xml:"gaugeMin"`
			GroupingSortProperties struct {
			} `xml:"groupingSortProperties"`
			Header *struct {
				Text string `xml:",chardata"`
			} `xml:"header"`
			IndicatorBreakpoint1 *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorBreakpoint1"`
			IndicatorBreakpoint2 *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorBreakpoint2"`
			IndicatorHighColor *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorHighColor"`
			IndicatorLowColor *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorLowColor"`
			IndicatorMiddleColor *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorMiddleColor"`
			LegendPosition *struct {
				Text string `xml:",chardata"`
			} `xml:"legendPosition"`
			MaxValuesDisplayed *struct {
				Text string `xml:",chardata"`
			} `xml:"maxValuesDisplayed"`
			MetricLabel *struct {
				Text string `xml:",chardata"`
			} `xml:"metricLabel"`
			Report         string `xml:"report"`
			ShowPercentage *struct {
				Text string `xml:",chardata"`
			} `xml:"showPercentage"`
			ShowRange *struct {
				Text string `xml:",chardata"`
			} `xml:"showRange"`
			ShowTotal *struct {
				Text string `xml:",chardata"`
			} `xml:"showTotal"`
			ShowValues *struct {
				Text string `xml:",chardata"`
			} `xml:"showValues"`
			SortBy *struct {
				Text string `xml:",chardata"`
			} `xml:"sortBy"`
			Title *struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			UseReportChart *struct {
				Text string `xml:",chardata"`
			} `xml:"useReportChart"`
		} `xml:"components"`
	} `xml:"leftSection"`
	MiddleSection *struct {
		ColumnSize struct {
			Text string `xml:",chardata"`
		} `xml:"columnSize"`
		Components []struct {
			AutoselectColumnsFromReport struct {
				Text string `xml:",chardata"`
			} `xml:"autoselectColumnsFromReport"`
			ChartAxisRange *struct {
				Text string `xml:",chardata"`
			} `xml:"chartAxisRange"`
			ChartSummary *struct {
				AxisBinding struct {
					Text string `xml:",chardata"`
				} `xml:"axisBinding"`
				Column struct {
					Text string `xml:",chardata"`
				} `xml:"column"`
			} `xml:"chartSummary"`
			ComponentType struct {
				Text string `xml:",chardata"`
			} `xml:"componentType"`
			DashboardFilterColumns []struct {
				Column struct {
					Text string `xml:",chardata"`
				} `xml:"column"`
			} `xml:"dashboardFilterColumns"`
			DashboardTableColumn []struct {
				CalculatePercent *struct {
					Text string `xml:",chardata"`
				} `xml:"calculatePercent"`
				Column struct {
					Text string `xml:",chardata"`
				} `xml:"column"`
				ShowTotal *struct {
					Text string `xml:",chardata"`
				} `xml:"showTotal"`
				SortBy *struct {
					Text string `xml:",chardata"`
				} `xml:"sortBy"`
			} `xml:"dashboardTableColumn"`
			DisplayUnits struct {
				Text string `xml:",chardata"`
			} `xml:"displayUnits"`
			DrillDownUrl *struct {
				Text string `xml:",chardata"`
			} `xml:"drillDownUrl"`
			DrillEnabled *struct {
				Text string `xml:",chardata"`
			} `xml:"drillEnabled"`
			DrillToDetailEnabled *struct {
				Text string `xml:",chardata"`
			} `xml:"drillToDetailEnabled"`
			EnableHover *struct {
				Text string `xml:",chardata"`
			} `xml:"enableHover"`
			ExpandOthers *struct {
				Text string `xml:",chardata"`
			} `xml:"expandOthers"`
			Footer *struct {
				Text string `xml:",chardata"`
			} `xml:"footer"`
			GaugeMax *struct {
				Text string `xml:",chardata"`
			} `xml:"gaugeMax"`
			GaugeMin *struct {
				Text string `xml:",chardata"`
			} `xml:"gaugeMin"`
			GroupingColumn []struct {
				Text string `xml:",chardata"`
			} `xml:"groupingColumn"`
			GroupingSortProperties struct {
			} `xml:"groupingSortProperties"`
			Header *struct {
				Text string `xml:",chardata"`
			} `xml:"header"`
			IndicatorBreakpoint1 *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorBreakpoint1"`
			IndicatorBreakpoint2 *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorBreakpoint2"`
			IndicatorHighColor *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorHighColor"`
			IndicatorLowColor *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorLowColor"`
			IndicatorMiddleColor *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorMiddleColor"`
			LegendPosition *struct {
				Text string `xml:",chardata"`
			} `xml:"legendPosition"`
			MaxValuesDisplayed *struct {
				Text string `xml:",chardata"`
			} `xml:"maxValuesDisplayed"`
			MetricLabel *struct {
				Text string `xml:",chardata"`
			} `xml:"metricLabel"`
			Report         string `xml:"report"`
			ShowPercentage *struct {
				Text string `xml:",chardata"`
			} `xml:"showPercentage"`
			ShowRange *struct {
				Text string `xml:",chardata"`
			} `xml:"showRange"`
			ShowTotal *struct {
				Text string `xml:",chardata"`
			} `xml:"showTotal"`
			ShowValues *struct {
				Text string `xml:",chardata"`
			} `xml:"showValues"`
			SortBy *struct {
				Text string `xml:",chardata"`
			} `xml:"sortBy"`
			Title *struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			UseReportChart *struct {
				Text string `xml:",chardata"`
			} `xml:"useReportChart"`
		} `xml:"components"`
	} `xml:"middleSection"`
	RightSection *struct {
		ColumnSize struct {
			Text string `xml:",chardata"`
		} `xml:"columnSize"`
		Components []struct {
			AutoselectColumnsFromReport struct {
				Text string `xml:",chardata"`
			} `xml:"autoselectColumnsFromReport"`
			ChartAxisRange *struct {
				Text string `xml:",chardata"`
			} `xml:"chartAxisRange"`
			ComponentType struct {
				Text string `xml:",chardata"`
			} `xml:"componentType"`
			DashboardFilterColumns []struct {
				Column struct {
					Text string `xml:",chardata"`
				} `xml:"column"`
			} `xml:"dashboardFilterColumns"`
			DashboardTableColumn []struct {
				AggregateType *struct {
					Text string `xml:",chardata"`
				} `xml:"aggregateType"`
				CalculatePercent *struct {
					Text string `xml:",chardata"`
				} `xml:"calculatePercent"`
				Column struct {
					Text string `xml:",chardata"`
				} `xml:"column"`
				ShowTotal *struct {
					Text string `xml:",chardata"`
				} `xml:"showTotal"`
				SortBy *struct {
					Text string `xml:",chardata"`
				} `xml:"sortBy"`
			} `xml:"dashboardTableColumn"`
			DisplayUnits struct {
				Text string `xml:",chardata"`
			} `xml:"displayUnits"`
			DrillEnabled *struct {
				Text string `xml:",chardata"`
			} `xml:"drillEnabled"`
			DrillToDetailEnabled *struct {
				Text string `xml:",chardata"`
			} `xml:"drillToDetailEnabled"`
			EnableHover *struct {
				Text string `xml:",chardata"`
			} `xml:"enableHover"`
			ExpandOthers *struct {
				Text string `xml:",chardata"`
			} `xml:"expandOthers"`
			Footer *struct {
				Text string `xml:",chardata"`
			} `xml:"footer"`
			GaugeMax *struct {
				Text string `xml:",chardata"`
			} `xml:"gaugeMax"`
			GaugeMin *struct {
				Text string `xml:",chardata"`
			} `xml:"gaugeMin"`
			GroupingSortProperties struct {
			} `xml:"groupingSortProperties"`
			Header *struct {
				Text string `xml:",chardata"`
			} `xml:"header"`
			IndicatorBreakpoint1 *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorBreakpoint1"`
			IndicatorBreakpoint2 *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorBreakpoint2"`
			IndicatorHighColor *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorHighColor"`
			IndicatorLowColor *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorLowColor"`
			IndicatorMiddleColor *struct {
				Text string `xml:",chardata"`
			} `xml:"indicatorMiddleColor"`
			LegendPosition *struct {
				Text string `xml:",chardata"`
			} `xml:"legendPosition"`
			MaxValuesDisplayed *struct {
				Text string `xml:",chardata"`
			} `xml:"maxValuesDisplayed"`
			MetricLabel *struct {
				Text string `xml:",chardata"`
			} `xml:"metricLabel"`
			Report         string `xml:"report"`
			ShowPercentage *struct {
				Text string `xml:",chardata"`
			} `xml:"showPercentage"`
			ShowRange *struct {
				Text string `xml:",chardata"`
			} `xml:"showRange"`
			ShowTotal *struct {
				Text string `xml:",chardata"`
			} `xml:"showTotal"`
			ShowValues *struct {
				Text string `xml:",chardata"`
			} `xml:"showValues"`
			SortBy *struct {
				Text string `xml:",chardata"`
			} `xml:"sortBy"`
			Title *struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			UseReportChart *struct {
				Text string `xml:",chardata"`
			} `xml:"useReportChart"`
		} `xml:"components"`
	} `xml:"rightSection"`
	Owner       *TextLiteral `xml:"owner"`
	RunningUser *TextLiteral `xml:"runningUser"`
	TextColor   struct {
		Text string `xml:",chardata"`
	} `xml:"textColor"`
	Title struct {
		Text string `xml:",chardata"`
	} `xml:"title"`
	TitleColor struct {
		Text string `xml:",chardata"`
	} `xml:"titleColor"`
	TitleSize struct {
		Text string `xml:",chardata"`
	} `xml:"titleSize"`
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
	if strings.Contains(originalPath, "dashboards/") {
		// Extract everything after "dashboards/"
		parts := strings.Split(originalPath, "dashboards/")
		if len(parts) > 1 {
			relativePath = parts[1]
		}
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
