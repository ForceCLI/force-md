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
		if recordType(pao[i]) != recordType(pao[j]) {
			// Overrides without a record type sort first
			return recordType(pao[i]) < recordType(pao[j])
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
		return content(pao[i]) < content(pao[j])
	})
}

func recordType(o ProfileActionOverride) string {
	if o.RecordType == nil {
		return ""
	}
	return *o.RecordType
}

func content(o ProfileActionOverride) string {
	if o.Content == nil {
		return ""
	}
	return *o.Content
}
