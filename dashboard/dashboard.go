package dashboard

import (
	"encoding/xml"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
)

type Dashboard struct {
	XMLName            xml.Name `xml:"Dashboard"`
	Xmlns              string   `xml:"xmlns,attr"`
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
					DecimalPrecision struct {
						Text string `xml:",chardata"`
					} `xml:"decimalPrecision"`
					FlexTableColumn []struct {
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
				GroupingColumn *struct {
					Text string `xml:",chardata"`
				} `xml:"groupingColumn"`
				GaugeMax *struct {
					Text string `xml:",chardata"`
				} `xml:"gaugeMax"`
				GaugeMin *struct {
					Text string `xml:",chardata"`
				} `xml:"gaugeMin"`
				GroupingSortProperties struct {
					GroupingSorts *struct {
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
				Report struct {
					Text string `xml:",chardata"`
				} `xml:"report"`
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
	DashboardType struct {
		Text string `xml:",chardata"`
	} `xml:"dashboardType"`
	Description  *TextLiteral `xml:"description"`
	IsGridLayout struct {
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
			Report struct {
				Text string `xml:",chardata"`
			} `xml:"report"`
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
			Report struct {
				Text string `xml:",chardata"`
			} `xml:"report"`
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
			Report struct {
				Text string `xml:",chardata"`
			} `xml:"report"`
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

func (p *Dashboard) MetaCheck() {}

func Open(path string) (*Dashboard, error) {
	p := &Dashboard{}
	return p, internal.ParseMetadataXml(p, path)
}
