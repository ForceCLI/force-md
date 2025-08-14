package validationrule

import (
	"encoding/xml"
	"path/filepath"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const VALIDATION_RULE_TRANSLATION_NAME = "ValidationRuleTranslation"

func init() {
	internal.TypeRegistry.Register(VALIDATION_RULE_TRANSLATION_NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ValidationRule struct {
	ErrorMessage struct {
		Text string `xml:",chardata"`
	} `xml:"errorMessage"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
}

type ValidationRuleTranslation struct {
	metadata.MetadataInfo
	XMLName      xml.Name `xml:"ValidationRuleTranslation"`
	Xmlns        string   `xml:"xmlns,attr"`
	ErrorMessage struct {
		Text string `xml:",chardata"`
	} `xml:"errorMessage"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
}

func (v *ValidationRuleTranslation) SetMetadata(m metadata.MetadataInfo) {
	v.MetadataInfo = m
}

func (v *ValidationRuleTranslation) Type() metadata.MetadataType {
	return VALIDATION_RULE_TRANSLATION_NAME
}

func (v *ValidationRuleTranslation) NameFromPath(path string) metadata.MetadataObjectName {
	// Extract parent-child name from path
	// e.g., /objectTranslations/Account-en_US/validationRules/Amount_Must_Be_Positive.validationRuleTranslation-meta.xml
	// or /objectTranslations/Account-en_US/Amount_Must_Be_Positive.validationRuleTranslation-meta.xml
	// should return "Account-en_US.Amount_Must_Be_Positive"
	dir := filepath.Dir(path)
	parentDir := filepath.Dir(dir)

	// Check if the file is in a validationRules subdirectory
	if filepath.Base(dir) == "validationRules" {
		parentName := filepath.Base(parentDir)
		baseName := filepath.Base(path)
		childName := strings.TrimSuffix(baseName, ".validationRuleTranslation-meta.xml")
		childName = strings.TrimSuffix(childName, ".validationRuleTranslation")
		return metadata.MetadataObjectName(parentName + "." + childName)
	} else {
		// File is directly in the object translation directory
		parentName := filepath.Base(dir)
		baseName := filepath.Base(path)
		childName := strings.TrimSuffix(baseName, ".validationRuleTranslation-meta.xml")
		childName = strings.TrimSuffix(childName, ".validationRuleTranslation")
		return metadata.MetadataObjectName(parentName + "." + childName)
	}
}

func Open(path string) (*ValidationRuleTranslation, error) {
	vr := &ValidationRuleTranslation{}
	return vr, metadata.ParseMetadataXml(vr, path)
}
