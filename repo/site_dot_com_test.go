package repo

import (
	"os"
	"path/filepath"
	"testing"

	siteDotComSite "github.com/ForceCLI/force-md/metadata/siteDotComSites"
	site "github.com/ForceCLI/force-md/metadata/sites"
	"github.com/stretchr/testify/require"
)

// CustomSite and SiteDotCom share the .site and .site-meta.xml suffixes; the
// root element decides which type a file is. A SiteDotCom's .site file holds the
// site's content, not XML, and is described by its .site-meta.xml.
func TestMetadataFromPathDistinguishesSiteTypes(t *testing.T) {
	root := t.TempDir()
	sitesDir := filepath.Join(root, "sites")
	siteDotComDir := filepath.Join(root, "siteDotComSites")
	require.NoError(t, os.MkdirAll(sitesDir, 0o755))
	require.NoError(t, os.MkdirAll(siteDotComDir, 0o755))

	customSite := filepath.Join(sitesDir, "Portal.site-meta.xml")
	require.NoError(t, os.WriteFile(customSite, []byte(`<?xml version="1.0" encoding="UTF-8"?>
<CustomSite xmlns="http://soap.sforce.com/2006/04/metadata">
    <active>true</active>
    <masterLabel>Portal</masterLabel>
    <siteType>ChatterNetwork</siteType>
</CustomSite>`), 0o644))

	siteDotComMeta := filepath.Join(siteDotComDir, "Portal1.site-meta.xml")
	require.NoError(t, os.WriteFile(siteDotComMeta, []byte(`<?xml version="1.0" encoding="UTF-8"?>
<SiteDotCom xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Portal</label>
    <siteType>ChatterNetworkPicasso</siteType>
</SiteDotCom>`), 0o644))
	siteDotComContent := filepath.Join(siteDotComDir, "Portal1.site")
	require.NoError(t, os.WriteFile(siteDotComContent, []byte{0x50, 0x4b, 0x03, 0x04}, 0o644))

	m, err := MetadataFromPath(customSite)
	require.NoError(t, err)
	require.Equal(t, site.NAME, m.Type())

	m, err = MetadataFromPath(siteDotComMeta)
	require.NoError(t, err)
	require.Equal(t, siteDotComSite.NAME, m.Type())
	require.Equal(t, "ChatterNetworkPicasso", m.(*siteDotComSite.SiteDotCom).SiteType.Text)

	m, err = MetadataFromPath(siteDotComContent)
	require.NoError(t, err)
	require.Equal(t, siteDotComSite.NAME, m.Type())
}
