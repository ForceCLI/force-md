package quickAction

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "QuickAction"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type QuickAction struct {
	internal.MetadataInfo
	XMLName xml.Name `xml:"QuickAction"`
	Xmlns   string   `xml:"xmlns,attr"`
	Label   struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	OptionsCreateFeedItem struct {
		Text string `xml:",chardata"`
	} `xml:"optionsCreateFeedItem"`
	QuickActionLayout struct {
		LayoutSectionStyle struct {
			Text string `xml:",chardata"`
		} `xml:"layoutSectionStyle"`
		QuickActionLayoutColumns []struct {
			QuickActionLayoutItems []struct {
				EmptySpace struct {
					Text string `xml:",chardata"`
				} `xml:"emptySpace"`
				Field struct {
					Text string `xml:",chardata"`
				} `xml:"field"`
				UiBehavior struct {
					Text string `xml:",chardata"`
				} `xml:"uiBehavior"`
			} `xml:"quickActionLayoutItems"`
		} `xml:"quickActionLayoutColumns"`
	} `xml:"quickActionLayout"`
	TargetObject struct {
		Text string `xml:",chardata"`
	} `xml:"targetObject"`
	TargetParentField struct {
		Text string `xml:",chardata"`
	} `xml:"targetParentField"`
	TargetRecordType struct {
		Text string `xml:",chardata"`
	} `xml:"targetRecordType"`
	QuickActionType struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	FieldOverrides []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Formula struct {
			Text string `xml:",chardata"`
		} `xml:"formula"`
		LiteralValue struct {
			Text string `xml:",chardata"`
		} `xml:"literalValue"`
	} `xml:"fieldOverrides"`
	StandardLabel struct {
		Text string `xml:",chardata"`
	} `xml:"standardLabel"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	SuccessMessage struct {
		Text string `xml:",chardata"`
	} `xml:"successMessage"`
	FlowDefinition struct {
		Text string `xml:",chardata"`
	} `xml:"flowDefinition"`
	Height struct {
		Text string `xml:",chardata"`
	} `xml:"height"`
	Page struct {
		Text string `xml:",chardata"`
	} `xml:"page"`
	Width struct {
		Text string `xml:",chardata"`
	} `xml:"width"`
	ActionSubtype struct {
		Text string `xml:",chardata"`
	} `xml:"actionSubtype"`
	LightningWebComponent struct {
		Text string `xml:",chardata"`
	} `xml:"lightningWebComponent"`
	LightningComponent struct {
		Text string `xml:",chardata"`
	} `xml:"lightningComponent"`
	Icon struct {
		Text string `xml:",chardata"`
	} `xml:"icon"`
}

func (c *QuickAction) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *QuickAction) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*QuickAction, error) {
	p := &QuickAction{}
	return p, internal.ParseMetadataXml(p, path)
}
