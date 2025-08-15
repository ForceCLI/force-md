package labels

import (
	"testing"
)

func TestCustomLabels_Tidy(t *testing.T) {
	labels := &CustomLabels{
		Labels: CustomLabelList{
			{
				FullName:         "ZTestLabel",
				Language:         "en_US",
				Protected:        "false",
				ShortDescription: "Z Test Label",
				Value:            "Z Value",
			},
			{
				FullName:         "ATestLabel",
				Language:         "en_US",
				Protected:        "false",
				ShortDescription: "A Test Label",
				Value:            "A Value",
			},
			{
				FullName:         "MidTestLabel",
				Language:         "en_US",
				Protected:        "false",
				ShortDescription: "Mid Test Label",
				Value:            "Mid Value",
			},
		},
	}

	// Verify initial order (unsorted)
	if labels.Labels[0].FullName != "ZTestLabel" {
		t.Errorf("Expected first label to be ZTestLabel, got %s", labels.Labels[0].FullName)
	}
	if labels.Labels[1].FullName != "ATestLabel" {
		t.Errorf("Expected second label to be ATestLabel, got %s", labels.Labels[1].FullName)
	}
	if labels.Labels[2].FullName != "MidTestLabel" {
		t.Errorf("Expected third label to be MidTestLabel, got %s", labels.Labels[2].FullName)
	}

	// Tidy (sort)
	labels.Tidy()

	// Verify sorted order
	if labels.Labels[0].FullName != "ATestLabel" {
		t.Errorf("Expected first label after tidy to be ATestLabel, got %s", labels.Labels[0].FullName)
	}
	if labels.Labels[1].FullName != "MidTestLabel" {
		t.Errorf("Expected second label after tidy to be MidTestLabel, got %s", labels.Labels[1].FullName)
	}
	if labels.Labels[2].FullName != "ZTestLabel" {
		t.Errorf("Expected third label after tidy to be ZTestLabel, got %s", labels.Labels[2].FullName)
	}
}

func TestCustomLabels_TidyEmptyLabels(t *testing.T) {
	labels := &CustomLabels{
		Labels: CustomLabelList{},
	}

	// Should not panic on empty labels
	labels.Tidy()

	if len(labels.Labels) != 0 {
		t.Errorf("Expected labels to remain empty, got %d labels", len(labels.Labels))
	}
}
