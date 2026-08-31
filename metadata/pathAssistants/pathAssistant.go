package pathassistant

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "PathAssistant"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type PathAssistant struct {
	metadata.MetadataInfo
	XMLName            xml.Name            `xml:"PathAssistant"`
	Xmlns              string              `xml:"xmlns,attr"`
	Active             BooleanText         `xml:"active"`
	EntityName         TextLiteral         `xml:"entityName"`
	FieldName          TextLiteral         `xml:"fieldName"`
	MasterLabel        TextLiteral         `xml:"masterLabel"`
	PathAssistantSteps []PathAssistantStep `xml:"pathAssistantSteps"`
	RecordTypeName     *TextLiteral        `xml:"recordTypeName"`
}

type PathAssistantStep struct {
	FieldNames        []TextLiteral `xml:"fieldNames"`
	Info              *TextLiteral  `xml:"info"`
	PicklistValueName TextLiteral   `xml:"picklistValueName"`
}

func (p *PathAssistant) SetMetadata(m metadata.MetadataInfo) {
	p.MetadataInfo = m
}

func (p *PathAssistant) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*PathAssistant, error) {
	p := &PathAssistant{}
	return p, metadata.ParseMetadataXml(p, path)
}
