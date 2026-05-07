package slackapps

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const sampleMetadataXML = `<?xml version="1.0" encoding="UTF-8"?>
<SlackApp xmlns="http://soap.sforce.com/2006/04/metadata">
    <appKey>*</appKey>
    <appToken>*</appToken>
    <botScopes>chat:write,chat:write.public</botScopes>
    <clientKey>*</clientKey>
    <clientSecret>*</clientSecret>
    <isProtected>false</isProtected>
    <masterLabel>Apex Example</masterLabel>
    <signingSecret>*</signingSecret>
    <userScopes>im:read,im:write</userScopes>
</SlackApp>
`

func TestOpen_ParsesSlackAppMetadataXml(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "ApexExample.slackapp-meta.xml")
	if err := os.WriteFile(path, []byte(sampleMetadataXML), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	app, err := Open(path)
	if err != nil {
		t.Fatalf("Open returned error: %v", err)
	}

	if app.MasterLabel.Text != "Apex Example" {
		t.Errorf("MasterLabel = %q, want %q", app.MasterLabel.Text, "Apex Example")
	}
	if app.AppKey.Text != "*" {
		t.Errorf("AppKey = %q, want %q", app.AppKey.Text, "*")
	}
	if app.BotScopes == nil || !strings.Contains(app.BotScopes.Text, "chat:write") {
		t.Errorf("BotScopes = %+v, want non-nil containing chat:write", app.BotScopes)
	}
	if app.UserScopes == nil || !strings.Contains(app.UserScopes.Text, "im:read") {
		t.Errorf("UserScopes = %+v, want non-nil containing im:read", app.UserScopes)
	}
	if app.IsProtected == nil || app.IsProtected.Text != "false" {
		t.Errorf("IsProtected = %+v, want non-nil 'false'", app.IsProtected)
	}
	if got, want := app.Type(), "SlackApp"; string(got) != want {
		t.Errorf("Type() = %q, want %q", got, want)
	}
}
