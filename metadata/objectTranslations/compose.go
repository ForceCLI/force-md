package objectTranslations

import (
	"encoding/xml"
	"strings"

	"github.com/ForceCLI/force-md/metadata"
	objTransField "github.com/ForceCLI/force-md/metadata/objectTranslations/field"
)

// ChildComponentProvider provides access to child components for composition
type ChildComponentProvider interface {
	Items(metadataType string) []metadata.RegisterableMetadata
}

// ComposeFromChildren creates a composed CustomObjectTranslation by merging child components from a provider.
// This is used when converting from source format (where components are separate) to metadata format.
// The returned CustomObjectTranslation needs to have SetName called before using Files() method.
func ComposeFromChildren(objectTranslationName string, provider ChildComponentProvider) *CustomObjectTranslation {
	// Try to get the base object translation if it exists
	var objTrans *CustomObjectTranslation
	baseObjectTransExists := false

	objectTranslationItems := provider.Items("CustomObjectTranslation")
	for _, item := range objectTranslationItems {
		if string(item.GetMetadataInfo().Name()) == objectTranslationName {
			if customObjTrans, ok := item.(*CustomObjectTranslation); ok {
				// Clone the object translation to avoid modifying the original
				cloned := *customObjTrans
				objTrans = &cloned
				baseObjectTransExists = true
				break
			}
		}
	}

	// If no base object translation exists, create a minimal one
	if !baseObjectTransExists {
		objTrans = &CustomObjectTranslation{
			XMLName:         xml.Name{Local: "CustomObjectTranslation"},
			Xmlns:           "http://soap.sforce.com/2006/04/metadata",
			MetadataInfo:    metadata.MetadataInfo{},
			translationName: objectTranslationName,
		}
	} else {
		// Don't clear anything - we want to merge additional items from separate files
		// with what's already in the base object translation
		// Ensure proper namespace is set
		objTrans.Xmlns = "http://soap.sforce.com/2006/04/metadata"
		// Ensure the translation name is set in case MetadataInfo doesn't have it
		if string(objTrans.MetadataInfo.Name()) == "" {
			objTrans.translationName = objectTranslationName
		}
	}

	// Helper function to find items with a specific prefix
	findItemsByPrefix := func(items []metadata.RegisterableMetadata, prefix string) []metadata.RegisterableMetadata {
		var result []metadata.RegisterableMetadata
		for _, item := range items {
			name := string(item.GetMetadataInfo().Name())
			if strings.HasPrefix(name, prefix) {
				result = append(result, item)
			}
		}
		return result
	}

	prefix := objectTranslationName + "."

	// Merge in field translations (only fields are decomposed, recordTypes and validationRules stay in base file)
	for _, ftMeta := range findItemsByPrefix(provider.Items("CustomFieldTranslation"), prefix) {
		if ft, ok := ftMeta.(*objTransField.CustomFieldTranslation); ok {
			objTrans.Fields = append(objTrans.Fields, ft.Field)
		}
	}

	// Tidy the composed object translation to ensure consistent ordering
	objTrans.Tidy()

	return objTrans
}
