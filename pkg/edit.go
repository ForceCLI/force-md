package pkg

import (
	"fmt"
	"strings"
)

func (p *Package) Add(metadataType string, member string) error {
	for i, t := range p.Types {
		if strings.ToLower(t.Name) == strings.ToLower(metadataType) {
			for _, m := range t.Members {
				if strings.ToLower(m) == strings.ToLower(member) {
					return fmt.Errorf("%s of type %s already exists", member, metadataType)
				}
			}
			p.Types[i].Members = append(p.Types[i].Members, member)
			p.Types[i].Tidy()
			return nil
		}
	}
	p.Types = append(p.Types, MetadataItems{
		Name:    metadataType,
		Members: []string{member},
	})
	return nil
}

func (p *Package) Delete(metadataType string, member string) error {
	found := false
	for i, t := range p.Types {
		newMembers := t.Members[:0]
		if strings.ToLower(t.Name) == strings.ToLower(metadataType) {
			for _, m := range t.Members {
				if strings.ToLower(m) == strings.ToLower(member) {
					found = true
				} else {
					newMembers = append(newMembers, m)
				}
			}
			p.Types[i].Members = newMembers
		}
	}
	if !found {
		return fmt.Errorf("%s of type %s not found", member, metadataType)
	}
	return nil
}
