package lwc

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "LightningComponentBundle"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type LightningComponentBundle struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"LightningComponentBundle"`
	Xmlns      string   `xml:"xmlns,attr"`
	Fqn        string   `xml:"fqn,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	IsExposed struct {
		Text string `xml:",chardata"`
	} `xml:"isExposed"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Targets struct {
		Target []struct {
			Text string `xml:",chardata"`
		} `xml:"target"`
	} `xml:"targets"`
	TargetConfigs struct {
		TargetConfig []struct {
			Targets  string `xml:"targets,attr"`
			Xmlns    string `xml:"xmlns,attr"`
			Property []struct {
				Name        string `xml:"name,attr"`
				Label       string `xml:"label,attr"`
				Type        string `xml:"type,attr"`
				Role        string `xml:"role,attr"`
				Description string `xml:"description,attr"`
				Datasource  string `xml:"datasource,attr"`
				Default     string `xml:"default,attr"`
				Required    string `xml:"required,attr"`
			} `xml:"property"`
			SupportedFormFactors struct {
				SupportedFormFactor []struct {
					Type string `xml:"type,attr"`
				} `xml:"supportedFormFactor"`
			} `xml:"supportedFormFactors"`
			ActionType struct {
				Text string `xml:",chardata"`
			} `xml:"actionType"`
		} `xml:"targetConfig"`
	} `xml:"targetConfigs"`
	RuntimeNamespace struct {
		Text string `xml:",chardata"`
	} `xml:"runtimeNamespace"`
}

func (c *LightningComponentBundle) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *LightningComponentBundle) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*LightningComponentBundle, error) {
	p := &LightningComponentBundle{}
	return p, metadata.ParseMetadataXml(p, path)
}
