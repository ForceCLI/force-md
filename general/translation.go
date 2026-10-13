package general

import (
	"encoding/xml"

	nbioxml "github.com/nbio/xml"
)

// TranslationText is a translated value in a translation file. Text holds
// the value with XML entities decoded. Salesforce writes a value that has not
// been translated as an XML comment holding the untranslated text, such as
// <label><!-- Account Name --></label>. Text is empty for such a value, and
// Comment holds the comment's content so the value can be written back
// unchanged.
//
// UnmarshalXML takes github.com/nbio/xml types because metadata files are
// parsed with that package, while MarshalXML takes encoding/xml types
// because they are written with encoding/xml.
type TranslationText struct {
	Text    string
	Comment string
}

func (t *TranslationText) UnmarshalXML(d *nbioxml.Decoder, start nbioxml.StartElement) error {
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch tok := tok.(type) {
		case nbioxml.CharData:
			t.Text += string(tok)
		case nbioxml.Comment:
			t.Comment += string(tok)
		case nbioxml.StartElement:
			if err := d.Skip(); err != nil {
				return err
			}
		case nbioxml.EndElement:
			return nil
		}
	}
}

func (t TranslationText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if t.Text != "" || t.Comment == "" {
		if err := e.EncodeToken(xml.CharData(t.Text)); err != nil {
			return err
		}
	} else if err := e.EncodeToken(xml.Comment(t.Comment)); err != nil {
		return err
	}
	return e.EncodeToken(start.End())
}
