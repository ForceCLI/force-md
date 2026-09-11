package objects

import (
	"testing"

	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objects/field"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type itemsByType map[string][]metadata.RegisterableMetadata

func (p itemsByType) Items(metadataType string) []metadata.RegisterableMetadata {
	return p[metadataType]
}

func TestComposeFromChildrenNamesObjectWithoutBaseFile(t *testing.T) {
	f := &field.CustomField{Field: field.Field{FullName: "Custom__c"}}
	f.SetMetadata(metadata.NewMetadataInfo("Account.Custom__c", "objects/Account/fields/Custom__c.field-meta.xml"))
	provider := itemsByType{"CustomField": {f}}

	obj := ComposeFromChildren("Account", provider)

	assert.Equal(t, metadata.MetadataObjectName("Account"), obj.GetMetadataInfo().Name())
	require.Len(t, obj.Fields, 1)
	assert.Equal(t, "Custom__c", obj.Fields[0].FullName)

	files, err := obj.Files(metadata.MetadataFormat)
	require.NoError(t, err)
	_, ok := files["objects/Account.object"]
	assert.True(t, ok, "expected objects/Account.object, got %v", files)
}
