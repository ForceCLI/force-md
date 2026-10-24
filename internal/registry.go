package internal

import (
	"strings"

	"github.com/ForceCLI/force-md/metadata"
)

type metadataTypeRegistry map[string]OpenMetadataFunc

var TypeRegistry metadataTypeRegistry

type suffixRegistration struct {
	suffix string
	open   OpenMetadataFunc
}

var suffixRegistrations []suffixRegistration

func init() {
	TypeRegistry = make(metadataTypeRegistry)
}

type OpenMetadataFunc func(path string) (metadata.RegisterableMetadata, error)

func (r *metadataTypeRegistry) Register(metadataType string, openFunc OpenMetadataFunc) {
	if _, ok := (*r)[metadataType]; ok {
		panic("Duplicate metadata registration for " + metadataType)
	}
	(*r)[metadataType] = openFunc
}

// RegisterSuffix registers openFunc for metadata whose root element name ends
// with suffix and that has no registration of its own, such as the
// <Feature>Settings types, which Salesforce adds to in each release.
func (r *metadataTypeRegistry) RegisterSuffix(suffix string, openFunc OpenMetadataFunc) {
	for _, s := range suffixRegistrations {
		if s.suffix == suffix {
			panic("Duplicate metadata suffix registration for " + suffix)
		}
	}
	suffixRegistrations = append(suffixRegistrations, suffixRegistration{suffix: suffix, open: openFunc})
}

// Lookup returns the function that opens metadata whose root element is
// named element: the one registered for that name, or else one registered
// for a suffix of it.
func (r metadataTypeRegistry) Lookup(element string) (OpenMetadataFunc, bool) {
	if open, ok := r[element]; ok {
		return open, true
	}
	for _, s := range suffixRegistrations {
		if strings.HasSuffix(element, s.suffix) && element != s.suffix {
			return s.open, true
		}
	}
	return nil, false
}
