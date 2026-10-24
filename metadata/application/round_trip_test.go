package application

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// A console application keeps its console configuration, keyboard shortcuts,
// preferences and the Omni-Channel and Service Console flags.
func TestConsoleApplicationRoundTrip(t *testing.T) {
	contents := `<?xml version="1.0" encoding="UTF-8"?>
<CustomApplication xmlns="http://soap.sforce.com/2006/04/metadata">
    <consoleConfig>
        <detailPageRefreshMethod>none</detailPageRefreshMethod>
        <keyboardShortcuts>
            <defaultShortcuts>
                <action>FOCUS_CONSOLE</action>
                <active>true</active>
                <keyCommand>ESC</keyCommand>
            </defaultShortcuts>
            <defaultShortcuts>
                <action>REFRESH_TAB</action>
                <active>false</active>
                <keyCommand>SHIFT+R</keyCommand>
            </defaultShortcuts>
        </keyboardShortcuts>
        <listPlacement>
            <location>full</location>
        </listPlacement>
        <listRefreshMethod>none</listRefreshMethod>
    </consoleConfig>
    <defaultLandingTab>standard-home</defaultLandingTab>
    <isNavAutoTempTabsDisabled>false</isNavAutoTempTabsDisabled>
    <isNavPersonalizationDisabled>false</isNavPersonalizationDisabled>
    <isNavTabPersistenceDisabled>false</isNavTabPersistenceDisabled>
    <isOmniPinnedViewEnabled>false</isOmniPinnedViewEnabled>
    <isServiceCloudConsole>true</isServiceCloudConsole>
    <label>Service Console</label>
    <preferences>
        <enableCustomizeMyTabs>false</enableCustomizeMyTabs>
        <enableKeyboardShortcuts>true</enableKeyboardShortcuts>
        <enableListViewHover>true</enableListViewHover>
        <enableListViewReskin>true</enableListViewReskin>
        <enableMultiMonitorComponents>true</enableMultiMonitorComponents>
        <enablePinTabs>true</enablePinTabs>
        <enableTabHover>false</enableTabHover>
        <enableTabLimits>false</enableTabLimits>
        <saveUserSessions>true</saveUserSessions>
    </preferences>
    <tabs>standard-Case</tabs>
    <uiType>Lightning</uiType>
</CustomApplication>
`
	path := filepath.Join(t.TempDir(), "Service_Console.app")
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}
	a, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(a)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != contents {
		t.Errorf("round trip changed the application:\n%s", out)
	}
}
