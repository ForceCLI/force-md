package recordtype

import "sort"

func (values ValueSetOptionList) Tidy() {
	sort.Slice(values, func(i, j int) bool {
		return values[i].FullName < values[j].FullName
	})
}

func (picklists PicklistList) Tidy() {
	sort.Slice(picklists, func(i, j int) bool {
		return picklists[i].Picklist < picklists[j].Picklist
	})
}

// Tidy implements the general.Tidyable interface for standalone RecordType metadata.
// It sorts the list of picklists and then sorts each picklist's values.
func (r *RecordTypeMetadata) Tidy() {
	// sort the picklist entries by picklist (field) name
	r.PicklistValues.Tidy()
	// sort values within each picklist
	for i := range r.PicklistValues {
		// sort the values of this picklist
		r.PicklistValues[i].Values.Tidy()
	}
}
