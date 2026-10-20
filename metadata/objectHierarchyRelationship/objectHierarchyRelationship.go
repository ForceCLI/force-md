package objectHierarchyRelationship

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ObjectHierarchyRelationship"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// ObjectHierarchyRelationship maps the fields of an input object to an output
// object for a Manufacturing Cloud usage: converting a record to a sales
// agreement (ConvertToSalesAgreement), a business transformation
// (TransformationMapping), and others. A record either carries a child
// object mapping for the input object's child records, or names the record
// it is the child of in ParentRecord. Stored as
// ObjectHierarchyRelationship/<name>.settings.
type ObjectHierarchyRelationship struct {
	metadata.MetadataInfo
	XMLName                        xml.Name       `xml:"ObjectHierarchyRelationship"`
	Xmlns                          string         `xml:"xmlns,attr"`
	ChildObjectMapping             *ObjectMapping `xml:"childObjectMapping"`
	InputObjRecordsGrpFieldName    *TextLiteral   `xml:"inputObjRecordsGrpFieldName"`
	MappingType                    *TextLiteral   `xml:"mappingType"`
	MasterLabel                    *TextLiteral   `xml:"masterLabel"`
	OutputPntRelationshipFieldName *TextLiteral   `xml:"outputPntRelationshipFieldName"`
	ParentObjectMapping            *ObjectMapping `xml:"parentObjectMapping"`
	ParentRecord                   *TextLiteral   `xml:"parentRecord"`
	ParentRelationshipFieldName    *TextLiteral   `xml:"parentRelationshipFieldName"`
	SourceReferenceRelaFieldName   *TextLiteral   `xml:"sourceReferenceRelaFieldName"`
	UsageType                      *TextLiteral   `xml:"usageType"`
}

// ObjectMapping maps the fields of an input object to those of an output
// object.
type ObjectMapping struct {
	InputObject   *TextLiteral         `xml:"inputObject"`
	MappingFields []ObjectMappingField `xml:"mappingFields"`
	OutputObject  *TextLiteral         `xml:"outputObject"`
}

// ObjectMappingField maps an input object field to an output object field.
// A lookup field is named by its relationship (Account for AccountId).
type ObjectMappingField struct {
	InputField  *TextLiteral `xml:"inputField"`
	OutputField *TextLiteral `xml:"outputField"`
}

func (r *ObjectHierarchyRelationship) SetMetadata(m metadata.MetadataInfo) {
	r.MetadataInfo = m
}

func (r *ObjectHierarchyRelationship) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ObjectHierarchyRelationship, error) {
	r := &ObjectHierarchyRelationship{}
	return r, metadata.ParseMetadataXml(r, path)
}
