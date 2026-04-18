package uibundles

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUIBundleOpen(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "demoBundle.uibundle-meta.xml")
	require.NoError(t, os.WriteFile(path, []byte(`<?xml version="1.0" encoding="UTF-8"?>
<UIBundle xmlns="http://soap.sforce.com/2006/04/metadata">
    <masterLabel>Demo Bundle</masterLabel>
    <description>Preview bundle.</description>
    <isActive>true</isActive>
    <version>1</version>
</UIBundle>`), 0o644))

	bundle, err := Open(path)
	require.NoError(t, err)
	require.NotNil(t, bundle.MasterLabel)
	assert.Equal(t, "demoBundle", string(bundle.Name()))
	assert.Equal(t, "Demo Bundle", bundle.MasterLabel.Text)
	assert.Equal(t, path, string(bundle.Path()))
}

func TestUIBundleFiles(t *testing.T) {
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
	require.NoError(t, os.WriteFile(filepath.Join(bundleDir, "index.html"), []byte("<!doctype html><div id=app></div>"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(bundleDir, "ui-bundle.json"), []byte(`{"outputDir":"dist"}`), 0o644))

	bundle, err := Open(metadataPath)
	require.NoError(t, err)

	files, err := bundle.Files(metadata.SourceFormat)
	require.NoError(t, err)
	assert.Contains(t, files, filepath.Join("uiBundles", "demoBundle", "demoBundle.uibundle-meta.xml"))
	assert.Contains(t, files, filepath.Join("uiBundles", "demoBundle", "index.html"))
	assert.Contains(t, files, filepath.Join("uiBundles", "demoBundle", "ui-bundle.json"))
}
