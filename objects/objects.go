package objects

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"

	"github.com/ForceCLI/force-md/objects/businessprocess"
	"github.com/ForceCLI/force-md/objects/compactlayout"
	"github.com/ForceCLI/force-md/objects/field"
	"github.com/ForceCLI/force-md/objects/fieldset"
	"github.com/ForceCLI/force-md/objects/index"
	"github.com/ForceCLI/force-md/objects/listview"
	"github.com/ForceCLI/force-md/objects/recordtype"
	"github.com/ForceCLI/force-md/objects/sharingreason"
	"github.com/ForceCLI/force-md/objects/validationrule"
	"github.com/ForceCLI/force-md/objects/weblink"
)

type FieldList []field.Field
type IndexList []index.BigObjectIndex
type ListViewList []listview.ListView

type ActionOverride struct {
	ActionName           string       `xml:"actionName"`
	Comment              *string      `xml:"comment"`
	Content              *string      `xml:"content"`
	FormFactor           *string      `xml:"formFactor"`
	SkipRecordTypeSelect *BooleanText `xml:"skipRecordTypeSelect"`
	Type                 string       `xml:"type"`
}

type CustomObject struct {
	Metadata
	XMLName              xml.Name         `xml:"CustomObject"`
	Xmlns                string           `xml:"xmlns,attr"`
	ActionOverrides      []ActionOverride `xml:"actionOverrides"`
	AllowInChatterGroups *struct {
		Text string `xml:",chardata"`
	} `xml:"allowInChatterGroups"`
	BusinessProcesses       []businessprocess.BusinessProcess `xml:"businessProcesses"`
	CompactLayoutAssignment *struct {
		Text string `xml:",chardata"`
	} `xml:"compactLayoutAssignment"`
	CompactLayouts []compactlayout.CompactLayout `xml:"compactLayouts"`
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
	FieldSets []fieldset.FieldSet `xml:"fieldSets"`
	Fields    FieldList           `xml:"fields"`
	Indexes   IndexList           `xml:"indexes"`
	Label     *struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	ListViews ListViewList `xml:"listViews"`
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
	ProfileSearchLayouts []struct {
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
	RecordTypes   []recordtype.RecordType `xml:"recordTypes"`
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
	SharingReasons []sharingreason.SharingReason `xml:"sharingReasons"`
	StartsWith     *struct {
		Text string `xml:",chardata"`
	} `xml:"startsWith"`
	ValidationRules []validationrule.Rule `xml:"validationRules"`
	Visibility      *struct {
		Text string `xml:",chardata"`
	} `xml:"visibility"`
	WebLinks []weblink.WebLink `xml:"webLinks"`
}

func (c *CustomObject) SetMetadata(m Metadata) {
	c.Metadata = m
}

func Open(path string) (*CustomObject, error) {
	p := &CustomObject{}
	return p, internal.ParseMetadataXml(p, path)
}
