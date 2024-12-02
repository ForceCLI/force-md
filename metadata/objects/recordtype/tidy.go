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
