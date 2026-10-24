package application

import (
	"reflect"
	"testing"
)

func strPtr(s string) *string {
	return &s
}

// Overrides with and without a record type sort the same way whatever their
// starting order, so tidying an application twice gives the same result.
func TestTidyProfileActionOverridesOrderDoesNotDependOnInput(t *testing.T) {
	override := func(recordType *string, profile string) ProfileActionOverride {
		return ProfileActionOverride{
			ActionName:        "View",
			Content:           strPtr("Campaign_Lightning1"),
			FormFactor:        "Large",
			PageOrSobjectType: "Campaign",
			RecordType:        recordType,
			Type:              "Flexipage",
			Profile:           profile,
		}
	}
	want := ProfileActionOverrideList{
		override(nil, "Admin"),
		override(nil, "Read Only"),
		override(strPtr("Campaign.Marketing_User"), "Admin"),
		override(strPtr("Campaign.Marketing_User"), "Read Only"),
		override(strPtr("Campaign.Outreach_User"), "Admin"),
	}
	orders := [][]int{
		{0, 1, 2, 3, 4},
		{4, 3, 2, 1, 0},
		{2, 0, 4, 1, 3},
		{3, 4, 1, 2, 0},
	}
	for _, order := range orders {
		var list ProfileActionOverrideList
		for _, i := range order {
			list = append(list, want[i])
		}
		list.Tidy()
		if !reflect.DeepEqual(list, want) {
			t.Errorf("order %v sorted to %v", order, list)
		}
	}
}

// An override without content does not stop the list from being sorted.
func TestTidyProfileActionOverridesWithoutContent(t *testing.T) {
	list := ProfileActionOverrideList{
		{ActionName: "View", FormFactor: "Large", PageOrSobjectType: "Account", Type: "Default", Profile: "Admin", Content: strPtr("Account_Page")},
		{ActionName: "View", FormFactor: "Large", PageOrSobjectType: "Account", Type: "Default", Profile: "Admin"},
	}
	list.Tidy()
	if list[0].Content != nil {
		t.Errorf("override without content sorted after one with content: %v", list)
	}
}
