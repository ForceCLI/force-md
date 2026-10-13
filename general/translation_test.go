package general

import (
	"encoding/xml"
	"testing"

	nbioxml "github.com/nbio/xml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type translationDoc struct {
	XMLName xml.Name         `xml:"fields"`
	Label   *TranslationText `xml:"label"`
	Help    *TranslationText `xml:"help"`
	Name    TranslationText  `xml:"name"`
}

func TestTranslationTextReadsDecodedTextAndComment(t *testing.T) {
	var doc translationDoc
	require.NoError(t, nbioxml.Unmarshal([]byte(`<fields><label>R&amp;D</label><help><!-- Help Text --></help><name>Dept__c</name></fields>`), &doc))
	assert.Equal(t, "R&D", doc.Label.Text)
	assert.Equal(t, "", doc.Label.Comment)
	assert.Equal(t, "", doc.Help.Text)
	assert.Equal(t, " Help Text ", doc.Help.Comment)
	assert.Equal(t, "Dept__c", doc.Name.Text)
}

func TestTranslationTextWritesTextAndComment(t *testing.T) {
	input := `<fields>
    <label>R&amp;D</label>
    <help><!-- Help Text --></help>
    <name>Dept__c</name>
</fields>`
	var doc translationDoc
	require.NoError(t, nbioxml.Unmarshal([]byte(input), &doc))
	out, err := xml.MarshalIndent(doc, "", "    ")
	require.NoError(t, err)
	assert.Equal(t, input, string(out))
}
