package permissionset

import (
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

var CustomMetadataTypeExistsError = errors.New("custom metadata type already exists")

func (p *PermissionSet) AddCustomMetadataType(metadataType string) error {
	for _, c := range p.CustomMetadataTypeAccesses {
		if c.Name == metadataType {
			return CustomMetadataTypeExistsError
		}
	}
	p.CustomMetadataTypeAccesses = append(p.CustomMetadataTypeAccesses, CustomMetadataType{
		Name:    metadataType,
		Enabled: TrueText,
	})
	p.CustomMetadataTypeAccesses.Tidy()
	return nil
}

func (p *PermissionSet) DeleteCustomMetadataType(metadataType string) error {
	found := false
	newTypes := p.CustomMetadataTypeAccesses[:0]
	for _, f := range p.CustomMetadataTypeAccesses {
		if f.Name == metadataType {
			found = true
		} else {
			newTypes = append(newTypes, f)
		}
	}
	if !found {
		return errors.New("metadata type not found")
	}
	p.CustomMetadataTypeAccesses = newTypes
	return nil
}

func (p *PermissionSet) GetCustomMetadataTypes() CustomMetadataTypeList {
	return p.CustomMetadataTypeAccesses
}
