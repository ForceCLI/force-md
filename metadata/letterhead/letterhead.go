package letterhead

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Letterhead"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Letterhead struct {
	metadata.MetadataInfo
	XMLName   xml.Name `xml:"Letterhead"`
	Xmlns     string   `xml:"xmlns,attr"`
	Available struct {
		Text string `xml:",chardata"`
	} `xml:"available"`
	BackgroundColor struct {
		Text string `xml:",chardata"`
	} `xml:"backgroundColor"`
	BodyColor struct {
		Text string `xml:",chardata"`
	} `xml:"bodyColor"`
	BottomLine struct {
		Color struct {
			Text string `xml:",chardata"`
		} `xml:"color"`
		Height struct {
			Text string `xml:",chardata"`
		} `xml:"height"`
	} `xml:"bottomLine"`
	Footer struct {
		BackgroundColor struct {
			Text string `xml:",chardata"`
		} `xml:"backgroundColor"`
		Height struct {
			Text string `xml:",chardata"`
		} `xml:"height"`
		HorizontalAlignment struct {
			Text string `xml:",chardata"`
		} `xml:"horizontalAlignment"`
		VerticalAlignment struct {
			Text string `xml:",chardata"`
		} `xml:"verticalAlignment"`
	} `xml:"footer"`
	Header struct {
		BackgroundColor struct {
			Text string `xml:",chardata"`
		} `xml:"backgroundColor"`
		Height struct {
			Text string `xml:",chardata"`
		} `xml:"height"`
		HorizontalAlignment struct {
			Text string `xml:",chardata"`
		} `xml:"horizontalAlignment"`
		Logo struct {
			Text string `xml:",chardata"`
		} `xml:"logo"`
		VerticalAlignment struct {
			Text string `xml:",chardata"`
		} `xml:"verticalAlignment"`
	} `xml:"header"`
	MiddleLine struct {
		Color struct {
			Text string `xml:",chardata"`
		} `xml:"color"`
		Height struct {
			Text string `xml:",chardata"`
		} `xml:"height"`
	} `xml:"middleLine"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	TopLine struct {
		Color struct {
			Text string `xml:",chardata"`
		} `xml:"color"`
		Height struct {
			Text string `xml:",chardata"`
		} `xml:"height"`
	} `xml:"topLine"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *Letterhead) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Letterhead) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Letterhead, error) {
	p := &Letterhead{}
	return p, metadata.ParseMetadataXml(p, path)
}
