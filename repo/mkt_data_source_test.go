package repo

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/metadata/dataSources"
	"github.com/ForceCLI/force-md/metadata/mktDataSources"
	"github.com/stretchr/testify/require"
)

// ExternalDataSource and DataSource share the .dataSource suffix; the root
// element decides which type a file is.
func TestMetadataFromPathDistinguishesDataSourceTypes(t *testing.T) {
	root := t.TempDir()
	externalDir := filepath.Join(root, "dataSources")
	mktDir := filepath.Join(root, "mktDataSources")
	require.NoError(t, os.MkdirAll(externalDir, 0o755))
	require.NoError(t, os.MkdirAll(mktDir, 0o755))

	external := filepath.Join(externalDir, "Sample.dataSource-meta.xml")
	require.NoError(t, os.WriteFile(external, []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ExternalDataSource xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Sample</label>
    <type>SimpleURL</type>
</ExternalDataSource>`), 0o644))

	mkt := filepath.Join(mktDir, "Sample.dataSource-meta.xml")
	require.NoError(t, os.WriteFile(mkt, []byte(`<?xml version="1.0" encoding="UTF-8"?>
<DataSource xmlns="http://soap.sforce.com/2006/04/metadata">
    <masterLabel>Sample</masterLabel>
    <prefix>SMP</prefix>
</DataSource>`), 0o644))

	m, err := MetadataFromPath(external)
	require.NoError(t, err)
	require.Equal(t, datasource.NAME, m.Type())

	m, err = MetadataFromPath(mkt)
	require.NoError(t, err)
	require.Equal(t, mktDataSources.NAME, m.Type())
}
