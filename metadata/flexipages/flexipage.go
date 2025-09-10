package flexipage

import (
	"encoding/xml"
	"fmt"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "FlexiPage"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type FlexiPage struct {
	metadata.MetadataInfo
	XMLName          xml.Name          `xml:"FlexiPage"`
	Xmlns            string            `xml:"xmlns,attr"`
	Description      *TextLiteral      `xml:"description"`
	FlexiPageRegions []FlexiPageRegion `xml:"flexiPageRegions"`
	MasterLabel      TextLiteral       `xml:"masterLabel"`
	ParentFlexiPage  *TextLiteral      `xml:"parentFlexiPage"`
	SobjectType      *TextLiteral      `xml:"sobjectType"`
	Template         Template          `xml:"template"`
	FlexiPageType    TextLiteral       `xml:"type"`
}

func (c *FlexiPage) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *FlexiPage) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*FlexiPage, error) {
	p := &FlexiPage{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (p *FlexiPage) DeleteField(fieldName string) error {
	fieldDeleted := false

	// Iterate through flexipage regions
	for regionIdx := range p.FlexiPageRegions {
		region := &p.FlexiPageRegions[regionIdx]

		// Filter out the field from item instances
		filteredInstances := region.ItemInstances[:0]
		for _, instance := range region.ItemInstances {
			// Check if this is a field instance with the field we want to delete
			if instance.FieldInstance != nil {
				fieldItemText := instance.FieldInstance.FieldItem.Text

				// Match exact field name or field name with Record. prefix
				if fieldItemText == fieldName || fieldItemText == "Record."+fieldName {
					fieldDeleted = true
					// Skip this instance (don't add to filteredInstances)
					continue
				}
			}
			filteredInstances = append(filteredInstances, instance)
		}

		region.ItemInstances = filteredInstances
	}

	if !fieldDeleted {
		return fmt.Errorf("field '%s' not found in flexipage", fieldName)
	}

	return nil
}
