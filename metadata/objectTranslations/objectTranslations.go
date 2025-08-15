package objectTranslations

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objectTranslations/field"
	"github.com/ForceCLI/force-md/metadata/objectTranslations/recordtype"
	"github.com/ForceCLI/force-md/metadata/objectTranslations/validationrule"
)

const OBJECT_TRANSLATIONS_NAME = "CustomObjectTranslation"

func init() {
	internal.TypeRegistry.Register(OBJECT_TRANSLATIONS_NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type FieldList []field.Field
type RecordTypeList []recordtype.RecordType
type ValidationRuleList []validationrule.ValidationRule

type CustomObjectTranslation struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"CustomObjectTranslation"`
	Xmlns      string   `xml:"xmlns,attr"`
	CaseValues []struct {
		Article *struct {
			Text string `xml:",chardata"`
		} `xml:"article"`
		CaseType *struct {
			Text string `xml:",chardata"`
		} `xml:"caseType"`
		Plural *struct {
			Text string `xml:",chardata"`
		} `xml:"plural"`
		PossessivePlural *struct {
			Text string `xml:",chardata"`
		} `xml:"possessivePlural"`
		Value *struct {
			Text string `xml:",chardata"`
		} `xml:"value"`
	} `xml:"caseValues"`
	FieldSets []struct {
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"fieldSets"`
	Gender *struct {
		Text string `xml:",chardata"`
	} `xml:"gender"`
	Layouts []struct {
		Layout struct {
			Text string `xml:",chardata"`
		} `xml:"layout"`
		Sections []struct {
			Label struct {
				Text string `xml:",chardata"`
			} `xml:"label"`
			Section struct {
				Text string `xml:",chardata"`
			} `xml:"section"`
		} `xml:"sections"`
	} `xml:"layouts"`
	NameFieldLabel *struct {
		Text string `xml:",chardata"`
	} `xml:"nameFieldLabel"`
	QuickActions []struct {
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"quickActions"`
	SharingReasons []struct {
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"sharingReasons"`
	StartsWith *struct {
		Text string `xml:",chardata"`
	} `xml:"startsWith"`
	WebLinks []struct {
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"webLinks"`
	WorkflowTasks []struct {
		Description *struct {
			Text string `xml:",chardata"`
		} `xml:"description"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Subject *struct {
			Text string `xml:",chardata"`
		} `xml:"subject"`
	} `xml:"workflowTasks"`
	Fields          FieldList          `xml:"fields"`
	RecordTypes     RecordTypeList     `xml:"recordTypes"`
	ValidationRules ValidationRuleList `xml:"validationRules"`
	// translationName is used when the object is created via composition
	translationName string
}

func (c *CustomObjectTranslation) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomObjectTranslation) Type() metadata.MetadataType {
	return OBJECT_TRANSLATIONS_NAME
}

func (c *CustomObjectTranslation) Files(format metadata.Format) (map[string][]byte, error) {
	if format == metadata.SourceFormat {
		// For source format, decompose the object translation into separate files
		return c.decompose()
	}

	// For metadata format, use default behavior (single file)
	files := make(map[string][]byte)
	content, err := internal.Marshal(c)
	if err != nil {
		return nil, err
	}

	// Use the stored translation name if MetadataInfo.Name() is empty
	translationName := string(c.MetadataInfo.Name())
	if translationName == "" && c.translationName != "" {
		translationName = c.translationName
	}
	fileName := "objectTranslations/" + translationName + ".objectTranslation"
	files[fileName] = content

	return files, nil
}

func (c *CustomObjectTranslation) decompose() (map[string][]byte, error) {
	files := make(map[string][]byte)
	// Use the stored translation name if MetadataInfo.Name() is empty
	translationName := string(c.MetadataInfo.Name())
	if translationName == "" && c.translationName != "" {
		translationName = c.translationName
	}
	baseDir := "objectTranslations/" + translationName

	// Clone the translation and remove only field translations that will be written as separate files
	// Keep recordTypes and validationRules in the base file per SFDX convention
	minimalTranslation := *c
	minimalTranslation.XMLName = xml.Name{Local: "CustomObjectTranslation"}
	minimalTranslation.Xmlns = "http://soap.sforce.com/2006/04/metadata"

	// Only clear out fields - recordTypes and validationRules remain in base file
	minimalTranslation.Fields = nil

	// Marshal the minimal translation
	content, err := internal.Marshal(&minimalTranslation)
	if err != nil {
		return nil, err
	}
	files[baseDir+"/"+translationName+".objectTranslation-meta.xml"] = content

	// Write field translations as separate files
	for _, f := range c.Fields {
		fieldTransMeta := &field.CustomFieldTranslation{
			MetadataInfo: metadata.MetadataInfo{},
			XMLName:      xml.Name{Local: "CustomFieldTranslation"},
			Xmlns:        "http://soap.sforce.com/2006/04/metadata",
			Field:        f,
		}
		fieldContent, err := internal.Marshal(fieldTransMeta)
		if err != nil {
			return nil, err
		}
		fieldName := f.Name.Text
		if fieldName == "" {
			continue
		}
		files[baseDir+"/"+fieldName+".fieldTranslation-meta.xml"] = fieldContent
	}

	return files, nil
}

func Open(path string) (*CustomObjectTranslation, error) {
	p := &CustomObjectTranslation{}
	return p, metadata.ParseMetadataXml(p, path)
}
