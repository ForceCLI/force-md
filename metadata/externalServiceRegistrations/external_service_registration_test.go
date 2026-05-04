package externalServiceRegistrations

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenParsesAllKnownFields(t *testing.T) {
	tmpDir := t.TempDir()
	metaPath := filepath.Join(tmpDir, "OpenLibrary.externalServiceRegistration-meta.xml")
	metaContent := `<?xml version="1.0" encoding="UTF-8"?>
<ExternalServiceRegistration xmlns="http://soap.sforce.com/2006/04/metadata">
    <description>Open Library catalog API</description>
    <externalServiceProviderType>Http</externalServiceProviderType>
    <label>Open Library</label>
    <namedCredential>OpenLibraryNC</namedCredential>
    <registrationProviderType>Custom</registrationProviderType>
    <schema>{&quot;openapi&quot;:&quot;3.0.0&quot;,&quot;info&quot;:{&quot;title&quot;:&quot;OpenLibrary&quot;,&quot;version&quot;:&quot;1&quot;},&quot;paths&quot;:{}}</schema>
    <schemaType>OpenApi3</schemaType>
    <schemaUploadFileExtension>json</schemaUploadFileExtension>
    <schemaUploadFileName>openlibrary</schemaUploadFileName>
    <status>Complete</status>
    <systemVersion>3</systemVersion>
    <operations>
        <active>true</active>
        <name>getBooks</name>
    </operations>
</ExternalServiceRegistration>`

	require.NoError(t, os.WriteFile(metaPath, []byte(metaContent), 0o644))

	registration, err := Open(metaPath)
	require.NoError(t, err)
	require.NotNil(t, registration)

	assert.Equal(t, "OpenLibrary", string(registration.GetMetadataInfo().Name()))
	require.NotNil(t, registration.Label)
	assert.Equal(t, "Open Library", registration.Label.Text)
	require.NotNil(t, registration.NamedCredential)
	assert.Equal(t, "OpenLibraryNC", registration.NamedCredential.Text)
	require.NotNil(t, registration.SchemaType)
	assert.Equal(t, "OpenApi3", registration.SchemaType.Text)
	require.NotNil(t, registration.Schema)
	assert.True(t, strings.HasPrefix(registration.Schema.Text, `{"openapi":"3.0.0"`),
		"schema body should be unescaped JSON, got %q", registration.Schema.Text)
	require.NotNil(t, registration.SchemaUploadFileExtension)
	assert.Equal(t, "json", registration.SchemaUploadFileExtension.Text)
	require.NotNil(t, registration.Status)
	assert.Equal(t, "Complete", registration.Status.Text)
	require.Len(t, registration.Operations, 1)
	require.NotNil(t, registration.Operations[0].Name)
	assert.Equal(t, "getBooks", registration.Operations[0].Name.Text)
}

func TestOpenAcceptsMinimalRegistration(t *testing.T) {
	tmpDir := t.TempDir()
	metaPath := filepath.Join(tmpDir, "Minimal.externalServiceRegistration-meta.xml")
	metaContent := `<?xml version="1.0" encoding="UTF-8"?>
<ExternalServiceRegistration xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Minimal</label>
    <namedCredential>NC</namedCredential>
    <schemaType>OpenApi3</schemaType>
    <schema>{&quot;openapi&quot;:&quot;3.0.0&quot;}</schema>
</ExternalServiceRegistration>`
	require.NoError(t, os.WriteFile(metaPath, []byte(metaContent), 0o644))

	registration, err := Open(metaPath)
	require.NoError(t, err)
	require.NotNil(t, registration)
	assert.Nil(t, registration.Description)
	assert.Nil(t, registration.RegistrationUrl)
	assert.Nil(t, registration.Status)
	assert.Empty(t, registration.Operations)
}
