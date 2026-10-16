package skills

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Skill"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Skill struct {
	metadata.MetadataInfo
	XMLName     xml.Name          `xml:"Skill"`
	Xmlns       string            `xml:"xmlns,attr"`
	Assignments *SkillAssignments `xml:"assignments"`
	Description *TextLiteral      `xml:"description"`
	Label       *TextLiteral      `xml:"label"`
}

type SkillAssignments struct {
	Profiles *SkillProfileAssignments `xml:"profiles"`
	Users    *SkillUserAssignments    `xml:"users"`
}

type SkillProfileAssignments struct {
	Profile []TextLiteral `xml:"profile"`
}

type SkillUserAssignments struct {
	User []TextLiteral `xml:"user"`
}

func (s *Skill) SetMetadata(m metadata.MetadataInfo) {
	s.MetadataInfo = m
}

func (s *Skill) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Skill, error) {
	s := &Skill{}
	return s, metadata.ParseMetadataXml(s, path)
}
