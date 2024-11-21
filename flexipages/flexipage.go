package flexipage

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "FlexiPage"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type FlexiPage struct {
	internal.MetadataInfo
	XMLName          xml.Name `xml:"FlexiPage"`
	Xmlns            string   `xml:"xmlns,attr"`
	FlexiPageRegions []struct {
		ItemInstances []struct {
			ComponentInstance struct {
				ComponentInstanceProperties []struct {
					Name struct {
						Text string `xml:",chardata"`
					} `xml:"name"`
					ValueList struct {
						ValueListItems []struct {
							Value struct {
								Text string `xml:",chardata"`
							} `xml:"value"`
							VisibilityRule struct {
								BooleanFilter struct {
									Text string `xml:",chardata"`
								} `xml:"booleanFilter"`
								Criteria []struct {
									LeftValue struct {
										Text string `xml:",chardata"`
									} `xml:"leftValue"`
									Operator struct {
										Text string `xml:",chardata"`
									} `xml:"operator"`
									RightValue struct {
										Text string `xml:",chardata"`
									} `xml:"rightValue"`
								} `xml:"criteria"`
							} `xml:"visibilityRule"`
						} `xml:"valueListItems"`
					} `xml:"valueList"`
					Value struct {
						Text string `xml:",chardata"`
					} `xml:"value"`
					Type struct {
						Text string `xml:",chardata"`
					} `xml:"type"`
				} `xml:"componentInstanceProperties"`
				ComponentName struct {
					Text string `xml:",chardata"`
				} `xml:"componentName"`
				Identifier struct {
					Text string `xml:",chardata"`
				} `xml:"identifier"`
				VisibilityRule struct {
					Criteria []struct {
						LeftValue struct {
							Text string `xml:",chardata"`
						} `xml:"leftValue"`
						Operator struct {
							Text string `xml:",chardata"`
						} `xml:"operator"`
						RightValue struct {
							Text string `xml:",chardata"`
						} `xml:"rightValue"`
					} `xml:"criteria"`
					BooleanFilter struct {
						Text string `xml:",chardata"`
					} `xml:"booleanFilter"`
				} `xml:"visibilityRule"`
			} `xml:"componentInstance"`
			FieldInstance struct {
				FieldInstanceProperties struct {
					Name struct {
						Text string `xml:",chardata"`
					} `xml:"name"`
					Value struct {
						Text string `xml:",chardata"`
					} `xml:"value"`
				} `xml:"fieldInstanceProperties"`
				FieldItem struct {
					Text string `xml:",chardata"`
				} `xml:"fieldItem"`
				Identifier struct {
					Text string `xml:",chardata"`
				} `xml:"identifier"`
				VisibilityRule struct {
					Criteria []struct {
						LeftValue struct {
							Text string `xml:",chardata"`
						} `xml:"leftValue"`
						Operator struct {
							Text string `xml:",chardata"`
						} `xml:"operator"`
						RightValue struct {
							Text string `xml:",chardata"`
						} `xml:"rightValue"`
					} `xml:"criteria"`
					BooleanFilter struct {
						Text string `xml:",chardata"`
					} `xml:"booleanFilter"`
				} `xml:"visibilityRule"`
			} `xml:"fieldInstance"`
		} `xml:"itemInstances"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Type struct {
			Text string `xml:",chardata"`
		} `xml:"type"`
		Mode struct {
			Text string `xml:",chardata"`
		} `xml:"mode"`
	} `xml:"flexiPageRegions"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	SobjectType struct {
		Text string `xml:",chardata"`
	} `xml:"sobjectType"`
	Template struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Properties struct {
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			Value struct {
				Text string `xml:",chardata"`
			} `xml:"value"`
		} `xml:"properties"`
	} `xml:"template"`
	FlexiPageType struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	ParentFlexiPage struct {
		Text string `xml:",chardata"`
	} `xml:"parentFlexiPage"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *FlexiPage) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *FlexiPage) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*FlexiPage, error) {
	p := &FlexiPage{}
	return p, internal.ParseMetadataXml(p, path)
}
