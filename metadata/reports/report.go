package report

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Report"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Report struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"Report"`
	Xmlns   string   `xml:"xmlns,attr"`
	Columns []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		AggregateTypes struct {
			Text string `xml:",chardata"`
		} `xml:"aggregateTypes"`
	} `xml:"columns"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Filter struct {
		CriteriaItems []struct {
			Column struct {
				Text string `xml:",chardata"`
			} `xml:"column"`
			ColumnToColumn struct {
				Text string `xml:",chardata"`
			} `xml:"columnToColumn"`
			IsUnlocked struct {
				Text string `xml:",chardata"`
			} `xml:"isUnlocked"`
			Operator struct {
				Text string `xml:",chardata"`
			} `xml:"operator"`
			Value struct {
				Text string `xml:",chardata"`
			} `xml:"value"`
		} `xml:"criteriaItems"`
		BooleanFilter struct {
			Text string `xml:",chardata"`
		} `xml:"booleanFilter"`
	} `xml:"filter"`
	Format struct {
		Text string `xml:",chardata"`
	} `xml:"format"`
	GroupingsDown []struct {
		DateGranularity struct {
			Text string `xml:",chardata"`
		} `xml:"dateGranularity"`
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		SortOrder struct {
			Text string `xml:",chardata"`
		} `xml:"sortOrder"`
		AggregateType struct {
			Text string `xml:",chardata"`
		} `xml:"aggregateType"`
		SortByName struct {
			Text string `xml:",chardata"`
		} `xml:"sortByName"`
		SortType struct {
			Text string `xml:",chardata"`
		} `xml:"sortType"`
	} `xml:"groupingsDown"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Params []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value struct {
			Text string `xml:",chardata"`
		} `xml:"value"`
	} `xml:"params"`
	ReportType struct {
		Text string `xml:",chardata"`
	} `xml:"reportType"`
	Scope struct {
		Text string `xml:",chardata"`
	} `xml:"scope"`
	ShowDetails struct {
		Text string `xml:",chardata"`
	} `xml:"showDetails"`
	ShowGrandTotal struct {
		Text string `xml:",chardata"`
	} `xml:"showGrandTotal"`
	ShowSubTotals struct {
		Text string `xml:",chardata"`
	} `xml:"showSubTotals"`
	TimeFrameFilter struct {
		DateColumn struct {
			Text string `xml:",chardata"`
		} `xml:"dateColumn"`
		Interval struct {
			Text string `xml:",chardata"`
		} `xml:"interval"`
		StartDate struct {
			Text string `xml:",chardata"`
		} `xml:"startDate"`
	} `xml:"timeFrameFilter"`
	SortColumn struct {
		Text string `xml:",chardata"`
	} `xml:"sortColumn"`
	SortOrder struct {
		Text string `xml:",chardata"`
	} `xml:"sortOrder"`
	Chart struct {
		BackgroundColor1 struct {
			Text string `xml:",chardata"`
		} `xml:"backgroundColor1"`
		BackgroundColor2 struct {
			Text string `xml:",chardata"`
		} `xml:"backgroundColor2"`
		BackgroundFadeDir struct {
			Text string `xml:",chardata"`
		} `xml:"backgroundFadeDir"`
		ChartSummaries struct {
			AxisBinding struct {
				Text string `xml:",chardata"`
			} `xml:"axisBinding"`
			Column struct {
				Text string `xml:",chardata"`
			} `xml:"column"`
			Aggregate struct {
				Text string `xml:",chardata"`
			} `xml:"aggregate"`
		} `xml:"chartSummaries"`
		ChartType struct {
			Text string `xml:",chardata"`
		} `xml:"chartType"`
		EnableHoverLabels struct {
			Text string `xml:",chardata"`
		} `xml:"enableHoverLabels"`
		ExpandOthers struct {
			Text string `xml:",chardata"`
		} `xml:"expandOthers"`
		GroupingColumn struct {
			Text string `xml:",chardata"`
		} `xml:"groupingColumn"`
		Location struct {
			Text string `xml:",chardata"`
		} `xml:"location"`
		ShowAxisLabels struct {
			Text string `xml:",chardata"`
		} `xml:"showAxisLabels"`
		ShowPercentage struct {
			Text string `xml:",chardata"`
		} `xml:"showPercentage"`
		ShowTotal struct {
			Text string `xml:",chardata"`
		} `xml:"showTotal"`
		ShowValues struct {
			Text string `xml:",chardata"`
		} `xml:"showValues"`
		Size struct {
			Text string `xml:",chardata"`
		} `xml:"size"`
		SummaryAxisRange struct {
			Text string `xml:",chardata"`
		} `xml:"summaryAxisRange"`
		TextColor struct {
			Text string `xml:",chardata"`
		} `xml:"textColor"`
		TextSize struct {
			Text string `xml:",chardata"`
		} `xml:"textSize"`
		TitleColor struct {
			Text string `xml:",chardata"`
		} `xml:"titleColor"`
		TitleSize struct {
			Text string `xml:",chardata"`
		} `xml:"titleSize"`
		Title struct {
			Text string `xml:",chardata"`
		} `xml:"title"`
		LegendPosition struct {
			Text string `xml:",chardata"`
		} `xml:"legendPosition"`
		SecondaryGroupingColumn struct {
			Text string `xml:",chardata"`
		} `xml:"secondaryGroupingColumn"`
	} `xml:"chart"`
	GroupingsAcross []struct {
		DateGranularity struct {
			Text string `xml:",chardata"`
		} `xml:"dateGranularity"`
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		SortOrder struct {
			Text string `xml:",chardata"`
		} `xml:"sortOrder"`
	} `xml:"groupingsAcross"`
	Buckets struct {
		BucketType struct {
			Text string `xml:",chardata"`
		} `xml:"bucketType"`
		DeveloperName struct {
			Text string `xml:",chardata"`
		} `xml:"developerName"`
		MasterLabel struct {
			Text string `xml:",chardata"`
		} `xml:"masterLabel"`
		NullTreatment struct {
			Text string `xml:",chardata"`
		} `xml:"nullTreatment"`
		SourceColumnName struct {
			Text string `xml:",chardata"`
		} `xml:"sourceColumnName"`
		UseOther struct {
			Text string `xml:",chardata"`
		} `xml:"useOther"`
		Values []struct {
			SourceValues []struct {
				SourceValue struct {
					Text string `xml:",chardata"`
				} `xml:"sourceValue"`
			} `xml:"sourceValues"`
			Value struct {
				Text string `xml:",chardata"`
			} `xml:"value"`
		} `xml:"values"`
		OtherBucketLabel struct {
			Text string `xml:",chardata"`
		} `xml:"otherBucketLabel"`
	} `xml:"buckets"`
	CustomDetailFormulas struct {
		CalculatedFormula struct {
			Text string `xml:",chardata"`
		} `xml:"calculatedFormula"`
		DataType struct {
			Text string `xml:",chardata"`
		} `xml:"dataType"`
		DeveloperName struct {
			Text string `xml:",chardata"`
		} `xml:"developerName"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		Scale struct {
			Text string `xml:",chardata"`
		} `xml:"scale"`
	} `xml:"customDetailFormulas"`
	CrossFilters struct {
		Operation struct {
			Text string `xml:",chardata"`
		} `xml:"operation"`
		PrimaryTableColumn struct {
			Text string `xml:",chardata"`
		} `xml:"primaryTableColumn"`
		RelatedTable struct {
			Text string `xml:",chardata"`
		} `xml:"relatedTable"`
		RelatedTableJoinColumn struct {
			Text string `xml:",chardata"`
		} `xml:"relatedTableJoinColumn"`
	} `xml:"crossFilters"`
	Aggregates []struct {
		CalculatedFormula struct {
			Text string `xml:",chardata"`
		} `xml:"calculatedFormula"`
		Datatype struct {
			Text string `xml:",chardata"`
		} `xml:"datatype"`
		DeveloperName struct {
			Text string `xml:",chardata"`
		} `xml:"developerName"`
		IsActive struct {
			Text string `xml:",chardata"`
		} `xml:"isActive"`
		IsCrossBlock struct {
			Text string `xml:",chardata"`
		} `xml:"isCrossBlock"`
		MasterLabel struct {
			Text string `xml:",chardata"`
		} `xml:"masterLabel"`
		Scale struct {
			Text string `xml:",chardata"`
		} `xml:"scale"`
		AcrossGroupingContext struct {
			Text string `xml:",chardata"`
		} `xml:"acrossGroupingContext"`
		DownGroupingContext struct {
			Text string `xml:",chardata"`
		} `xml:"downGroupingContext"`
	} `xml:"aggregates"`
	ColorRanges struct {
		ColumnName struct {
			Text string `xml:",chardata"`
		} `xml:"columnName"`
		HighColor struct {
			Text string `xml:",chardata"`
		} `xml:"highColor"`
		LowBreakpoint struct {
			Text string `xml:",chardata"`
		} `xml:"lowBreakpoint"`
		LowColor struct {
			Text string `xml:",chardata"`
		} `xml:"lowColor"`
		MidColor struct {
			Text string `xml:",chardata"`
		} `xml:"midColor"`
		HighBreakpoint struct {
			Text string `xml:",chardata"`
		} `xml:"highBreakpoint"`
	} `xml:"colorRanges"`
	FormattingRules struct {
		ColumnName struct {
			Text string `xml:",chardata"`
		} `xml:"columnName"`
		Values []struct {
			RangeUpperBound struct {
				Text string `xml:",chardata"`
			} `xml:"rangeUpperBound"`
			BackgroundColor struct {
				Text string `xml:",chardata"`
			} `xml:"backgroundColor"`
		} `xml:"values"`
	} `xml:"formattingRules"`
	RoleHierarchyFilter struct {
		Text string `xml:",chardata"`
	} `xml:"roleHierarchyFilter"`
}

func (c *Report) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Report) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Report, error) {
	p := &Report{}
	return p, metadata.ParseMetadataXml(p, path)
}
