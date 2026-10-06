package standardValueSetTranslation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenReadsValueTranslations(t *testing.T) {
	path := filepath.Join(t.TempDir(), "LeadSource-fr.standardValueSetTranslation-meta.xml")
	require.NoError(t, os.WriteFile(path, []byte(`<?xml version="1.0" encoding="UTF-8"?>
<StandardValueSetTranslation xmlns="http://soap.sforce.com/2006/04/metadata">
    <valueTranslation>
        <masterLabel>Web</masterLabel>
        <translation>Internet</translation>
    </valueTranslation>
    <valueTranslation>
        <masterLabel>Other</masterLabel>
        <translation><!-- Other --></translation>
    </valueTranslation>
</StandardValueSetTranslation>
`), 0o644))

	trans, err := Open(path)
	require.NoError(t, err)
	require.Len(t, trans.ValueTranslation, 2)
	assert.Equal(t, "Web", trans.ValueTranslation[0].MasterLabel.Text)
	assert.Equal(t, "Internet", trans.ValueTranslation[0].Translation.Text)
	assert.Equal(t, "", trans.ValueTranslation[1].Translation.Text)
	assert.Equal(t, NAME, string(trans.Type()))
}
