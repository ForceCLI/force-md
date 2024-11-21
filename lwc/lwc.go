package lwc

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "LightningComponentBundle"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type LightningComponentBundle struct {
	internal.MetadataInfo
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

func (c *LightningComponentBundle) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *LightningComponentBundle) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*LightningComponentBundle, error) {
	p := &LightningComponentBundle{}
	return p, internal.ParseMetadataXml(p, path)
}
