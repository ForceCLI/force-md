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

type Field struct {
	FullName          string       `xml:"fullName"`
	BusinessStatus    *TextLiteral `xml:"businessStatus"`
	BusinessOwnerUser *TextLiteral `xml:"businessOwnerUser"`
	CaseSensitive     *struct {
		Text string `xml:",chardata"`
	} `xml:"caseSensitive"`
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
	EncryptionScheme   *TextLiteral `xml:"encryptionScheme"`
	ExternalId         *BooleanText `xml:"externalId"`
	FieldManageability *struct {
		Text string `xml:",chardata"`
	} `xml:"fieldManageability"`
	Formula              *TextLiteral `xml:"formula"`
	FormulaTreatBlanksAs *struct {
		Text string `xml:",chardata"`
	} `xml:"formulaTreatBlanksAs"`
	InlineHelpText      *TextLiteral `xml:"inlineHelpText"`
	IsFilteringDisabled *BooleanText `xml:"isFilteringDisabled"`
	IsNameField         *BooleanText `xml:"isNameField"`
	IsSortingDisabled   *BooleanText `xml:"isSortingDisabled"`
	Label               *TextLiteral `xml:"label"`
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
	Precision              *IntegerText `xml:"precision"`
	Length                 *IntegerText `xml:"length"`
	MaskChar               *TextLiteral `xml:"maskChar"`
	MaskType               *TextLiteral `xml:"maskType"`
	ReferenceTo            *TextLiteral `xml:"referenceTo"`
	RelationshipLabel      *TextLiteral `xml:"relationshipLabel"`
	RelationshipName       *TextLiteral `xml:"relationshipName"`
	RestrictedAdminField   *TextLiteral `xml:"restrictedAdminField"`
	Required               *BooleanText `xml:"required"`
	Scale                  *IntegerText `xml:"scale"`
	SecurityClassification *TextLiteral `xml:"securityClassification"`
	TrackFeedHistory       *struct {
		Text string `xml:",chardata"`
	} `xml:"trackFeedHistory"`
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
	RelationshipOrder *struct {
		Text string `xml:",chardata"`
	} `xml:"relationshipOrder"`
	ReparentableMasterDetail *struct {
		Text string `xml:",chardata"`
	} `xml:"reparentableMasterDetail"`
	MetadataRelationshipControllingField *struct {
		Text string `xml:",chardata"`
	} `xml:"metadataRelationshipControllingField"`
	TrackHistory            *BooleanText `xml:"trackHistory"`
	TrackTrending           *BooleanText `xml:"trackTrending"`
	Type                    *TextLiteral `xml:"type"`
	Unique                  *BooleanText `xml:"unique"`
	WriteRequiresMasterRead *struct {
		Text string `xml:",chardata"`
	} `xml:"writeRequiresMasterRead"`
	ValueSet *struct {
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
			Value []struct {
				FullName string `xml:"fullName"`
				Default  struct {
					Text string `xml:",chardata"`
				} `xml:"default"`
				IsActive *BooleanText `xml:"isActive"`
				Label    struct {
					Text string `xml:",innerxml"`
				} `xml:"label"`
				Color *struct {
					Text string `xml:",chardata"`
				} `xml:"color"`
			} `xml:"value"`
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
