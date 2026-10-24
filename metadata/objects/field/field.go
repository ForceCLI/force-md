package field

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objects/split"
)

const NAME = "CustomField"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type FieldFilter func(Field) bool

type CustomField struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"CustomField"`
	Xmlns   string   `xml:"xmlns,attr"`
	Field
}

// ValueSetValue is one value of a picklist field's value set.
type ValueSetValue struct {
	FullName string `xml:"fullName"`
	Color    *struct {
		Text string `xml:",chardata"`
	} `xml:"color"`
	Default struct {
		Text string `xml:",chardata"`
	} `xml:"default"`
	Description *TextLiteral `xml:"description"`
	IsActive    *BooleanText `xml:"isActive"`
	Label       struct {
		Text string `xml:",innerxml"`
	} `xml:"label"`
}

type Field struct {
	FullName          string       `xml:"fullName"`
	BusinessOwnerUser *TextLiteral `xml:"businessOwnerUser"`
	BusinessStatus    *TextLiteral `xml:"businessStatus"`
	CaseSensitive     *struct {
		Text string `xml:",chardata"`
	} `xml:"caseSensitive"`
	ComplianceGroup  *TextLiteral `xml:"complianceGroup"`
	DefaultValue     *TextLiteral `xml:"defaultValue"`
	DeleteConstraint *struct {
		Text string `xml:",chardata"`
	} `xml:"deleteConstraint"`
	Deprecated    *TextLiteral `xml:"deprecated"`
	Description   *TextLiteral `xml:"description"`
	DisplayFormat *struct {
		Text string `xml:",chardata"`
	} `xml:"displayFormat"`
	DisplayLocationInDecimal *struct {
		Text string `xml:",chardata"`
	} `xml:"displayLocationInDecimal"`
	EncryptionScheme *TextLiteral `xml:"encryptionScheme"`
	// ExternalDeveloperName is an external object field's name in the
	// external system.
	ExternalDeveloperName *TextLiteral `xml:"externalDeveloperName"`
	ExternalId            *BooleanText `xml:"externalId"`
	FieldManageability    *struct {
		Text string `xml:",chardata"`
	} `xml:"fieldManageability"`
	Formula              *TextLiteral `xml:"formula"`
	FormulaTreatBlanksAs *struct {
		Text string `xml:",chardata"`
	} `xml:"formulaTreatBlanksAs"`
	// GlobalPicklist names the global value set of a picklist field in API
	// versions before 38.0, which replaced it with valueSet.
	GlobalPicklist      *TextLiteral `xml:"globalPicklist"`
	InlineHelpText      *TextLiteral `xml:"inlineHelpText"`
	IsAIPredictionField *BooleanText `xml:"isAIPredictionField"`
	IsFilteringDisabled *BooleanText `xml:"isFilteringDisabled"`
	IsNameField         *BooleanText `xml:"isNameField"`
	IsSortingDisabled   *BooleanText `xml:"isSortingDisabled"`
	Label               *TextLiteral `xml:"label"`
	Length              *IntegerText `xml:"length"`
	LookupFilter        *struct {
		Active struct {
			Text string `xml:",chardata"`
		} `xml:"active"`
		BooleanFilter *TextLiteral `xml:"booleanFilter"`
		ErrorMessage  *struct {
			Text string `xml:",innerxml"`
		} `xml:"errorMessage"`
		FilterItems []struct {
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			Operation struct {
				Text string `xml:",chardata"`
			} `xml:"operation"`
			Value *struct {
				Text string `xml:",chardata"`
			} `xml:"value"`
			ValueField *struct {
				Text string `xml:",chardata"`
			} `xml:"valueField"`
		} `xml:"filterItems"`
		InfoMessage *TextLiteral `xml:"infoMessage"`
		IsOptional  struct {
			Text string `xml:",chardata"`
		} `xml:"isOptional"`
	} `xml:"lookupFilter"`
	MaskChar                             *TextLiteral `xml:"maskChar"`
	MaskType                             *TextLiteral `xml:"maskType"`
	MetadataRelationshipControllingField *struct {
		Text string `xml:",chardata"`
	} `xml:"metadataRelationshipControllingField"`
	// MktDataLakeFieldAttributes is present on a Data Lake Object (__dll)
	// field; MktDataModelFieldAttributes on a Data Model Object (__dlm) field.
	MktDataLakeFieldAttributes  *MktDataLakeFieldAttributes  `xml:"mktDataLakeFieldAttributes"`
	MktDataModelFieldAttributes *MktDataModelFieldAttributes `xml:"mktDataModelFieldAttributes"`
	// Picklist holds a picklist field's values in API versions before 38.0,
	// which replaced it with valueSet.
	Picklist  *Picklist    `xml:"picklist"`
	Precision *IntegerText `xml:"precision"`
	// ReadOnlyProxy marks an external object field whose value the external
	// system supplies.
	ReadOnlyProxy     *BooleanText `xml:"readOnlyProxy"`
	ReferenceTo       *TextLiteral `xml:"referenceTo"`
	RelationshipLabel *TextLiteral `xml:"relationshipLabel"`
	RelationshipName  *TextLiteral `xml:"relationshipName"`
	RelationshipOrder *struct {
		Text string `xml:",chardata"`
	} `xml:"relationshipOrder"`
	ReparentableMasterDetail *struct {
		Text string `xml:",chardata"`
	} `xml:"reparentableMasterDetail"`
	Required               *BooleanText `xml:"required"`
	RestrictedAdminField   *TextLiteral `xml:"restrictedAdminField"`
	Scale                  *IntegerText `xml:"scale"`
	SecurityClassification *TextLiteral `xml:"securityClassification"`
	// StartingNumber is an auto-number field's first number.
	StartingNumber  *IntegerText `xml:"startingNumber"`
	SummarizedField *struct {
		Text string `xml:",chardata"`
	} `xml:"summarizedField"`
	SummaryFilterItems []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Operation struct {
			Text string `xml:",chardata"`
		} `xml:"operation"`
		Value struct {
			Text string `xml:",chardata"`
		} `xml:"value"`
	} `xml:"summaryFilterItems"`
	SummaryForeignKey *struct {
		Text string `xml:",chardata"`
	} `xml:"summaryForeignKey"`
	SummaryOperation *struct {
		Text string `xml:",chardata"`
	} `xml:"summaryOperation"`
	TrackFeedHistory *struct {
		Text string `xml:",chardata"`
	} `xml:"trackFeedHistory"`
	TrackHistory  *BooleanText `xml:"trackHistory"`
	TrackTrending *BooleanText `xml:"trackTrending"`
	Type          *TextLiteral `xml:"type"`
	Unique        *BooleanText `xml:"unique"`
	ValueSet      *struct {
		ControllingField *struct {
			Text string `xml:",chardata"`
		} `xml:"controllingField"`
		Restricted *struct {
			Text string `xml:",chardata"`
		} `xml:"restricted"`
		ValueSetDefinition *struct {
			Sorted struct {
				Text string `xml:",chardata"`
			} `xml:"sorted"`
			Value []ValueSetValue `xml:"value"`
		} `xml:"valueSetDefinition"`
		ValueSetName *struct {
			Text string `xml:",chardata"`
		} `xml:"valueSetName"`
		ValueSettings []struct {
			ControllingFieldValue []struct {
				Text string `xml:",innerxml"`
			} `xml:"controllingFieldValue"`
			ValueName struct {
				Text string `xml:",chardata"`
			} `xml:"valueName"`
		} `xml:"valueSettings"`
	} `xml:"valueSet"`
	VisibleLines *struct {
		Text string `xml:",chardata"`
	} `xml:"visibleLines"`
	WriteRequiresMasterRead *struct {
		Text string `xml:",chardata"`
	} `xml:"writeRequiresMasterRead"`
}

func (c *CustomField) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomField) Type() metadata.MetadataType {
	return NAME
}

func (c *CustomField) NameFromPath(path string) metadata.MetadataObjectName {
	return split.NameFromPath(path)
}

func Open(path string) (*CustomField, error) {
	p := &CustomField{}
	return p, metadata.ParseMetadataXml(p, path)
}

type MktDataLakeFieldAttributes struct {
	DefinitionCreationType *struct {
		Text string `xml:",chardata"`
	} `xml:"definitionCreationType"`
	ExternalName *struct {
		Text string `xml:",chardata"`
	} `xml:"externalName"`
	IsEventDate *struct {
		Text string `xml:",chardata"`
	} `xml:"isEventDate"`
	IsInternalOrganization *struct {
		Text string `xml:",chardata"`
	} `xml:"isInternalOrganization"`
	IsRecordModified *struct {
		Text string `xml:",chardata"`
	} `xml:"isRecordModified"`
	PrimaryIndexOrder *struct {
		Text string `xml:",chardata"`
	} `xml:"primaryIndexOrder"`
	UsageTag *struct {
		Text string `xml:",chardata"`
	} `xml:"usageTag"`
}

type MktDataModelFieldAttributes struct {
	DefinitionCreationType *struct {
		Text string `xml:",chardata"`
	} `xml:"definitionCreationType"`
	InvalidMergeActionType *struct {
		Text string `xml:",chardata"`
	} `xml:"invalidMergeActionType"`
	IsDynamicLookup *struct {
		Text string `xml:",chardata"`
	} `xml:"isDynamicLookup"`
	LabelOverride *struct {
		Text string `xml:",chardata"`
	} `xml:"labelOverride"`
	MasterLabel *struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	PrimaryIndexOrder *struct {
		Text string `xml:",chardata"`
	} `xml:"primaryIndexOrder"`
	RefAttrDeveloperName *struct {
		Text string `xml:",chardata"`
	} `xml:"refAttrDeveloperName"`
	UsageTag *struct {
		Text string `xml:",chardata"`
	} `xml:"usageTag"`
}
