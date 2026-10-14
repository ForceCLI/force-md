package navigationMenus

import (
	"os"
	"path/filepath"
	"testing"
)

const sampleNavigationMenu = `<?xml version="1.0" encoding="UTF-8"?>
<NavigationMenu xmlns="http://soap.sforce.com/2006/04/metadata">
    <container>Customer Portal</container>
    <containerType>Network</containerType>
    <label>Default Navigation</label>
    <navigationMenuItem>
        <defaultListViewId>AllCases</defaultListViewId>
        <label>My Cases</label>
        <menuItemBranding>
            <tileImage>case_120</tileImage>
        </menuItemBranding>
        <position>1</position>
        <publiclyAvailable>false</publiclyAvailable>
        <target>Case</target>
        <type>SalesforceObject</type>
    </navigationMenuItem>
    <navigationMenuItem>
        <label>Locations</label>
        <position>2</position>
        <publiclyAvailable>true</publiclyAvailable>
        <subMenu>
            <navigationMenuItem>
                <label>Main Office</label>
                <position>0</position>
                <publiclyAvailable>true</publiclyAvailable>
                <target>https://example.com/main</target>
                <targetPreference>OpenExternalLinkInSameTab</targetPreference>
                <type>ExternalLink</type>
            </navigationMenuItem>
        </subMenu>
        <type>MenuLabel</type>
    </navigationMenuItem>
</NavigationMenu>
`

func TestOpen(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "SFDC_Default_Navigation_Customer_Portal.navigationMenu-meta.xml")
	if err := os.WriteFile(path, []byte(sampleNavigationMenu), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	n, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if got := n.Container.String(); got != "Customer Portal" {
		t.Errorf("Container = %q, want %q", got, "Customer Portal")
	}
	if got := n.ContainerType.String(); got != "Network" {
		t.Errorf("ContainerType = %q, want %q", got, "Network")
	}
	if got := n.Label.String(); got != "Default Navigation" {
		t.Errorf("Label = %q, want %q", got, "Default Navigation")
	}
	if got := string(n.Type()); got != "NavigationMenu" {
		t.Errorf("Type = %q, want NavigationMenu", got)
	}
	if got := len(n.NavigationMenuItems); got != 2 {
		t.Fatalf("len(NavigationMenuItems) = %d, want 2", got)
	}

	cases := n.NavigationMenuItems[0]
	if got := cases.DefaultListViewId.String(); got != "AllCases" {
		t.Errorf("item[0].DefaultListViewId = %q, want %q", got, "AllCases")
	}
	if cases.MenuItemBranding == nil || cases.MenuItemBranding.TileImage.String() != "case_120" {
		t.Errorf("item[0].MenuItemBranding.TileImage = %v, want case_120", cases.MenuItemBranding)
	}
	if got := cases.PubliclyAvailable.String(); got != "false" {
		t.Errorf("item[0].PubliclyAvailable = %q, want false", got)
	}
	if cases.SubMenu != nil {
		t.Errorf("item[0].SubMenu = %v, want nil", cases.SubMenu)
	}

	locations := n.NavigationMenuItems[1]
	if got := locations.Type.String(); got != "MenuLabel" {
		t.Errorf("item[1].Type = %q, want MenuLabel", got)
	}
	if locations.Target != nil {
		t.Errorf("item[1].Target = %v, want nil", locations.Target)
	}
	if locations.SubMenu == nil || len(locations.SubMenu.NavigationMenuItems) != 1 {
		t.Fatalf("item[1].SubMenu = %v, want one item", locations.SubMenu)
	}
	office := locations.SubMenu.NavigationMenuItems[0]
	if got := office.Label.String(); got != "Main Office" {
		t.Errorf("submenu item Label = %q, want %q", got, "Main Office")
	}
	if got := office.Position.String(); got != "0" {
		t.Errorf("submenu item Position = %q, want 0", got)
	}
	if got := office.TargetPreference.String(); got != "OpenExternalLinkInSameTab" {
		t.Errorf("submenu item TargetPreference = %q, want OpenExternalLinkInSameTab", got)
	}
}
