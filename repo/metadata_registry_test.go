package repo

import (
	"testing"
)

func TestGetMetadataDirectory(t *testing.T) {
	tests := []struct {
		metadataType string
		expected     string
	}{
		{"customobject", "objects"},
		{"CustomObject", "objects"},
		{"auradefinitionbundle", "aura"},
		{"lightningcomponentbundle", "lwc"},
		{"CustomApplication", "applications"},
	}

	for _, tt := range tests {
		result := GetMetadataDirectory(tt.metadataType)
		if result != tt.expected {
			t.Errorf("GetMetadataDirectory(%s) = %s; want %s", tt.metadataType, result, tt.expected)
		}
	}
}

func TestGetParentType(t *testing.T) {
	tests := []struct {
		childType string
		expected  string
	}{
		{"customfield", "customobject"},
		{"CustomField", "customobject"},
		{"validationrule", "customobject"},
		{"recordtype", "customobject"},
		{"workflowalert", "workflow"},
		{"sharingcriteriarule", "sharingrules"},
		{"notachild", ""},
	}

	for _, tt := range tests {
		result := GetParentType(tt.childType)
		if result != tt.expected {
			t.Errorf("GetParentType(%s) = %s; want %s", tt.childType, result, tt.expected)
		}
	}
}

func TestIsChildType(t *testing.T) {
	tests := []struct {
		metadataType string
		expected     bool
	}{
		{"customfield", true},
		{"validationrule", true},
		{"recordtype", true},
		{"workflowalert", true},
		{"customobject", false},
		{"profile", false},
		{"flow", false},
	}

	for _, tt := range tests {
		result := IsChildType(tt.metadataType)
		if result != tt.expected {
			t.Errorf("IsChildType(%s) = %v; want %v", tt.metadataType, result, tt.expected)
		}
	}
}

func TestGetCanonicalDirectoryName(t *testing.T) {
	tests := []struct {
		metadataType string
		expected     string
	}{
		{"customfield", "objects"},
		{"validationrule", "objects"},
		{"CustomObject", "objects"},
		{"Profile", "profiles"},
		{"lightningcomponentbundle", "lwc"},
	}

	for _, tt := range tests {
		result := GetCanonicalDirectoryName(tt.metadataType)
		if result != tt.expected {
			t.Errorf("GetCanonicalDirectoryName(%s) = %s; want %s", tt.metadataType, result, tt.expected)
		}
	}
}
