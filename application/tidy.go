package application

import "sort"

func (a *CustomApplication) Tidy() {
	a.ProfileActionOverrides.Tidy()
}

func (pao ProfileActionOverrideList) Tidy() {
	sort.Slice(pao, func(i, j int) bool {
		if pao[i].PageOrSobjectType != pao[j].PageOrSobjectType {
			return pao[i].PageOrSobjectType < pao[j].PageOrSobjectType
		}
		if pao[i].ActionName != pao[j].ActionName {
			return pao[i].ActionName < pao[j].ActionName
		}
		if pao[i].RecordType != nil && pao[j].RecordType != nil && *pao[i].RecordType != *pao[j].RecordType {
			return *pao[i].RecordType < *pao[j].RecordType
		}
		if pao[i].FormFactor != pao[j].FormFactor {
			return pao[i].FormFactor > pao[j].FormFactor
		}
		if pao[i].Profile != pao[j].Profile {
			return pao[i].Profile < pao[j].Profile
		}
		if pao[i].Type != pao[j].Type {
			return pao[i].Type < pao[j].Type
		}
		return pao[i].Content < pao[j].Content
	})
}
