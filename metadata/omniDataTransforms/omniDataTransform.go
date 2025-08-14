package omniDataTransform

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
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
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	ErrorIgnored struct {
		Text string `xml:",chardata"`
	} `xml:"errorIgnored"`
	ExpectedInputJson         *TextLiteral `xml:"expectedInputJson"`
	ExpectedOutputJson        *TextLiteral `xml:"expectedOutputJson"`
	FieldLevelSecurityEnabled struct {
		Text string `xml:",chardata"`
	} `xml:"fieldLevelSecurityEnabled"`
	InputType struct {
		Text string `xml:",chardata"`
	} `xml:"inputType"`
	IsManagedUsingStdDesigner *struct {
		Text string `xml:",chardata"`
	} `xml:"isManagedUsingStdDesigner"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	NullInputsIncludedInOutput struct {
		Text string `xml:",chardata"`
	} `xml:"nullInputsIncludedInOutput"`
	OmniDataTransformItem []struct {
		DefaultValue *struct {
			Text string `xml:",chardata"`
		} `xml:"defaultValue"`
		Disabled struct {
			Text string `xml:",chardata"`
		} `xml:"disabled"`
		FilterDataType *struct {
			Text string `xml:",chardata"`
		} `xml:"filterDataType"`
		FilterGroup struct {
			Text string `xml:",chardata"`
		} `xml:"filterGroup"`
		FormulaConverted  *TextLiteral `xml:"formulaConverted"`
		FormulaExpression *TextLiteral `xml:"formulaExpression"`
		FilterOperator    *struct {
			Text string `xml:",chardata"`
		} `xml:"filterOperator"`
		FilterValue *struct {
			Text string `xml:",chardata"`
		} `xml:"filterValue"`
		FormulaResultPath *struct {
			Text string `xml:",chardata"`
		} `xml:"formulaResultPath"`
		FormulaSequence *struct {
			Text string `xml:",chardata"`
		} `xml:"formulaSequence"`
		GlobalKey struct {
			Text string `xml:",chardata"`
		} `xml:"globalKey"`
		InputFieldName *struct {
			Text string `xml:",chardata"`
		} `xml:"inputFieldName"`
		InputObjectName *struct {
			Text string `xml:",chardata"`
		} `xml:"inputObjectName"`
		InputObjectQuerySequence struct {
			Text string `xml:",chardata"`
		} `xml:"inputObjectQuerySequence"`
		LinkedObjectSequence struct {
			Text string `xml:",chardata"`
		} `xml:"linkedObjectSequence"`
		MigrationValue *struct {
			Text string `xml:",chardata"`
		} `xml:"migrationValue"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		OutputCreationSequence struct {
			Text string `xml:",chardata"`
		} `xml:"outputCreationSequence"`
		OutputFieldFormat *struct {
			Text string `xml:",chardata"`
		} `xml:"outputFieldFormat"`
		OutputFieldName struct {
			Text string `xml:",chardata"`
		} `xml:"outputFieldName"`
		OutputObjectName struct {
			Text string `xml:",chardata"`
		} `xml:"outputObjectName"`
		RequiredForUpsert struct {
			Text string `xml:",chardata"`
		} `xml:"requiredForUpsert"`
		TransformValuesMappings *struct {
			Text string `xml:",chardata"`
		} `xml:"transformValuesMappings"`
		UpsertKey struct {
			Text string `xml:",chardata"`
		} `xml:"upsertKey"`
	} `xml:"omniDataTransformItem"`
	OutputType struct {
		Text string `xml:",chardata"`
	} `xml:"outputType"`
	PreviewJsonData  *TextLiteral `xml:"previewJsonData"`
	ProcessSuperBulk struct {
		Text string `xml:",chardata"`
	} `xml:"processSuperBulk"`
	ResponseCacheTtlMinutes struct {
		Text string `xml:",chardata"`
	} `xml:"responseCacheTtlMinutes"`
	ResponseCacheType *struct {
		Text string `xml:",chardata"`
	} `xml:"responseCacheType"`
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
	TargetOutputFileName *struct {
		Text string `xml:",chardata"`
	} `xml:"targetOutputFileName"`
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
