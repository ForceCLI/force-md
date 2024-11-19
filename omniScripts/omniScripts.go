package omniscript

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "OmniScript"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type OmniScript struct {
	internal.MetadataInfo
	XMLName                     xml.Name `xml:"OmniScript"`
	Xmlns                       string   `xml:"xmlns,attr"`
	ElementTypeComponentMapping struct {
		Text string `xml:",chardata"`
	} `xml:"elementTypeComponentMapping"`
	IsActive struct {
		Text string `xml:",chardata"`
	} `xml:"isActive"`
	IsIntegrationProcedure struct {
		Text string `xml:",chardata"`
	} `xml:"isIntegrationProcedure"`
	IsMetadataCacheDisabled struct {
		Text string `xml:",chardata"`
	} `xml:"isMetadataCacheDisabled"`
	IsOmniScriptEmbeddable struct {
		Text string `xml:",chardata"`
	} `xml:"isOmniScriptEmbeddable"`
	IsTestProcedure struct {
		Text string `xml:",chardata"`
	} `xml:"isTestProcedure"`
	IsWebCompEnabled struct {
		Text string `xml:",chardata"`
	} `xml:"isWebCompEnabled"`
	Language struct {
		Text string `xml:",chardata"`
	} `xml:"language"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	OmniProcessElements []struct {
		IsActive struct {
			Text string `xml:",chardata"`
		} `xml:"isActive"`
		IsOmniScriptEmbeddable struct {
			Text string `xml:",chardata"`
		} `xml:"isOmniScriptEmbeddable"`
		Level struct {
			Text string `xml:",chardata"`
		} `xml:"level"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		OmniProcessVersionNumber struct {
			Text string `xml:",chardata"`
		} `xml:"omniProcessVersionNumber"`
		PropertySetConfig struct {
			Text string `xml:",chardata"`
		} `xml:"propertySetConfig"`
		SequenceNumber struct {
			Text string `xml:",chardata"`
		} `xml:"sequenceNumber"`
		Type struct {
			Text string `xml:",chardata"`
		} `xml:"type"`
		ChildElements []struct {
			ChildElements []struct {
				IsActive struct {
					Text string `xml:",chardata"`
				} `xml:"isActive"`
				IsOmniScriptEmbeddable struct {
					Text string `xml:",chardata"`
				} `xml:"isOmniScriptEmbeddable"`
				Level struct {
					Text string `xml:",chardata"`
				} `xml:"level"`
				Name struct {
					Text string `xml:",chardata"`
				} `xml:"name"`
				OmniProcessVersionNumber struct {
					Text string `xml:",chardata"`
				} `xml:"omniProcessVersionNumber"`
				PropertySetConfig struct {
					Text string `xml:",chardata"`
				} `xml:"propertySetConfig"`
				SequenceNumber struct {
					Text string `xml:",chardata"`
				} `xml:"sequenceNumber"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
				ChildElements []struct {
					IsActive struct {
						Text string `xml:",chardata"`
					} `xml:"isActive"`
					IsOmniScriptEmbeddable struct {
						Text string `xml:",chardata"`
					} `xml:"isOmniScriptEmbeddable"`
					Level struct {
						Text string `xml:",chardata"`
					} `xml:"level"`
					Name struct {
						Text string `xml:",chardata"`
					} `xml:"name"`
					OmniProcessVersionNumber struct {
						Text string `xml:",chardata"`
					} `xml:"omniProcessVersionNumber"`
					PropertySetConfig struct {
						Text string `xml:",chardata"`
					} `xml:"propertySetConfig"`
					SequenceNumber struct {
						Text string `xml:",chardata"`
					} `xml:"sequenceNumber"`
					Type struct {
						Text string `xml:",chardata"`
					} `xml:"type"`
				} `xml:"childElements"`
			} `xml:"childElements"`
			IsActive struct {
				Text string `xml:",chardata"`
			} `xml:"isActive"`
			IsOmniScriptEmbeddable struct {
				Text string `xml:",chardata"`
			} `xml:"isOmniScriptEmbeddable"`
			Level struct {
				Text string `xml:",chardata"`
			} `xml:"level"`
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			OmniProcessVersionNumber struct {
				Text string `xml:",chardata"`
			} `xml:"omniProcessVersionNumber"`
			PropertySetConfig struct {
				Text string `xml:",chardata"`
			} `xml:"propertySetConfig"`
			SequenceNumber struct {
				Text string `xml:",chardata"`
			} `xml:"sequenceNumber"`
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"type"`
		} `xml:"childElements"`
		Description struct {
			Text string `xml:",chardata"`
		} `xml:"description"`
	} `xml:"omniProcessElements"`
	OmniProcessType struct {
		Text string `xml:",chardata"`
	} `xml:"omniProcessType"`
	PropertySetConfig struct {
		Text string `xml:",chardata"`
	} `xml:"propertySetConfig"`
	SubType struct {
		Text string `xml:",chardata"`
	} `xml:"subType"`
	OmniScriptType struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	UniqueName struct {
		Text string `xml:",chardata"`
	} `xml:"uniqueName"`
	VersionNumber struct {
		Text string `xml:",chardata"`
	} `xml:"versionNumber"`
	WebComponentKey struct {
		Text string `xml:",chardata"`
	} `xml:"webComponentKey"`
}

func (c *OmniScript) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *OmniScript) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*OmniScript, error) {
	p := &OmniScript{}
	return p, internal.ParseMetadataXml(p, path)
}
