package layout

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// A layout's feed layout, which sets the feed-based page's filters and
// sidebar components, is written back unchanged.
func TestFeedLayoutRoundTrip(t *testing.T) {
	content := `<?xml version="1.0" encoding="UTF-8"?>
<Layout xmlns="http://soap.sforce.com/2006/04/metadata">
    <feedLayout>
        <autocollapsePublisher>true</autocollapsePublisher>
        <compactFeed>true</compactFeed>
        <feedFilterPosition>CenterDropDown</feedFilterPosition>
        <feedFilters>
            <feedFilterType>AllUpdates</feedFilterType>
        </feedFilters>
        <feedFilters>
            <feedFilterType>FeedItemType</feedFilterType>
            <feedItemType>TextPost</feedItemType>
        </feedFilters>
        <fullWidthFeed>true</fullWidthFeed>
        <hideSidebar>false</hideSidebar>
        <highlightExternalFeedItems>true</highlightExternalFeedItems>
        <rightComponents>
            <componentType>Following</componentType>
        </rightComponents>
        <rightComponents>
            <componentType>Followers</componentType>
        </rightComponents>
        <useInlineFiltersInConsole>false</useInlineFiltersInConsole>
    </feedLayout>
    <layoutSections>
        <customLabel>false</customLabel>
        <detailHeading>false</detailHeading>
        <editHeading>true</editHeading>
        <label>Information</label>
        <layoutColumns>
            <layoutItems>
                <behavior>Required</behavior>
                <field>Name</field>
            </layoutItems>
        </layoutColumns>
        <style>OneColumn</style>
    </layoutSections>
    <showEmailCheckbox>false</showEmailCheckbox>
    <showHighlightsPanel>false</showHighlightsPanel>
    <showInteractionLogPanel>false</showInteractionLogPanel>
    <showRunAssignmentRulesCheckbox>false</showRunAssignmentRulesCheckbox>
    <showSubmitAndAttachButton>false</showSubmitAndAttachButton>
</Layout>
`
	path := filepath.Join(t.TempDir(), "Case-Case Layout.layout")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	l, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(l)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != content {
		t.Errorf("round trip changed the layout:\n%s", out)
	}
}
