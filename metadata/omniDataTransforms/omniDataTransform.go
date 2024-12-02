package omniDataTransform

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "OmniDataTransform"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type OmniDataTransform struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"OmniDataTransform"`
	Xmlns   string   `xml:"xmlns,attr"`
	Active  struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
	AssignmentRulesUsed struct {
		Text string `xml:",chardata"`
	} `xml:"assignmentRulesUsed"`
	DeletedOnSuccess struct {
		Text string `xml:",chardata"`
	} `xml:"deletedOnSuccess"`
	ErrorIgnored struct {
		Text string `xml:",chardata"`
	} `xml:"errorIgnored"`
	FieldLevelSecurityEnabled struct {
		Text string `xml:",chardata"`
	} `xml:"fieldLevelSecurityEnabled"`
	InputType struct {
		Text string `xml:",chardata"`
	} `xml:"inputType"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	NullInputsIncludedInOutput struct {
		Text string `xml:",chardata"`
	} `xml:"nullInputsIncludedInOutput"`
	OmniDataTransformItem []struct {
		Disabled struct {
			Text string `xml:",chardata"`
		} `xml:"disabled"`
		FilterGroup struct {
			Text string `xml:",chardata"`
		} `xml:"filterGroup"`
		GlobalKey struct {
			Text string `xml:",chardata"`
		} `xml:"globalKey"`
		InputFieldName struct {
			Text string `xml:",chardata"`
		} `xml:"inputFieldName"`
		InputObjectQuerySequence struct {
			Text string `xml:",chardata"`
		} `xml:"inputObjectQuerySequence"`
		LinkedObjectSequence struct {
			Text string `xml:",chardata"`
		} `xml:"linkedObjectSequence"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		OutputCreationSequence struct {
			Text string `xml:",chardata"`
		} `xml:"outputCreationSequence"`
		OutputFieldName struct {
			Text string `xml:",chardata"`
		} `xml:"outputFieldName"`
		OutputObjectName struct {
			Text string `xml:",chardata"`
		} `xml:"outputObjectName"`
		RequiredForUpsert struct {
			Text string `xml:",chardata"`
		} `xml:"requiredForUpsert"`
		UpsertKey struct {
			Text string `xml:",chardata"`
		} `xml:"upsertKey"`
		FormulaConverted struct {
			Text string `xml:",chardata"`
		} `xml:"formulaConverted"`
		FormulaExpression struct {
			Text string `xml:",chardata"`
		} `xml:"formulaExpression"`
		FormulaResultPath struct {
			Text string `xml:",chardata"`
		} `xml:"formulaResultPath"`
		FormulaSequence struct {
			Text string `xml:",chardata"`
		} `xml:"formulaSequence"`
		FilterOperator struct {
			Text string `xml:",chardata"`
		} `xml:"filterOperator"`
		MigrationValue struct {
			Text string `xml:",chardata"`
		} `xml:"migrationValue"`
		DefaultValue struct {
			Text string `xml:",chardata"`
		} `xml:"defaultValue"`
		FilterValue struct {
			Text string `xml:",chardata"`
		} `xml:"filterValue"`
		InputObjectName struct {
			Text string `xml:",chardata"`
		} `xml:"inputObjectName"`
		OutputFieldFormat struct {
			Text string `xml:",chardata"`
		} `xml:"outputFieldFormat"`
		TransformValuesMappings struct {
			Text string `xml:",chardata"`
		} `xml:"transformValuesMappings"`
		FilterDataType struct {
			Text string `xml:",chardata"`
		} `xml:"filterDataType"`
	} `xml:"omniDataTransformItem"`
	OutputType struct {
		Text string `xml:",chardata"`
	} `xml:"outputType"`
	PreviewJsonData struct {
		Text string `xml:",chardata"`
	} `xml:"previewJsonData"`
	ProcessSuperBulk struct {
		Text string `xml:",chardata"`
	} `xml:"processSuperBulk"`
	ResponseCacheTtlMinutes struct {
		Text string `xml:",chardata"`
	} `xml:"responseCacheTtlMinutes"`
	RollbackOnError struct {
		Text string `xml:",chardata"`
	} `xml:"rollbackOnError"`
	SourceObject struct {
		Text string `xml:",chardata"`
	} `xml:"sourceObject"`
	SourceObjectDefault struct {
		Text string `xml:",chardata"`
	} `xml:"sourceObjectDefault"`
	SynchronousProcessThreshold struct {
		Text string `xml:",chardata"`
	} `xml:"synchronousProcessThreshold"`
	OmniDataTransformType struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	UniqueName struct {
		Text string `xml:",chardata"`
	} `xml:"uniqueName"`
	VersionNumber struct {
		Text string `xml:",chardata"`
	} `xml:"versionNumber"`
	XmlDeclarationRemoved struct {
		Text string `xml:",chardata"`
	} `xml:"xmlDeclarationRemoved"`
	ExpectedOutputJson struct {
		Text string `xml:",chardata"`
	} `xml:"expectedOutputJson"`
	ResponseCacheType struct {
		Text string `xml:",chardata"`
	} `xml:"responseCacheType"`
	ExpectedInputJson struct {
		Text string `xml:",chardata"`
	} `xml:"expectedInputJson"`
	TargetOutputFileName struct {
		Text string `xml:",chardata"`
	} `xml:"targetOutputFileName"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *OmniDataTransform) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *OmniDataTransform) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*OmniDataTransform, error) {
	p := &OmniDataTransform{}
	return p, metadata.ParseMetadataXml(p, path)
}
