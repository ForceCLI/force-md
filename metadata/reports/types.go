package report

import (
	. "github.com/ForceCLI/force-md/general"
)

// ReportColumn represents a column in a report
type ReportColumn struct {
	AggregateTypes []TextLiteral `xml:"aggregateTypes"`
	Field          TextLiteral   `xml:"field"`
}

// ReportGroupingAcross represents a grouping across in a report
type ReportGroupingAcross struct {
	DateGranularity *TextLiteral `xml:"dateGranularity"`
	Field           TextLiteral  `xml:"field"`
	SortOrder       TextLiteral  `xml:"sortOrder"`
}

// ReportGroupingDown represents a grouping down in a report
type ReportGroupingDown struct {
	AggregateType   *TextLiteral `xml:"aggregateType"`
	DateGranularity *TextLiteral `xml:"dateGranularity"`
	Field           TextLiteral  `xml:"field"`
	SortByName      *TextLiteral `xml:"sortByName"`
	SortOrder       TextLiteral  `xml:"sortOrder"`
	SortType        *TextLiteral `xml:"sortType"`
}

// ReportFilterCriteriaItem represents a filter criteria item in a report
type ReportFilterCriteriaItem struct {
	Column         TextLiteral  `xml:"column"`
	ColumnToColumn *TextLiteral `xml:"columnToColumn"`
	IsUnlocked     *TextLiteral `xml:"isUnlocked"`
	Operator       TextLiteral  `xml:"operator"`
	Value          TextLiteral  `xml:"value"`
}

// ReportFilter represents a filter in a report
type ReportFilter struct {
	BooleanFilter *TextLiteral               `xml:"booleanFilter"`
	CriteriaItems []ReportFilterCriteriaItem `xml:"criteriaItems"`
}

// ReportBucketSourceValue is one source of a bucket value: a picklist or text
// value, or a numeric range bounded by from and to
type ReportBucketSourceValue struct {
	From        *TextLiteral `xml:"from"`
	SourceValue *TextLiteral `xml:"sourceValue"`
	To          *TextLiteral `xml:"to"`
}

// ReportBucketValue represents a bucket value
type ReportBucketValue struct {
	SourceValues []ReportBucketSourceValue `xml:"sourceValues"`
	Value        TextLiteral               `xml:"value"`
}

// ReportBucket represents a bucket in a report
type ReportBucket struct {
	BucketType       TextLiteral         `xml:"bucketType"`
	DeveloperName    TextLiteral         `xml:"developerName"`
	MasterLabel      TextLiteral         `xml:"masterLabel"`
	NullTreatment    TextLiteral         `xml:"nullTreatment"`
	OtherBucketLabel *TextLiteral        `xml:"otherBucketLabel"`
	SourceColumnName TextLiteral         `xml:"sourceColumnName"`
	UseOther         TextLiteral         `xml:"useOther"`
	Values           []ReportBucketValue `xml:"values"`
}

// ReportParam represents a report parameter
type ReportParam struct {
	Name  TextLiteral `xml:"name"`
	Value TextLiteral `xml:"value"`
}

// ReportTimeFrameFilter represents the date range a report covers
type ReportTimeFrameFilter struct {
	DateColumn TextLiteral  `xml:"dateColumn"`
	EndDate    *TextLiteral `xml:"endDate"`
	Interval   TextLiteral  `xml:"interval"`
	StartDate  *TextLiteral `xml:"startDate"`
}

// ReportCrossFilter represents a cross filter, which filters a report by its
// related records
type ReportCrossFilter struct {
	CriteriaItems          []ReportFilterCriteriaItem `xml:"criteriaItems"`
	Operation              TextLiteral                `xml:"operation"`
	PrimaryTableColumn     TextLiteral                `xml:"primaryTableColumn"`
	RelatedTable           TextLiteral                `xml:"relatedTable"`
	RelatedTableJoinColumn TextLiteral                `xml:"relatedTableJoinColumn"`
}

// ReportBlockInfo identifies a block of a joined report, or, at the top
// level of a joined report, the report itself
type ReportBlockInfo struct {
	AggregateReferences []struct {
		Aggregate TextLiteral `xml:"aggregate"`
	} `xml:"aggregateReferences"`
	BlockId   NillableText `xml:"blockId"`
	JoinTable TextLiteral  `xml:"joinTable"`
}

// ReportBlock represents one block of a joined report
type ReportBlock struct {
	BlockInfo       ReportBlockInfo        `xml:"blockInfo"`
	Columns         []ReportColumn         `xml:"columns"`
	Filter          *ReportFilter          `xml:"filter"`
	Format          TextLiteral            `xml:"format"`
	Name            TextLiteral            `xml:"name"`
	Params          []ReportParam          `xml:"params"`
	ReportType      TextLiteral            `xml:"reportType"`
	Scope           *TextLiteral           `xml:"scope"`
	ShowGrandTotal  *TextLiteral           `xml:"showGrandTotal"`
	ShowSubTotals   *TextLiteral           `xml:"showSubTotals"`
	SortColumn      *TextLiteral           `xml:"sortColumn"`
	SortOrder       *TextLiteral           `xml:"sortOrder"`
	TimeFrameFilter *ReportTimeFrameFilter `xml:"timeFrameFilter"`
}
