package testSuites

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ApexTestSuite"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ApexTestSuite struct {
	metadata.MetadataInfo
	XMLName        xml.Name      `xml:"ApexTestSuite"`
	Xmlns          string        `xml:"xmlns,attr"`
	TestClassNames []TextLiteral `xml:"testClassName"`
}

func (s *ApexTestSuite) SetMetadata(m metadata.MetadataInfo) {
	s.MetadataInfo = m
}

func (s *ApexTestSuite) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ApexTestSuite, error) {
	s := &ApexTestSuite{}
	return s, metadata.ParseMetadataXml(s, path)
}
