package field

import (
	. "github.com/ForceCLI/force-md/general"
)

// Picklist is the picklist element of API versions before 38.0.
type Picklist struct {
	ControllingField   *TextLiteral    `xml:"controllingField"`
	PicklistValues     []PicklistValue `xml:"picklistValues"`
	RestrictedPicklist *BooleanText    `xml:"restrictedPicklist"`
	Sorted             *BooleanText    `xml:"sorted"`
}

type PicklistValue struct {
	FullName               string        `xml:"fullName"`
	AllowEmail             *BooleanText  `xml:"allowEmail"`
	Closed                 *BooleanText  `xml:"closed"`
	Color                  *TextLiteral  `xml:"color"`
	ControllingFieldValues []TextLiteral `xml:"controllingFieldValues"`
	Converted              *BooleanText  `xml:"converted"`
	CssExposed             *BooleanText  `xml:"cssExposed"`
	Default                *BooleanText  `xml:"default"`
	Description            *TextLiteral  `xml:"description"`
	ForecastCategory       *TextLiteral  `xml:"forecastCategory"`
	HighPriority           *BooleanText  `xml:"highPriority"`
	Probability            *IntegerText  `xml:"probability"`
	ReverseRole            *TextLiteral  `xml:"reverseRole"`
	Reviewed               *BooleanText  `xml:"reviewed"`
	Won                    *BooleanText  `xml:"won"`
}
