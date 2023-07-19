package sharingrules

import "sort"

func (r *SharingRules) Tidy() {
	r.SharingCriteriaRules.Tidy()
	r.SharingOwnerRules.Tidy()
}

func (r CriteriaRuleList) Tidy() {
	sort.Slice(r, func(i, j int) bool {
		return r[i].FullName < r[j].FullName
	})
}

func (r OwnerRuleList) Tidy() {
	sort.Slice(r, func(i, j int) bool {
		return r[i].FullName < r[j].FullName
	})
}
