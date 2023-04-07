package objects

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type FieldList []Field

type CustomField struct {
	XMLName xml.Name `xml:"CustomField"`
	Xmlns   string   `xml:"xmlns,attr"`
	Field
}

type ValidationRule struct {
	XMLName xml.Name `xml:"CustomField"`
	Xmlns   string   `xml:"xmlns,attr"`
	Rule
}

type Rule struct {
	FullName string `xml:"fullName"`
	Active   struct {
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
	ExternalId         *BooleanText `xml:"externalId"`
	FieldManageability *struct {
		Text string `xml:",chardata"`
	} `xml:"fieldManageability"`
	Formula *struct {
		Text string `xml:",innerxml"`
	} `xml:"formula"`
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

type WebLink struct {
	FullName     string `xml:"fullName"`
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
	RequireRowSelection *struct {
		Text string `xml:",chardata"`
	} `xml:"requireRowSelection"`
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

type ValueSetOption struct {
	FullName string      `xml:"fullName"`
	Default  BooleanText `xml:"default"`
}

type ValueSetOptionList []ValueSetOption

type Picklist struct {
	Picklist string             `xml:"picklist"`
	Values   ValueSetOptionList `xml:"values"`
}

type PicklistList []Picklist

type RecordType struct {
	FullName string `xml:"fullName"`
	Active   struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
	BusinessProcess *struct {
		Text string `xml:",chardata"`
	} `xml:"businessProcess"`
	CompactLayoutAssignment *struct {
		Text string `xml:",chardata"`
	} `xml:"compactLayoutAssignment"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	PicklistValues PicklistList `xml:"picklistValues"`
}

type ActionOverride struct {
	ActionName           string       `xml:"actionName"`
	Comment              *string      `xml:"comment"`
	Content              *string      `xml:"content"`
	FormFactor           *string      `xml:"formFactor"`
	SkipRecordTypeSelect *BooleanText `xml:"skipRecordTypeSelect"`
	Type                 string       `xml:"type"`
}

type CustomObject struct {
	XMLName              xml.Name         `xml:"CustomObject"`
	Xmlns                string           `xml:"xmlns,attr"`
	ActionOverrides      []ActionOverride `xml:"actionOverrides"`
	AllowInChatterGroups *struct {
		Text string `xml:",chardata"`
	} `xml:"allowInChatterGroups"`
	BusinessProcesses []struct {
		FullName struct {
			Text string `xml:",chardata"`
		} `xml:"fullName"`
		Description *struct {
			Text string `xml:",chardata"`
		} `xml:"description"`
		IsActive struct {
			Text string `xml:",chardata"`
		} `xml:"isActive"`
		Values []struct {
			FullName struct {
				Text string `xml:",chardata"`
			} `xml:"fullName"`
			Default struct {
				Text string `xml:",chardata"`
			} `xml:"default"`
		} `xml:"values"`
	} `xml:"businessProcesses"`
	CompactLayoutAssignment *struct {
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
	CustomHelpPage *struct {
		Text string `xml:",chardata"`
	} `xml:"customHelpPage"`
	CustomSettingsType *struct {
		Text string `xml:",chardata"`
	} `xml:"customSettingsType"`
	DeploymentStatus *struct {
		Text string `xml:",chardata"`
	} `xml:"deploymentStatus"`
	Deprecated *struct {
		Text string `xml:",chardata"`
	} `xml:"deprecated"`
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
	EnableFeeds *struct {
		Text string `xml:",chardata"`
	} `xml:"enableFeeds"`
	EnableHistory *struct {
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
	EventType *struct {
		Text string `xml:",chardata"`
	} `xml:"eventType"`
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
		SharedTo *struct {
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
		TrackFeedHistory *struct {
			Text string `xml:",chardata"`
		} `xml:"trackFeedHistory"`
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
	ProfileSearchLayouts *struct {
		Fields []struct {
			Text string `xml:",chardata"`
		} `xml:"fields"`
		ProfileName struct {
			Text string `xml:",chardata"`
		} `xml:"profileName"`
	} `xml:"profileSearchLayouts"`
	PublishBehavior *struct {
		Text string `xml:",chardata"`
	} `xml:"publishBehavior"`
	RecordTypeTrackFeedHistory *struct {
		Text string `xml:",chardata"`
	} `xml:"recordTypeTrackFeedHistory"`
	RecordTypeTrackHistory *struct {
		Text string `xml:",chardata"`
	} `xml:"recordTypeTrackHistory"`
	RecordTypes   []RecordType `xml:"recordTypes"`
	SearchLayouts *struct {
		CustomTabListAdditionalFields []struct {
			Text string `xml:",chardata"`
		} `xml:"customTabListAdditionalFields"`
		ExcludedStandardButtons []struct {
			Text string `xml:",chardata"`
		} `xml:"excludedStandardButtons"`
		ListViewButtons []struct {
			Text string `xml:",chardata"`
		} `xml:"listViewButtons"`
		LookupDialogsAdditionalFields []struct {
			Text string `xml:",chardata"`
		} `xml:"lookupDialogsAdditionalFields"`
		LookupFilterFields []struct {
			Text string `xml:",chardata"`
		} `xml:"lookupFilterFields"`
		LookupPhoneDialogsAdditionalFields []struct {
			Text string `xml:",chardata"`
		} `xml:"lookupPhoneDialogsAdditionalFields"`
		SearchFilterFields []struct {
			Text string `xml:",chardata"`
		} `xml:"searchFilterFields"`
		SearchResultsAdditionalFields []struct {
			Text string `xml:",chardata"`
		} `xml:"searchResultsAdditionalFields"`
		SearchResultsCustomButtons []struct {
			Text string `xml:",chardata"`
		} `xml:"searchResultsCustomButtons"`
	} `xml:"searchLayouts"`
	SharingModel *struct {
		Text string `xml:",chardata"`
	} `xml:"sharingModel"`
	SharingReasons []struct {
		FullName string `xml:"fullName"`
		Label    string `xml:"label"`
	} `xml:"sharingReasons"`
	StartsWith *struct {
		Text string `xml:",chardata"`
	} `xml:"startsWith"`
	ValidationRules []Rule `xml:"validationRules"`
	Visibility      *struct {
		Text string `xml:",chardata"`
	} `xml:"visibility"`
	WebLinks []WebLink `xml:"webLinks"`
}

func (p *CustomObject) MetaCheck() {}

func Open(path string) (*CustomObject, error) {
	p := &CustomObject{}
	return p, internal.ParseMetadataXml(p, path)
}
