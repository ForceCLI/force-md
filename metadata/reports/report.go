package report

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

const NAME = "Report"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Report struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"Report"`
	Xmlns      string   `xml:"xmlns,attr"`
	Aggregates []struct {
		AcrossGroupingContext *struct {
			Text string `xml:",chardata"`
		} `xml:"acrossGroupingContext"`
		CalculatedFormula *TextLiteral `xml:"calculatedFormula"`
		Datatype          struct {
			Text string `xml:",chardata"`
		} `xml:"datatype"`
		DeveloperName struct {
			Text string `xml:",chardata"`
		} `xml:"developerName"`
		DownGroupingContext *struct {
			Text string `xml:",chardata"`
		} `xml:"downGroupingContext"`
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
	} `xml:"aggregates"`
	Buckets []struct {
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
		OtherBucketLabel *struct {
			Text string `xml:",chardata"`
		} `xml:"otherBucketLabel"`
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
	} `xml:"buckets"`
	Chart *struct {
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
			Aggregate *struct {
				Text string `xml:",chardata"`
			} `xml:"aggregate"`
			AxisBinding struct {
				Text string `xml:",chardata"`
			} `xml:"axisBinding"`
			Column struct {
				Text string `xml:",chardata"`
			} `xml:"column"`
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
		LegendPosition *struct {
			Text string `xml:",chardata"`
		} `xml:"legendPosition"`
		Location struct {
			Text string `xml:",chardata"`
		} `xml:"location"`
		SecondaryGroupingColumn *struct {
			Text string `xml:",chardata"`
		} `xml:"secondaryGroupingColumn"`
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
		Title *struct {
			Text string `xml:",chardata"`
		} `xml:"title"`
		TitleColor struct {
			Text string `xml:",chardata"`
		} `xml:"titleColor"`
		TitleSize struct {
			Text string `xml:",chardata"`
		} `xml:"titleSize"`
	} `xml:"chart"`
	ColorRanges []struct {
		ColumnName struct {
			Text string `xml:",chardata"`
		} `xml:"columnName"`
		HighBreakpoint *struct {
			Text string `xml:",chardata"`
		} `xml:"highBreakpoint"`
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
	} `xml:"colorRanges"`
	Columns []struct {
		AggregateTypes *struct {
			Text string `xml:",chardata"`
		} `xml:"aggregateTypes"`
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
	} `xml:"columns"`
	CrossFilters []struct {
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
	CustomDetailFormulas []struct {
		CalculatedFormula struct {
			Text string `xml:",chardata"`
		} `xml:"calculatedFormula"`
		DataType struct {
			Text string `xml:",chardata"`
		} `xml:"dataType"`
		Description *struct {
			Text string `xml:",chardata"`
		} `xml:"description"`
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
	Description *TextLiteral `xml:"description"`
	Filter      *struct {
		BooleanFilter *struct {
			Text string `xml:",chardata"`
		} `xml:"booleanFilter"`
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
	} `xml:"filter"`
	Format struct {
		Text string `xml:",chardata"`
	} `xml:"format"`
	FormattingRules []struct {
		ColumnName struct {
			Text string `xml:",chardata"`
		} `xml:"columnName"`
		Values []struct {
			BackgroundColor *struct {
				Text string `xml:",chardata"`
			} `xml:"backgroundColor"`
			RangeUpperBound *struct {
				Text string `xml:",chardata"`
			} `xml:"rangeUpperBound"`
		} `xml:"values"`
	} `xml:"formattingRules"`
	GroupingsAcross []struct {
		DateGranularity *struct {
			Text string `xml:",chardata"`
		} `xml:"dateGranularity"`
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		SortOrder struct {
			Text string `xml:",chardata"`
		} `xml:"sortOrder"`
	} `xml:"groupingsAcross"`
	GroupingsDown []struct {
		AggregateType *struct {
			Text string `xml:",chardata"`
		} `xml:"aggregateType"`
		DateGranularity *struct {
			Text string `xml:",chardata"`
		} `xml:"dateGranularity"`
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		SortByName *struct {
			Text string `xml:",chardata"`
		} `xml:"sortByName"`
		SortOrder struct {
			Text string `xml:",chardata"`
		} `xml:"sortOrder"`
		SortType *struct {
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
	RoleHierarchyFilter []struct {
		Text string `xml:",chardata"`
	} `xml:"roleHierarchyFilter"`
	Scope *struct {
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
	SortColumn *struct {
		Text string `xml:",chardata"`
	} `xml:"sortColumn"`
	SortOrder *struct {
		Text string `xml:",chardata"`
	} `xml:"sortOrder"`
	TimeFrameFilter struct {
		DateColumn struct {
			Text string `xml:",chardata"`
		} `xml:"dateColumn"`
		Interval struct {
			Text string `xml:",chardata"`
		} `xml:"interval"`
		StartDate *struct {
			Text string `xml:",chardata"`
		} `xml:"startDate"`
	} `xml:"timeFrameFilter"`
}

func (c *Report) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Report) Type() metadata.MetadataType {
	return NAME
}

func (c *Report) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the original path from metadata info
	originalPath := string(c.MetadataInfo.Path())

	// Extract the folder structure from the original path
	// e.g., reports/CRM_Admin_Exception_Reports/SystemAdminExceptionDashboard/Accounts_For_Merging_DSj.report-meta.xml
	// Should preserve: CRM_Admin_Exception_Reports/SystemAdminExceptionDashboard/Accounts_For_Merging_DSj

	// Get the directory name for reports
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Get relative path within the reports directory
	var relativePath string
	if strings.Contains(originalPath, "reports/") {
		// Extract everything after "reports/"
		parts := strings.Split(originalPath, "reports/")
		if len(parts) > 1 {
			relativePath = parts[1]
		}
	}

	if relativePath == "" {
		return nil, fmt.Errorf("could not extract report path from %s", originalPath)
	}

	// Remove the file extension and -meta.xml suffix to get the clean relative path
	relativePath = strings.TrimSuffix(relativePath, "-meta.xml")
	relativePath = strings.TrimSuffix(relativePath, ".report")

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal report metadata: %w", err)
	}

	files := make(map[string][]byte)

	var fileName string
	switch format {
	case metadata.SourceFormat:
		// Source format: preserve folder structure and add -meta.xml suffix
		fileName = relativePath + ".report-meta.xml"
	case metadata.MetadataFormat:
		// Metadata format: preserve folder structure, no -meta.xml suffix
		fileName = relativePath + ".report"
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	files[filepath.Join(dirName, fileName)] = xmlContent

	return files, nil
}

func Open(path string) (*Report, error) {
	p := &Report{}
	return p, metadata.ParseMetadataXml(p, path)
}
