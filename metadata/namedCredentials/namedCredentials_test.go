package namedCredentials

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNamedCredentialOpenSupportsLegacyRootFields(t *testing.T) {
	tmpDir := t.TempDir()
	metaPath := filepath.Join(tmpDir, "LegacyDadJokes.namedCredential-meta.xml")
	metaContent := `<?xml version="1.0" encoding="UTF-8"?>
<NamedCredential xmlns="http://soap.sforce.com/2006/04/metadata">
    <allowMergeFieldsInBody>false</allowMergeFieldsInBody>
    <allowMergeFieldsInHeader>false</allowMergeFieldsInHeader>
    <calloutStatus>Enabled</calloutStatus>
    <endpoint>https://icanhazdadjoke.com/</endpoint>
    <generateAuthorizationHeader>false</generateAuthorizationHeader>
    <label>Legacy Dad Jokes</label>
    <principalType>Anonymous</principalType>
    <protocol>NoAuthentication</protocol>
</NamedCredential>`

	require.NoError(t, os.WriteFile(metaPath, []byte(metaContent), 0o644))

	namedCredential, err := Open(metaPath)
	require.NoError(t, err)
	require.NotNil(t, namedCredential)

	assert.Equal(t, "LegacyDadJokes", string(namedCredential.GetMetadataInfo().Name()))
	assert.Equal(t, "Legacy Dad Jokes", namedCredential.Label.Text)
	require.NotNil(t, namedCredential.Endpoint)
	require.NotNil(t, namedCredential.PrincipalType)
	require.NotNil(t, namedCredential.Protocol)
	assert.Equal(t, "https://icanhazdadjoke.com/", namedCredential.Endpoint.Text)
	assert.Equal(t, "Anonymous", namedCredential.PrincipalType.Text)
	assert.Equal(t, "NoAuthentication", namedCredential.Protocol.Text)
	assert.Equal(t, "Enabled", namedCredential.CalloutStatus.Text)
	assert.Equal(t, "false", namedCredential.GenerateAuthorizationHeader.Text)
	assert.Empty(t, namedCredential.NamedCredentialParameters)
	assert.Nil(t, namedCredential.AuthProvider)
}

func TestNamedCredentialMarshalOmitsUnsetLegacyRootFields(t *testing.T) {
	namedCredential := &NamedCredential{}
	namedCredential.Xmlns = "http://soap.sforce.com/2006/04/metadata"
	namedCredential.Label.Text = "Modern Credential"
	namedCredential.NamedCredentialType.Text = "SecuredEndpoint"
	namedCredential.GenerateAuthorizationHeader.Text = "true"
	namedCredential.NamedCredentialParameters = []struct {
		ExternalCredential *struct {
			Text string `xml:",chardata"`
		} `xml:"externalCredential"`
		ParameterName struct {
			Text string `xml:",chardata"`
		} `xml:"parameterName"`
		ParameterType struct {
			Text string `xml:",chardata"`
		} `xml:"parameterType"`
		ParameterValue *struct {
			Text string `xml:",chardata"`
		} `xml:"parameterValue"`
	}{
		{
			ParameterName: struct {
				Text string `xml:",chardata"`
			}{Text: "Url"},
			ParameterType: struct {
				Text string `xml:",chardata"`
			}{Text: "Url"},
			ParameterValue: &struct {
				Text string `xml:",chardata"`
			}{Text: "https://api.example.com"},
		},
	}

	marshaled, err := internal.Marshal(namedCredential)
	require.NoError(t, err)

	assert.NotContains(t, string(marshaled), "<endpoint>")
	assert.NotContains(t, string(marshaled), "<principalType>")
	assert.NotContains(t, string(marshaled), "<protocol>")
	assert.NotContains(t, string(marshaled), "<authProvider>")
}
