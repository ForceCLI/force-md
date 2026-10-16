package territory2model

import (
	"encoding/xml"
	"fmt"
	"path/filepath"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "Territory2Model"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// FieldValue is the value of a custom field on a territory or territory
// model.
type FieldValue struct {
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Value TypedValue `xml:"value"`
}

// TypedValue is a custom field value written with its XML schema type,
// <value xsi:type="xsd:string">, or as <value xsi:nil="true"/> when the field
// has no value.
type TypedValue struct {
	Text string `xml:",innerxml"`
	Type string `xml:"xsi:type,attr,omitempty"`
	Nil  string `xml:"xsi:nil,attr,omitempty"`
}

type Territory2Model struct {
	metadata.MetadataInfo
	XMLName      xml.Name     `xml:"Territory2Model"`
	Xmlns        string       `xml:"xmlns,attr"`
	Xsi          string       `xml:"xmlns:xsi,attr,omitempty"`
	Xsd          string       `xml:"xmlns:xsd,attr,omitempty"`
	CustomFields []FieldValue `xml:"customFields"`
	Description  *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	RecordVisibility *struct {
		Text string `xml:",chardata"`
	} `xml:"recordVisibility"`
	UserRecordAccess *struct {
		Text string `xml:",chardata"`
	} `xml:"userRecordAccess"`
}

func (c *Territory2Model) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Territory2Model) Type() metadata.MetadataType {
	return NAME
}

// Files writes the model into its own directory, which also holds its
// territories and rules: territory2Models/<Model>/<Model>.territory2Model.
func (c *Territory2Model) Files(format metadata.Format) (map[string][]byte, error) {
	name := string(c.MetadataInfo.Name())
	return ModelComponentFiles(c, name, "", name, NAME, format)
}

// ModelName returns the DeveloperName of the model that owns a territory or
// rule file, the directory above its territories or rules directory:
// territory2Models/<Model>/territories/<Name>.territory2.
func ModelName(path string) string {
	return filepath.Base(filepath.Dir(filepath.Dir(path)))
}

// ComponentName returns the package.xml member name of a territory or rule
// file, <Model>.<Name>.
func ComponentName(path string) metadata.MetadataObjectName {
	return metadata.MetadataObjectName(ModelName(path) + "." + string(metadata.NameFromPath(path)))
}

// ModelComponentFiles returns the file for a component of the model named
// model, territory2Models/<model>/<dir>/<name>.<suffix>, with -meta.xml
// appended in source format. An empty dir places the file directly in the
// model's directory.
func ModelComponentFiles(m metadata.RegisterableMetadata, model, dir, name string, metadataType metadata.MetadataType, format metadata.Format) (map[string][]byte, error) {
	xmlContent, err := internal.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %s: %w", metadataType, err)
	}
	fileName := name + "." + registry.GetMetadataSuffix(metadataType)
	switch format {
	case metadata.SourceFormat:
		fileName += "-meta.xml"
	case metadata.MetadataFormat:
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}
	path := filepath.Join(registry.GetMetadataDirectory(NAME), model, dir, fileName)
	return map[string][]byte{path: xmlContent}, nil
}

func Open(path string) (*Territory2Model, error) {
	p := &Territory2Model{}
	return p, metadata.ParseMetadataXml(p, path)
}
