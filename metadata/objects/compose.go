package objects

import (
	"encoding/xml"
	"strings"

	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objects/businessprocess"
	"github.com/ForceCLI/force-md/metadata/objects/compactlayout"
	"github.com/ForceCLI/force-md/metadata/objects/field"
	"github.com/ForceCLI/force-md/metadata/objects/fieldset"
	"github.com/ForceCLI/force-md/metadata/objects/index"
	"github.com/ForceCLI/force-md/metadata/objects/listview"
	"github.com/ForceCLI/force-md/metadata/objects/recordtype"
	"github.com/ForceCLI/force-md/metadata/objects/sharingreason"
	"github.com/ForceCLI/force-md/metadata/objects/validationrule"
	"github.com/ForceCLI/force-md/metadata/objects/weblink"
)

// ChildComponentProvider provides access to child components for composition
type ChildComponentProvider interface {
	Items(metadataType string) []metadata.RegisterableMetadata
}

// ComposeFromChildren creates a composed CustomObject by merging child components from a provider.
// This is used when converting from source format (where components are separate) to metadata format.
func ComposeFromChildren(objectName string, provider ChildComponentProvider) *CustomObject {
	// Try to get the base object if it exists
	var obj *CustomObject
	baseObjectExists := false

	objectItems := provider.Items("CustomObject")
	for _, item := range objectItems {
		if string(item.GetMetadataInfo().Name()) == objectName {
			if customObj, ok := item.(*CustomObject); ok {
				// Clone the object to avoid modifying the original
				cloned := *customObj
				obj = &cloned
				baseObjectExists = true
				break
			}
		}
	}

	// If no base object exists, create a minimal one
	if !baseObjectExists {
		obj = &CustomObject{
			XMLName: xml.Name{Local: "CustomObject"},
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
		}
	} else {
		// Clear child component arrays since we'll rebuild them from the provider
		obj.Fields = nil
		obj.RecordTypes = nil
		obj.ValidationRules = nil
		obj.Indexes = nil
		obj.FieldSets = nil
		obj.WebLinks = nil
		obj.CompactLayouts = nil
		obj.SharingReasons = nil
		obj.BusinessProcesses = nil
		obj.ListViews = nil
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

	prefix := objectName + "."

	// Merge in business processes
	for _, bpMeta := range findItemsByPrefix(provider.Items("BusinessProcess"), prefix) {
		if bp, ok := bpMeta.(*businessprocess.BusinessProcessMetadata); ok {
			obj.BusinessProcesses = append(obj.BusinessProcesses, bp.BusinessProcess)
		}
	}

	// Merge in fields
	for _, fieldMeta := range findItemsByPrefix(provider.Items("CustomField"), prefix) {
		if f, ok := fieldMeta.(*field.CustomField); ok {
			obj.Fields = append(obj.Fields, f.Field)
		}
	}

	// Merge in record types
	for _, rtMeta := range findItemsByPrefix(provider.Items("RecordType"), prefix) {
		if rt, ok := rtMeta.(*recordtype.RecordTypeMetadata); ok {
			obj.RecordTypes = append(obj.RecordTypes, rt.RecordType)
		}
	}

	// Merge in validation rules
	for _, vrMeta := range findItemsByPrefix(provider.Items("ValidationRule"), prefix) {
		if vr, ok := vrMeta.(*validationrule.ValidationRule); ok {
			obj.ValidationRules = append(obj.ValidationRules, vr.Rule)
		}
	}

	// Merge in field sets
	for _, fsMeta := range findItemsByPrefix(provider.Items("FieldSet"), prefix) {
		if fs, ok := fsMeta.(*fieldset.FieldSetMetadata); ok {
			obj.FieldSets = append(obj.FieldSets, fs.FieldSet)
		}
	}

	// Merge in web links
	for _, wlMeta := range findItemsByPrefix(provider.Items("WebLink"), prefix) {
		if wl, ok := wlMeta.(*weblink.WebLinkMetadata); ok {
			obj.WebLinks = append(obj.WebLinks, wl.WebLink)
		}
	}

	// Merge in compact layouts
	for _, clMeta := range findItemsByPrefix(provider.Items("CompactLayout"), prefix) {
		if cl, ok := clMeta.(*compactlayout.CompactLayoutMetadata); ok {
			obj.CompactLayouts = append(obj.CompactLayouts, cl.CompactLayout)
		}
	}

	// Merge in sharing reasons
	for _, srMeta := range findItemsByPrefix(provider.Items("SharingReason"), prefix) {
		if sr, ok := srMeta.(*sharingreason.SharingReasonMetadata); ok {
			obj.SharingReasons = append(obj.SharingReasons, sr.SharingReason)
		}
	}

	// Merge in indexes
	for _, idxMeta := range findItemsByPrefix(provider.Items("Index"), prefix) {
		if idx, ok := idxMeta.(*index.Index); ok {
			obj.Indexes = append(obj.Indexes, idx.BigObjectIndex)
		}
	}

	// Merge in list views
	for _, lvMeta := range findItemsByPrefix(provider.Items("ListView"), prefix) {
		if lv, ok := lvMeta.(*listview.ListViewMetadata); ok {
			obj.ListViews = append(obj.ListViews, lv.ListView)
		}
	}

	// Tidy the composed object to ensure consistent ordering
	obj.Tidy()

	return obj
}

