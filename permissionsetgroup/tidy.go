package permissionsetgroup

import "sort"

func (pl PermissionSetList) Tidy() {
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Text < pl[j].Text
	})
}
