package recordtype

import (
	"encoding/xml"
	"path/filepath"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const RECORD_TYPE_TRANSLATION_NAME = "RecordTypeTranslation"

func init() {
	internal.TypeRegistry.Register(RECORD_TYPE_TRANSLATION_NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type RecordType struct {
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
}

type RecordTypeTranslation struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"RecordTypeTranslation"`
	Xmlns       string   `xml:"xmlns,attr"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
}

func (r *RecordTypeTranslation) SetMetadata(m metadata.MetadataInfo) {
	r.MetadataInfo = m
}

func (r *RecordTypeTranslation) Type() metadata.MetadataType {
	return RECORD_TYPE_TRANSLATION_NAME
}

func (r *RecordTypeTranslation) NameFromPath(path string) metadata.MetadataObjectName {
	// Extract parent-child name from path
	// e.g., /objectTranslations/Account-en_US/recordTypes/Business.recordTypeTranslation-meta.xml
	// or /objectTranslations/Account-en_US/Business.recordTypeTranslation-meta.xml
	// should return "Account-en_US.Business"
	dir := filepath.Dir(path)
	parentDir := filepath.Dir(dir)

	// Check if the file is in a recordTypes subdirectory
	if filepath.Base(dir) == "recordTypes" {
		parentName := filepath.Base(parentDir)
		baseName := filepath.Base(path)
		childName := strings.TrimSuffix(baseName, ".recordTypeTranslation-meta.xml")
		childName = strings.TrimSuffix(childName, ".recordTypeTranslation")
		return metadata.MetadataObjectName(parentName + "." + childName)
	} else {
		// File is directly in the object translation directory
		parentName := filepath.Base(dir)
		baseName := filepath.Base(path)
		childName := strings.TrimSuffix(baseName, ".recordTypeTranslation-meta.xml")
		childName = strings.TrimSuffix(childName, ".recordTypeTranslation")
		return metadata.MetadataObjectName(parentName + "." + childName)
	}
}

func Open(path string) (*RecordTypeTranslation, error) {
	rt := &RecordTypeTranslation{}
	return rt, metadata.ParseMetadataXml(rt, path)
}
