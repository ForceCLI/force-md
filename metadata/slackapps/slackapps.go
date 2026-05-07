package slackapps

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "SlackApp"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// SlackApp represents the SlackApp metadata type. The accompanying
// .slackapp file (YAML, holding the app's command/shortcut/event handler
// definitions) sits next to the .slackapp-meta.xml file in the slackapps
// directory and is read separately by consumers.
type SlackApp struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"SlackApp"`
	Xmlns   string   `xml:"xmlns,attr"`

	AppKey struct {
		Text string `xml:",chardata"`
	} `xml:"appKey"`
	AppToken struct {
		Text string `xml:",chardata"`
	} `xml:"appToken"`
	BotScopes *struct {
		Text string `xml:",chardata"`
	} `xml:"botScopes,omitempty"`
	ClientKey struct {
		Text string `xml:",chardata"`
	} `xml:"clientKey"`
	ClientSecret struct {
		Text string `xml:",chardata"`
	} `xml:"clientSecret"`
	IsProtected *struct {
		Text string `xml:",chardata"`
	} `xml:"isProtected,omitempty"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	SigningSecret struct {
		Text string `xml:",chardata"`
	} `xml:"signingSecret"`
	UserScopes *struct {
		Text string `xml:",chardata"`
	} `xml:"userScopes,omitempty"`
}

func (s *SlackApp) SetMetadata(m metadata.MetadataInfo) {
	s.MetadataInfo = m
}

func Open(path string) (*SlackApp, error) {
	p := &SlackApp{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (s *SlackApp) Type() metadata.MetadataType {
	return NAME
}
