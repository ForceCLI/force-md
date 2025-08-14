package objects

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"

	"github.com/ForceCLI/force-md/metadata/objects/businessprocess"
	"github.com/ForceCLI/force-md/metadata/objects/compactlayout"
	"github.com/ForceCLI/force-md/metadata/objects/field"
	"github.com/ForceCLI/force-md/metadata/objects/fieldset"
	"github.com/ForceCLI/force-md/metadata/objects/index"
	"github.com/ForceCLI/force-md/metadata/objects/listview"
	"github.com/ForceCLI/force-md/metadata/objects/recordtype"
	"github.com/ForceCLI/force-md/metadata/objects/sharingreason"
	"github.com/ForceCLI/force-md/metadata/objects/validationrule"
	"github.com/ForceCLI/force-md/metadata/objects/weblink"
)

const NAME = "CustomObject"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

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
	metadata.MetadataInfo
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
	ValidationRules validationrule.ValidationRuleList `xml:"validationRules"`
	Visibility      *struct {
		Text string `xml:",chardata"`
	} `xml:"visibility"`
	WebLinks []weblink.WebLink `xml:"webLinks"`
}

func (c *CustomObject) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomObject) Type() metadata.MetadataType {
	return NAME
}

func (c *CustomObject) Files(format metadata.Format) (map[string][]byte, error) {
	if format == metadata.SourceFormat {
		// For source format, decompose the object into separate files
		return c.decompose()
	}

	// For metadata format, use default behavior (single file)
	files := make(map[string][]byte)
	content, err := internal.Marshal(c)
	if err != nil {
		return nil, err
	}

	objectName := c.MetadataInfo.Name()
	fileName := "objects/" + string(objectName) + ".object"
	files[fileName] = content

	return files, nil
}

func (c *CustomObject) decompose() (map[string][]byte, error) {
	files := make(map[string][]byte)
	objectName := string(c.MetadataInfo.Name())
	baseDir := "objects/" + objectName

	// Clone the object and remove child components that will be written as separate files
	minimalObj := *c
	minimalObj.XMLName = xml.Name{Local: "CustomObject"}
	minimalObj.Xmlns = "http://soap.sforce.com/2006/04/metadata"

	// Clear out the child components that will be written separately
	minimalObj.Fields = nil
	minimalObj.RecordTypes = nil
	minimalObj.ValidationRules = nil
	minimalObj.Indexes = nil
	minimalObj.FieldSets = nil
	minimalObj.WebLinks = nil
	minimalObj.CompactLayouts = nil
	minimalObj.SharingReasons = nil
	minimalObj.BusinessProcesses = nil
	minimalObj.ListViews = nil

	// Marshal the minimal object
	content, err := internal.Marshal(&minimalObj)
	if err != nil {
		return nil, err
	}
	files[baseDir+"/"+objectName+".object-meta.xml"] = content

	// Write fields as separate files
	for _, f := range c.Fields {
		fieldMeta := &field.CustomField{
			MetadataInfo: metadata.MetadataInfo{},
			XMLName:      xml.Name{Local: "CustomField"},
			Xmlns:        "http://soap.sforce.com/2006/04/metadata",
			Field:        f,
		}
		fieldContent, err := internal.Marshal(fieldMeta)
		if err != nil {
			return nil, err
		}
		if f.FullName == "" {
			continue
		}
		files[baseDir+"/fields/"+f.FullName+".field-meta.xml"] = fieldContent
	}

	// Write record types as separate files
	for _, rt := range c.RecordTypes {
		rtMeta := &recordtype.RecordTypeMetadata{
			MetadataInfo: metadata.MetadataInfo{},
			XMLName:      xml.Name{Local: "RecordType"},
			Xmlns:        "http://soap.sforce.com/2006/04/metadata",
			RecordType:   rt,
		}
		rtContent, err := internal.Marshal(rtMeta)
		if err != nil {
			return nil, err
		}
		files[baseDir+"/recordTypes/"+rt.FullName+".recordType-meta.xml"] = rtContent
	}

	// Write validation rules as separate files
	for _, vr := range c.ValidationRules {
		vrMeta := &validationrule.ValidationRule{
			MetadataInfo: metadata.MetadataInfo{},
			XMLName:      xml.Name{Local: "ValidationRule"},
			Xmlns:        "http://soap.sforce.com/2006/04/metadata",
			Rule:         vr,
		}
		vrContent, err := internal.Marshal(vrMeta)
		if err != nil {
			return nil, err
		}
		files[baseDir+"/validationRules/"+vr.FullName+".validationRule-meta.xml"] = vrContent
	}

	// Write indexes as separate files
	for _, idx := range c.Indexes {
		idxMeta := &index.Index{
			MetadataInfo:   metadata.MetadataInfo{},
			XMLName:        xml.Name{Local: "Index"},
			Xmlns:          "http://soap.sforce.com/2006/04/metadata",
			BigObjectIndex: idx,
		}
		idxContent, err := internal.Marshal(idxMeta)
		if err != nil {
			return nil, err
		}
		files[baseDir+"/indexes/"+idx.Label+".index-meta.xml"] = idxContent
	}

	// Write field sets as separate files
	for _, fs := range c.FieldSets {
		fsMeta := &fieldset.FieldSetMetadata{
			MetadataInfo: metadata.MetadataInfo{},
			XMLName:      xml.Name{Local: "FieldSet"},
			Xmlns:        "http://soap.sforce.com/2006/04/metadata",
			FieldSet:     fs,
		}
		fsContent, err := internal.Marshal(fsMeta)
		if err != nil {
			return nil, err
		}
		files[baseDir+"/fieldSets/"+fs.FullName+".fieldSet-meta.xml"] = fsContent
	}

	// Write web links as separate files
	for _, wl := range c.WebLinks {
		wlMeta := &weblink.WebLinkMetadata{
			MetadataInfo: metadata.MetadataInfo{},
			XMLName:      xml.Name{Local: "WebLink"},
			Xmlns:        "http://soap.sforce.com/2006/04/metadata",
			WebLink:      wl,
		}
		wlContent, err := internal.Marshal(wlMeta)
		if err != nil {
			return nil, err
		}
		files[baseDir+"/webLinks/"+wl.FullName+".webLink-meta.xml"] = wlContent
	}

	// Write compact layouts as separate files
	for _, cl := range c.CompactLayouts {
		clMeta := &compactlayout.CompactLayoutMetadata{
			MetadataInfo:  metadata.MetadataInfo{},
			XMLName:       xml.Name{Local: "CompactLayout"},
			Xmlns:         "http://soap.sforce.com/2006/04/metadata",
			CompactLayout: cl,
		}
		clContent, err := internal.Marshal(clMeta)
		if err != nil {
			return nil, err
		}
		files[baseDir+"/compactLayouts/"+cl.FullName.Text+".compactLayout-meta.xml"] = clContent
	}

	// Write sharing reasons as separate files
	for _, sr := range c.SharingReasons {
		srMeta := &sharingreason.SharingReasonMetadata{
			MetadataInfo:  metadata.MetadataInfo{},
			XMLName:       xml.Name{Local: "SharingReason"},
			Xmlns:         "http://soap.sforce.com/2006/04/metadata",
			SharingReason: sr,
		}
		srContent, err := internal.Marshal(srMeta)
		if err != nil {
			return nil, err
		}
		files[baseDir+"/sharingReasons/"+sr.FullName+".sharingReason-meta.xml"] = srContent
	}

	// Write business processes as separate files
	for _, bp := range c.BusinessProcesses {
		bpMeta := &businessprocess.BusinessProcessMetadata{
			MetadataInfo:    metadata.MetadataInfo{},
			XMLName:         xml.Name{Local: "BusinessProcess"},
			Xmlns:           "http://soap.sforce.com/2006/04/metadata",
			BusinessProcess: bp,
		}
		bpContent, err := internal.Marshal(bpMeta)
		if err != nil {
			return nil, err
		}
		files[baseDir+"/businessProcesses/"+bp.FullName.Text+".businessProcess-meta.xml"] = bpContent
	}

	// Write list views as separate files
	for _, lv := range c.ListViews {
		lvMeta := &listview.ListViewMetadata{
			MetadataInfo: metadata.MetadataInfo{},
			XMLName:      xml.Name{Local: "ListView"},
			Xmlns:        "http://soap.sforce.com/2006/04/metadata",
			ListView:     lv,
		}
		lvContent, err := internal.Marshal(lvMeta)
		if err != nil {
			return nil, err
		}
		files[baseDir+"/listViews/"+lv.FullName.Text+".listView-meta.xml"] = lvContent
	}

	return files, nil
}

func Open(path string) (*CustomObject, error) {
	p := &CustomObject{}
	return p, metadata.ParseMetadataXml(p, path)
}
