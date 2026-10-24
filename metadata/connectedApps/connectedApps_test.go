package connectedapps

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

func assertRoundTrip(t *testing.T, content string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "Test.connectedApp")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	c, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(c)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != content {
		t.Errorf("round trip changed the connected app:\n%s", out)
	}
}

// The OAuth flow settings, a contact phone, and the profiles granted access
// are written back, and a connected app without a certificate is written
// without an empty certificate element.
func TestConnectedAppOauthRoundTrip(t *testing.T) {
	assertRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
    <contactEmail>admin@example.com</contactEmail>
    <contactPhone>555-0100</contactPhone>
    <label>Integration</label>
    <oauthConfig>
        <callbackUrl>https://example.com/callback</callbackUrl>
        <consumerKey>key</consumerKey>
        <isAdminApproved>true</isAdminApproved>
        <isClientCredentialEnabled>false</isClientCredentialEnabled>
        <isCodeCredentialEnabled>false</isCodeCredentialEnabled>
        <isCodeCredentialPostOnly>false</isCodeCredentialPostOnly>
        <isConsumerSecretOptional>false</isConsumerSecretOptional>
        <isIntrospectAllTokens>false</isIntrospectAllTokens>
        <isNamedUserJwtEnabled>false</isNamedUserJwtEnabled>
        <isPkceRequired>true</isPkceRequired>
        <isRefreshTokenRotationEnabled>false</isRefreshTokenRotationEnabled>
        <isSecretRequiredForRefreshToken>true</isSecretRequiredForRefreshToken>
        <isSecretRequiredForTokenExchange>false</isSecretRequiredForTokenExchange>
        <isTokenExchangeEnabled>false</isTokenExchangeEnabled>
        <scopes>Api</scopes>
        <scopes>RefreshToken</scopes>
    </oauthConfig>
    <oauthPolicy>
        <ipRelaxation>ENFORCE</ipRelaxation>
        <isTokenExchangeFlowEnabled>false</isTokenExchangeFlowEnabled>
        <refreshTokenPolicy>infinite</refreshTokenPolicy>
    </oauthPolicy>
    <profileName>System Administrator</profileName>
    <profileName>Integration User</profileName>
</ConnectedApp>
`)
}

// A connected app without OAuth settings is written without empty
// oauthConfig and oauthPolicy elements.
func TestConnectedAppWithoutOauthRoundTrip(t *testing.T) {
	assertRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
    <contactEmail>admin@example.com</contactEmail>
    <label>Canvas</label>
</ConnectedApp>
`)
}
