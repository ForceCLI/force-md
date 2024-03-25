package permissionset

import (
	"fmt"
	"sort"
)

func (p *PermissionSet) Tidy() {
	p.ApplicationVisibilities.Tidy()
	p.ClassAccesses.Tidy()
	p.CustomPermissions.Tidy()
	p.CustomMetadataTypeAccesses.Tidy()
	p.FieldPermissions.Tidy()
	p.ObjectPermissions.Tidy()
	p.PageAccesses.Tidy()
	p.RecordTypeVisibilities.Tidy()
	p.TabSettings.Tidy()
	p.UserPermissions.Tidy()
}

func (op ObjectPermissionsList) Tidy() {
	sort.Slice(op, func(i, j int) bool {
		return op[i].Object < op[j].Object
	})
}

func (av ApplicationVisibilityList) Tidy() {
	sort.Slice(av, func(i, j int) bool {
		return av[i].Application < av[j].Application
	})
}

func (ca ApexClassList) Tidy() {
	sort.Slice(ca, func(i, j int) bool {
		return ca[i].ApexClass < ca[j].ApexClass
	})
}

func (pa PageAccessList) Tidy() {
	sort.Slice(pa, func(i, j int) bool {
		return pa[i].ApexPage < pa[j].ApexPage
	})
}

func (ts TabSettingsList) Tidy() {
	sort.Slice(ts, func(i, j int) bool {
		return ts[i].Tab < ts[j].Tab
	})
}

func (rt RecordTypeList) Tidy() {
	sort.Slice(rt, func(i, j int) bool {
		return rt[i].RecordType < rt[j].RecordType
	})
}

func (cp CustomPermissionList) Tidy() {
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].Name < cp[j].Name
	})
}

func (cp CustomMetadataTypeList) Tidy() {
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].Name < cp[j].Name
	})
}

func (cp CustomSettingList) Tidy() {
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].Name < cp[j].Name
	})
}

func (fp *FieldPermissionsList) Tidy() {
	if len(*fp) == 0 {
		return
	}
	sort.Slice(*fp, func(i, j int) bool {
		return (*fp)[i].Field < (*fp)[j].Field
	})
	lastUniqueIndex := 0
	for i := 1; i < len(*fp); i++ {
		// If the current element is not a duplicate, move it to the next position after the last unique element
		if (*fp)[i].Field != (*fp)[lastUniqueIndex].Field {
			lastUniqueIndex++
			(*fp)[lastUniqueIndex] = (*fp)[i]
		} else {
			fmt.Println("omitting duplicate permissions for", (*fp)[i].Field)
		}
	}
	// Slice the original slice to the correct length of unique elements
	*fp = (*fp)[:lastUniqueIndex+1]
}

func (up UserPermissionList) Tidy() {
	sort.Slice(up, func(i, j int) bool {
		return up[i].Name < up[j].Name
	})
}
