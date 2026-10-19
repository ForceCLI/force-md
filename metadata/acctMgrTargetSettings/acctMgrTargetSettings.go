package acctMgrTargetSettings

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "AcctMgrTargetSettings"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// AcctMgrTargetSettings holds the Manufacturing Cloud settings for account
// manager targets: the distribution period, the team member hierarchy and
// the default price book. An org has one, stored as
// acctMgrTargetSettings/AcctMgrTargetSettings.acctMgrTargetSetting.
type AcctMgrTargetSettings struct {
	metadata.MetadataInfo
	XMLName                      xml.Name     `xml:"AcctMgrTargetSettings"`
	Xmlns                        string       `xml:"xmlns,attr"`
	AcctMgrPeriodicTargetDstrCnt *IntegerText `xml:"acctMgrPeriodicTargetDstrCnt"`
	PeriodType                   *TextLiteral `xml:"periodType"`
	PricebookId                  *TextLiteral `xml:"pricebookId"`
	TeamMemberHierarchyType      *TextLiteral `xml:"teamMemberHierarchyType"`
}

func (s *AcctMgrTargetSettings) SetMetadata(m metadata.MetadataInfo) {
	s.MetadataInfo = m
}

func (s *AcctMgrTargetSettings) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*AcctMgrTargetSettings, error) {
	s := &AcctMgrTargetSettings{}
	return s, metadata.ParseMetadataXml(s, path)
}
