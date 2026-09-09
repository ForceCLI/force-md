package mktDataSources

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenDataSource(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "mktDataSources")
	require.NoError(t, os.MkdirAll(dir, 0o755))
	path := filepath.Join(dir, "SFMC100008473.dataSource-meta.xml")
	content := `<?xml version="1.0" encoding="UTF-8"?>
<DataSource xmlns="http://soap.sforce.com/2006/04/metadata">
    <masterLabel>SFMC100008473</masterLabel>
    <prefix>SFM</prefix>
</DataSource>`
	require.NoError(t, os.WriteFile(path, []byte(content), 0o644))

	ds, err := Open(path)
	require.NoError(t, err)
	assert.Equal(t, NAME, ds.Type())
	assert.Equal(t, "SFMC100008473", ds.MasterLabel.Text)
	assert.Equal(t, "SFM", ds.Prefix.Text)
}
