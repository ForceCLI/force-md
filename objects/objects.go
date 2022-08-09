package objects

import (
	"encoding/xml"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
)

type FieldList []Field

type CustomField struct {
	XMLName xml.Name `xml:"CustomField"`
	Xmlns   string   `xml:"xmlns,attr"`
	Field
}

type Field struct {
	FullName       string       `xml:"fullName"`
	BusinessStatus *TextLiteral `xml:"businessStatus"`
	CaseSensitive  *struct {
		Text string `xml:",chardata"`
	} `xml:"caseSensitive"`
	DefaultValue     *TextLiteral `xml:"defaultValue"`
	DeleteConstraint *struct {
		Text string `xml:",chardata"`
	} `xml:"deleteConstraint"`
	Description   *TextLiteral `xml:"description"`
	DisplayFormat *struct {
		Text string `xml:",chardata"`
	} `xml:"displayFormat"`
	DisplayLocationInDecimal *struct {
		Text string `xml:",chardata"`
	} `xml:"displayLocationInDecimal"`
	ExternalId *BooleanText `xml:"externalId"`
	Formula    *struct {
		Text string `xml:",innerxml"`
	} `xml:"formula"`
	FormulaTreatBlanksAs *struct {
		Text string `xml:",chardata"`
	} `xml:"formulaTreatBlanksAs"`
	InlineHelpText *TextLiteral `xml:"inlineHelpText"`
	Label          *TextLiteral `xml:"label"`
	LookupFilter   *struct {
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
	ReferenceTo            *TextLiteral `xml:"referenceTo"`
	RelationshipLabel      *TextLiteral `xml:"relationshipLabel"`
	RelationshipName       *TextLiteral `xml:"relationshipName"`
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
	FieldManageability *struct {
		Text string `xml:",chardata"`
	} `xml:"fieldManageability"`
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
				FullName struct {
					Text string `xml:",innerxml"`
				} `xml:"fullName"`
				Default struct {
					Text string `xml:",chardata"`
				} `xml:"default"`
				IsActive *struct {
					Text string `xml:",chardata"`
				} `xml:"isActive"`
				Label struct {
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

type FieldSet struct {
	FullName        string `xml:"fullName"`
	AvailableFields []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		IsFieldManaged struct {
			Text string `xml:",chardata"`
		} `xml:"isFieldManaged"`
		IsRequired struct {
			Text string `xml:",chardata"`
		} `xml:"isRequired"`
	} `xml:"availableFields"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	DisplayedFields []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		IsFieldManaged struct {
			Text string `xml:",chardata"`
		} `xml:"isFieldManaged"`
		IsRequired struct {
			Text string `xml:",chardata"`
		} `xml:"isRequired"`
	} `xml:"displayedFields"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
}

type RecordType struct {
	FullName string `xml:"fullName"`
	Active   struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
	CompactLayoutAssignment *struct {
		Text string `xml:",chardata"`
	} `xml:"compactLayoutAssignment"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	PicklistValues []struct {
		Picklist string `xml:"picklist"`
		Values   []struct {
			FullName string      `xml:"fullName"`
			Default  BooleanText `xml:"default"`
		} `xml:"values"`
	} `xml:"picklistValues"`
}
type CustomObject struct {
	XMLName         xml.Name `xml:"CustomObject"`
	Xmlns           string   `xml:"xmlns,attr"`
	ActionOverrides []struct {
		ActionName struct {
			Text string `xml:",chardata"`
		} `xml:"actionName"`
		Comment *struct {
			Text string `xml:",chardata"`
		} `xml:"comment"`
		Content *struct {
			Text string `xml:",chardata"`
		} `xml:"content"`
		FormFactor *struct {
			Text string `xml:",chardata"`
		} `xml:"formFactor"`
		SkipRecordTypeSelect *struct {
			Text string `xml:",chardata"`
		} `xml:"skipRecordTypeSelect"`
		Type struct {
			Text string `xml:",chardata"`
		} `xml:"type"`
	} `xml:"actionOverrides"`
	AllowInChatterGroups *struct {
		Text string `xml:",chardata"`
	} `xml:"allowInChatterGroups"`
	CompactLayoutAssignment struct {
		Text string `xml:",chardata"`
	} `xml:"compactLayoutAssignment"`
	CompactLayouts []struct {
		FullName struct {
			Text string `xml:",chardata"`
		} `xml:"fullName"`
		Fields []struct {
			Text string `xml:",chardata"`
		} `xml:"fields"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
	} `xml:"compactLayouts"`
	DeploymentStatus *struct {
		Text string `xml:",chardata"`
	} `xml:"deploymentStatus"`
	Description *struct {
		Text string `xml:",innerxml"`
	} `xml:"description"`
	EnableActivities *struct {
		Text string `xml:",chardata"`
	} `xml:"enableActivities"`
	EnableBulkApi *struct {
		Text string `xml:",chardata"`
	} `xml:"enableBulkApi"`
	EnableEnhancedLookup *struct {
		Text string `xml:",chardata"`
	} `xml:"enableEnhancedLookup"`
	EnableFeeds struct {
		Text string `xml:",chardata"`
	} `xml:"enableFeeds"`
	EnableHistory struct {
		Text string `xml:",chardata"`
	} `xml:"enableHistory"`
	EnableLicensing *struct {
		Text string `xml:",chardata"`
	} `xml:"enableLicensing"`
	EnableReports *struct {
		Text string `xml:",chardata"`
	} `xml:"enableReports"`
	EnableSearch *struct {
		Text string `xml:",chardata"`
	} `xml:"enableSearch"`
	EnableSharing *struct {
		Text string `xml:",chardata"`
	} `xml:"enableSharing"`
	EnableStreamingApi *struct {
		Text string `xml:",chardata"`
	} `xml:"enableStreamingApi"`
	ExternalSharingModel *struct {
		Text string `xml:",chardata"`
	} `xml:"externalSharingModel"`
	FieldSets []FieldSet `xml:"fieldSets"`
	Fields    FieldList  `xml:"fields"`
	Label     *struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	ListViews []struct {
		FullName struct {
			Text string `xml:",chardata"`
		} `xml:"fullName"`
		BooleanFilter *struct {
			Text string `xml:",chardata"`
		} `xml:"booleanFilter"`
		Columns []struct {
			Text string `xml:",chardata"`
		} `xml:"columns"`
		FilterScope struct {
			Text string `xml:",chardata"`
		} `xml:"filterScope"`
		Filters []struct {
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			Operation struct {
				Text string `xml:",chardata"`
			} `xml:"operation"`
			Value *struct {
				Text string `xml:",chardata"`
			} `xml:"value"`
		} `xml:"filters"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		Queue *struct {
			Text string `xml:",chardata"`
		} `xml:"queue"`
		SharedTo struct {
			Group []struct {
				Text string `xml:",chardata"`
			} `xml:"group"`
			Rol []struct {
				Text string `xml:",chardata"`
			} `xml:"role"`
			RoleAndSubordinates []struct {
				Text string `xml:",chardata"`
			} `xml:"roleAndSubordinates"`
			AllInternalUsers *struct {
				Text string `xml:",chardata"`
			} `xml:"allInternalUsers"`
		} `xml:"sharedTo"`
		Language *struct {
			Text string `xml:",chardata"`
		} `xml:"language"`
	} `xml:"listViews"`
	NameField *struct {
		DisplayFormat *struct {
			Text string `xml:",chardata"`
		} `xml:"displayFormat"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		TrackHistory *struct {
			Text string `xml:",chardata"`
		} `xml:"trackHistory"`
		Type struct {
			Text string `xml:",chardata"`
		} `xml:"type"`
	} `xml:"nameField"`
	PluralLabel *struct {
		Text string `xml:",chardata"`
	} `xml:"pluralLabel"`
	RecordTypeTrackFeedHistory *struct {
		Text string `xml:",chardata"`
	} `xml:"recordTypeTrackFeedHistory"`
	RecordTypeTrackHistory *struct {
		Text string `xml:",chardata"`
	} `xml:"recordTypeTrackHistory"`
	RecordTypes   []RecordType `xml:"recordTypes"`
	SearchLayouts struct {
		CustomTabListAdditionalFields []struct {
			Text string `xml:",chardata"`
		} `xml:"customTabListAdditionalFields"`
		ListViewButtons []struct {
			Text string `xml:",chardata"`
		} `xml:"listViewButtons"`
		LookupDialogsAdditionalFields []struct {
			Text string `xml:",chardata"`
		} `xml:"lookupDialogsAdditionalFields"`
		LookupPhoneDialogsAdditionalFields []struct {
			Text string `xml:",chardata"`
		} `xml:"lookupPhoneDialogsAdditionalFields"`
		SearchFilterFields []struct {
			Text string `xml:",chardata"`
		} `xml:"searchFilterFields"`
		SearchResultsAdditionalFields []struct {
			Text string `xml:",chardata"`
		} `xml:"searchResultsAdditionalFields"`
		ExcludedStandardButtons []struct {
			Text string `xml:",chardata"`
		} `xml:"excludedStandardButtons"`
	} `xml:"searchLayouts"`
	SharingModel struct {
		Text string `xml:",chardata"`
	} `xml:"sharingModel"`
	ValidationRules []struct {
		FullName struct {
			Text string `xml:",chardata"`
		} `xml:"fullName"`
		Active struct {
			Text string `xml:",chardata"`
		} `xml:"active"`
		Description *struct {
			Text string `xml:",innerxml"`
		} `xml:"description"`
		ErrorConditionFormula struct {
			Text string `xml:",innerxml"`
		} `xml:"errorConditionFormula"`
		ErrorDisplayField *struct {
			Text string `xml:",chardata"`
		} `xml:"errorDisplayField"`
		ErrorMessage struct {
			Text string `xml:",innerxml"`
		} `xml:"errorMessage"`
	} `xml:"validationRules"`
	Visibility *struct {
		Text string `xml:",chardata"`
	} `xml:"visibility"`
	WebLinks []struct {
		FullName struct {
			Text string `xml:",chardata"`
		} `xml:"fullName"`
		Availability struct {
			Text string `xml:",chardata"`
		} `xml:"availability"`
		Description *struct {
			Text string `xml:",chardata"`
		} `xml:"description"`
		DisplayType struct {
			Text string `xml:",chardata"`
		} `xml:"displayType"`
		EncodingKey *struct {
			Text string `xml:",chardata"`
		} `xml:"encodingKey"`
		HasMenubar *struct {
			Text string `xml:",chardata"`
		} `xml:"hasMenubar"`
		HasScrollbars *struct {
			Text string `xml:",chardata"`
		} `xml:"hasScrollbars"`
		HasToolbar *struct {
			Text string `xml:",chardata"`
		} `xml:"hasToolbar"`
		Height *struct {
			Text string `xml:",chardata"`
		} `xml:"height"`
		IsResizable *struct {
			Text string `xml:",chardata"`
		} `xml:"isResizable"`
		LinkType struct {
			Text string `xml:",chardata"`
		} `xml:"linkType"`
		MasterLabel struct {
			Text string `xml:",chardata"`
		} `xml:"masterLabel"`
		OpenType struct {
			Text string `xml:",chardata"`
		} `xml:"openType"`
		Page *struct {
			Text string `xml:",chardata"`
		} `xml:"page"`
		Position *struct {
			Text string `xml:",chardata"`
		} `xml:"position"`
		Protected struct {
			Text string `xml:",chardata"`
		} `xml:"protected"`
		ShowsLocation *struct {
			Text string `xml:",chardata"`
		} `xml:"showsLocation"`
		ShowsStatus *struct {
			Text string `xml:",chardata"`
		} `xml:"showsStatus"`
		URL *struct {
			Text string `xml:",innerxml"`
		} `xml:"url"`
		Width *struct {
			Text string `xml:",chardata"`
		} `xml:"width"`
		RequireRowSelection *struct {
			Text string `xml:",chardata"`
		} `xml:"requireRowSelection"`
	} `xml:"webLinks"`
	CustomSettingsType *struct {
		Text string `xml:",chardata"`
	} `xml:"customSettingsType"`
	StartsWith *struct {
		Text string `xml:",chardata"`
	} `xml:"startsWith"`
	Deprecated *struct {
		Text string `xml:",chardata"`
	} `xml:"deprecated"`
}

func (p *CustomObject) MetaCheck() {}

func Open(path string) (*CustomObject, error) {
	p := &CustomObject{}
	return p, internal.ParseMetadataXml(p, path)
}
