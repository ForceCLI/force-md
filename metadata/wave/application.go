package wave

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const APP_NAME = "WaveApplication"

func init() {
	internal.TypeRegistry.Register(APP_NAME, func(path string) (metadata.RegisterableMetadata, error) { return OpenApplication(path) })
}

type WaveApplication struct {
	metadata.MetadataInfo
	XMLName   xml.Name `xml:"WaveApplication"`
	Xmlns     string   `xml:"xmlns,attr"`
	AssetIcon struct {
		Text string `xml:",chardata"`
	} `xml:"assetIcon"`
	Folder struct {
		Text string `xml:",chardata"`
	} `xml:"folder"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Shares struct {
		AccessLevel struct {
			Text string `xml:",chardata"`
		} `xml:"accessLevel"`
		SharedTo struct {
			Text string `xml:",chardata"`
		} `xml:"sharedTo"`
		SharedToType struct {
			Text string `xml:",chardata"`
		} `xml:"sharedToType"`
	} `xml:"shares"`
	TemplateOrigin struct {
		Text string `xml:",chardata"`
	} `xml:"templateOrigin"`
	TemplateVersion struct {
		Text string `xml:",chardata"`
	} `xml:"templateVersion"`
}

func (c *WaveApplication) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *WaveApplication) Type() metadata.MetadataType {
	return APP_NAME
}

func OpenApplication(path string) (*WaveApplication, error) {
	p := &WaveApplication{}
	return p, metadata.ParseMetadataXml(p, path)
}
