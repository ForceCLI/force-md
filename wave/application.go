package wave

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const APP_NAME = "WaveApplication"

func init() {
	internal.TypeRegistry.Register(APP_NAME, func(path string) (internal.RegisterableMetadata, error) { return OpenApplication(path) })
}

type WaveApplication struct {
	internal.MetadataInfo
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

func (c *WaveApplication) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *WaveApplication) Type() internal.MetadataType {
	return APP_NAME
}

func OpenApplication(path string) (*WaveApplication, error) {
	p := &WaveApplication{}
	return p, internal.ParseMetadataXml(p, path)
}
