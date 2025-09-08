package report

import (
	. "github.com/ForceCLI/force-md/general"
)

// ReportColumn represents a column in a report
type ReportColumn struct {
	AggregateTypes *TextLiteral `xml:"aggregateTypes"`
	Field          TextLiteral  `xml:"field"`
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
	Column         TextLiteral `xml:"column"`
	ColumnToColumn TextLiteral `xml:"columnToColumn"`
	IsUnlocked     TextLiteral `xml:"isUnlocked"`
	Operator       TextLiteral `xml:"operator"`
	Value          TextLiteral `xml:"value"`
}

// ReportFilter represents a filter in a report
type ReportFilter struct {
	BooleanFilter *TextLiteral               `xml:"booleanFilter"`
	CriteriaItems []ReportFilterCriteriaItem `xml:"criteriaItems"`
}

// ReportBucketValue represents a bucket value
type ReportBucketValue struct {
	SourceValues []struct {
		SourceValue TextLiteral `xml:"sourceValue"`
	} `xml:"sourceValues"`
	Value TextLiteral `xml:"value"`
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
