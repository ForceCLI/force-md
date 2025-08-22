package profile

import (
	"testing"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/metadata/permissionset"
	"github.com/stretchr/testify/assert"
)

func TestCloneApexClassAccess(t *testing.T) {
	tests := []struct {
		name        string
		profile     *Profile
		sourceClass string
		destClass   string
		expectError bool
		errorMsg    string
	}{
		{
			name: "successful clone",
			profile: &Profile{
				ClassAccesses: permissionset.ApexClassList{
					{
						ApexClass: "SourceClass",
						Enabled:   BooleanText{Text: "true"},
					},
					{
						ApexClass: "OtherClass",
						Enabled:   BooleanText{Text: "false"},
					},
				},
			},
			sourceClass: "SourceClass",
			destClass:   "DestClass",
			expectError: false,
		},
		{
			name: "clone with disabled source",
			profile: &Profile{
				ClassAccesses: permissionset.ApexClassList{
					{
						ApexClass: "SourceClass",
						Enabled:   BooleanText{Text: "false"},
					},
				},
			},
			sourceClass: "SourceClass",
			destClass:   "DestClass",
			expectError: false,
		},
		{
			name: "source class not found",
			profile: &Profile{
				ClassAccesses: permissionset.ApexClassList{
					{
						ApexClass: "OtherClass",
						Enabled:   BooleanText{Text: "true"},
					},
				},
			},
			sourceClass: "SourceClass",
			destClass:   "DestClass",
			expectError: true,
			errorMsg:    "source apex class not found",
		},
		{
			name: "destination class already exists",
			profile: &Profile{
				ClassAccesses: permissionset.ApexClassList{
					{
						ApexClass: "SourceClass",
						Enabled:   BooleanText{Text: "true"},
					},
					{
						ApexClass: "DestClass",
						Enabled:   BooleanText{Text: "false"},
					},
				},
			},
			sourceClass: "SourceClass",
			destClass:   "DestClass",
			expectError: true,
			errorMsg:    "apex class already exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.profile.CloneApexClassAccess(tt.sourceClass, tt.destClass)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				assert.NoError(t, err)
				// Verify the class was added
				found := false
				var clonedClass permissionset.ApexClass
				for _, c := range tt.profile.ClassAccesses {
					if c.ApexClass == tt.destClass {
						found = true
						clonedClass = c
						break
					}
				}
				assert.True(t, found, "Cloned class should exist")

				// Verify it has the same enabled status as source
				var sourceClass permissionset.ApexClass
				for _, c := range tt.profile.ClassAccesses {
					if c.ApexClass == tt.sourceClass {
						sourceClass = c
						break
					}
				}
				assert.Equal(t, sourceClass.Enabled.Text, clonedClass.Enabled.Text)
			}
		})
	}
}
