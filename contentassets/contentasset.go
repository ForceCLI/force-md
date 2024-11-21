package contentasset

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "ContentAsset"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type ContentAsset struct {
	internal.MetadataInfo
	XMLName                  xml.Name `xml:"ContentAsset"`
	Xmlns                    string   `xml:"xmlns,attr"`
	IsVisibleByExternalUsers struct {
		Text string `xml:",chardata"`
	} `xml:"isVisibleByExternalUsers"`
	Language struct {
		Text string `xml:",chardata"`
	} `xml:"language"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Relationships struct {
		Organization struct {
			Access struct {
				Text string `xml:",chardata"`
			} `xml:"access"`
		} `xml:"organization"`
	} `xml:"relationships"`
	Versions struct {
		Version struct {
			Number struct {
				Text string `xml:",chardata"`
			} `xml:"number"`
			PathOnClient struct {
				Text string `xml:",chardata"`
			} `xml:"pathOnClient"`
		} `xml:"version"`
	} `xml:"versions"`
}

func (c *ContentAsset) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ContentAsset) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*ContentAsset, error) {
	p := &ContentAsset{}
	return p, internal.ParseMetadataXml(p, path)
}
