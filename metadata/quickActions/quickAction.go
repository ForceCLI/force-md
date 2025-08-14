package quickAction

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "QuickAction"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type QuickAction struct {
	metadata.MetadataInfo
	XMLName       xml.Name `xml:"QuickAction"`
	Xmlns         string   `xml:"xmlns,attr"`
	ActionSubtype *struct {
		Text string `xml:",chardata"`
	} `xml:"actionSubtype"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	FieldOverrides []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Formula *struct {
			Text string `xml:",chardata"`
		} `xml:"formula"`
		LiteralValue *struct {
			Text string `xml:",chardata"`
		} `xml:"literalValue"`
	} `xml:"fieldOverrides"`
	FlowDefinition *struct {
		Text string `xml:",chardata"`
	} `xml:"flowDefinition"`
	Height *struct {
		Text string `xml:",chardata"`
	} `xml:"height"`
	Icon *struct {
		Text string `xml:",chardata"`
	} `xml:"icon"`
	Label *struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	LightningComponent *struct {
		Text string `xml:",chardata"`
	} `xml:"lightningComponent"`
	LightningWebComponent *struct {
		Text string `xml:",chardata"`
	} `xml:"lightningWebComponent"`
	OptionsCreateFeedItem struct {
		Text string `xml:",chardata"`
	} `xml:"optionsCreateFeedItem"`
	Page *struct {
		Text string `xml:",chardata"`
	} `xml:"page"`
	QuickActionLayout *struct {
		LayoutSectionStyle struct {
			Text string `xml:",chardata"`
		} `xml:"layoutSectionStyle"`
		QuickActionLayoutColumns []struct {
			QuickActionLayoutItems []struct {
				EmptySpace struct {
					Text string `xml:",chardata"`
				} `xml:"emptySpace"`
				Field *struct {
					Text string `xml:",chardata"`
				} `xml:"field"`
				UiBehavior struct {
					Text string `xml:",chardata"`
				} `xml:"uiBehavior"`
			} `xml:"quickActionLayoutItems"`
		} `xml:"quickActionLayoutColumns"`
	} `xml:"quickActionLayout"`
	StandardLabel *struct {
		Text string `xml:",chardata"`
	} `xml:"standardLabel"`
	SuccessMessage *struct {
		Text string `xml:",chardata"`
	} `xml:"successMessage"`
	TargetObject *struct {
		Text string `xml:",chardata"`
	} `xml:"targetObject"`
	TargetParentField *struct {
		Text string `xml:",chardata"`
	} `xml:"targetParentField"`
	TargetRecordType *struct {
		Text string `xml:",chardata"`
	} `xml:"targetRecordType"`
	QuickActionType struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	Width *struct {
		Text string `xml:",chardata"`
	} `xml:"width"`
}

func (c *QuickAction) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *QuickAction) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*QuickAction, error) {
	p := &QuickAction{}
	return p, metadata.ParseMetadataXml(p, path)
}
