package flexipage

import (
	. "github.com/ForceCLI/force-md/general"
)

// VisibilityRuleCriteria represents visibility rule criteria
type VisibilityRuleCriteria struct {
	LeftValue  TextLiteral  `xml:"leftValue"`
	Operator   TextLiteral  `xml:"operator"`
	RightValue *TextLiteral `xml:"rightValue"`
}

// VisibilityRule represents a visibility rule
type VisibilityRule struct {
	BooleanFilter *TextLiteral             `xml:"booleanFilter"`
	Criteria      []VisibilityRuleCriteria `xml:"criteria"`
}

// ComponentInstanceProperty represents a component instance property
type ComponentInstanceProperty struct {
	Name      TextLiteral  `xml:"name"`
	Type      *TextLiteral `xml:"type"`
	ValueList *struct {
		ValueListItems []struct {
			Value          TextLiteral     `xml:"value"`
			VisibilityRule *VisibilityRule `xml:"visibilityRule"`
		} `xml:"valueListItems"`
	} `xml:"valueList"`
	Value *TextLiteral `xml:"value"`
}

// ComponentInstance represents a component instance
type ComponentInstance struct {
	ComponentInstanceProperties []ComponentInstanceProperty `xml:"componentInstanceProperties"`
	ComponentName               TextLiteral                 `xml:"componentName"`
	Identifier                  TextLiteral                 `xml:"identifier"`
	VisibilityRule              *VisibilityRule             `xml:"visibilityRule"`
}

// FieldInstanceProperty represents a field instance property
type FieldInstanceProperty struct {
	Name  TextLiteral `xml:"name"`
	Value TextLiteral `xml:"value"`
}

// FieldInstance represents a field instance
type FieldInstance struct {
	FieldInstanceProperties []FieldInstanceProperty `xml:"fieldInstanceProperties"`
	FieldItem               TextLiteral             `xml:"fieldItem"`
	Identifier              TextLiteral             `xml:"identifier"`
	VisibilityRule          *VisibilityRule         `xml:"visibilityRule"`
}

// ItemInstance represents an item instance in a FlexiPage region
type ItemInstance struct {
	ComponentInstance *ComponentInstance `xml:"componentInstance"`
	FieldInstance     *FieldInstance     `xml:"fieldInstance"`
}

// FlexiPageRegion represents a region in a FlexiPage
type FlexiPageRegion struct {
	ItemInstances []ItemInstance `xml:"itemInstances"`
	Mode          *TextLiteral   `xml:"mode"`
	Name          TextLiteral    `xml:"name"`
	Type          TextLiteral    `xml:"type"`
}

// TemplateProperty represents a template property
type TemplateProperty struct {
	Name  TextLiteral `xml:"name"`
	Value TextLiteral `xml:"value"`
}

// Template represents a FlexiPage template
type Template struct {
	Name       TextLiteral        `xml:"name"`
	Properties []TemplateProperty `xml:"properties"`
}
