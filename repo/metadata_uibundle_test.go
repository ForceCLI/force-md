package repo

import (
	"os"
	"path/filepath"
	"testing"

	uibundles "github.com/ForceCLI/force-md/metadata/uibundles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetadataFromPathFindsUIBundleMetadataFromBundleFile(t *testing.T) {
	root := t.TempDir()
	bundleDir := filepath.Join(root, "uiBundles", "demoBundle")
	require.NoError(t, os.MkdirAll(bundleDir, 0o755))
	metadataPath := filepath.Join(bundleDir, "demoBundle.uibundle-meta.xml")
	require.NoError(t, os.WriteFile(metadataPath, []byte(`<?xml version="1.0" encoding="UTF-8"?>
<UIBundle xmlns="http://soap.sforce.com/2006/04/metadata">
    <masterLabel>Demo Bundle</masterLabel>
    <description>Preview bundle.</description>
    <isActive>true</isActive>
    <version>1</version>
</UIBundle>`), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(bundleDir, "index.html"), []byte("<!doctype html>"), 0o644))

	found, err := MetadataFromPath(filepath.Join(bundleDir, "index.html"))
	require.NoError(t, err)
	bundle, ok := found.(*uibundles.UIBundle)
	require.True(t, ok)
	assert.Equal(t, "demoBundle", string(bundle.Name()))
}
