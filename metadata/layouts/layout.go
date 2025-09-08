package layout

import (
	"encoding/xml"
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Layout"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Layout struct {
	metadata.MetadataInfo
	XMLName       xml.Name `xml:"Layout"`
	Xmlns         string   `xml:"xmlns,attr"`
	CustomButtons []struct {
		Text string `xml:",chardata"`
	} `xml:"customButtons"`
	EmailDefault *struct {
		Text string `xml:",chardata"`
	} `xml:"emailDefault"`
	ExcludeButtons []struct {
		Text string `xml:",chardata"`
	} `xml:"excludeButtons"`
	LayoutSections []struct {
		CustomLabel struct {
			Text string `xml:",chardata"`
		} `xml:"customLabel"`
		DetailHeading struct {
			Text string `xml:",chardata"`
		} `xml:"detailHeading"`
		EditHeading struct {
			Text string `xml:",chardata"`
		} `xml:"editHeading"`
		Label *struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		LayoutColumns []struct {
			LayoutItems []struct {
				Behavior *struct {
					Text string `xml:",chardata"`
				} `xml:"behavior"`
				Field *struct {
					Text string `xml:",chardata"`
				} `xml:"field"`
				CustomLink *struct {
					Text string `xml:",chardata"`
				} `xml:"customLink"`
				EmptySpace *struct {
					Text string `xml:",chardata"`
				} `xml:"emptySpace"`
				Height *struct {
					Text string `xml:",chardata"`
				} `xml:"height"`
				Page *struct {
					Text string `xml:",chardata"`
				} `xml:"page"`
				ShowLabel *struct {
					Text string `xml:",chardata"`
				} `xml:"showLabel"`
				ShowScrollbars *struct {
					Text string `xml:",chardata"`
				} `xml:"showScrollbars"`
				Width *struct {
					Text string `xml:",chardata"`
				} `xml:"width"`
			} `xml:"layoutItems"`
		} `xml:"layoutColumns"`
		Style struct {
			Text string `xml:",chardata"`
		} `xml:"style"`
	} `xml:"layoutSections"`
	MultilineLayoutFields []struct {
		Text string `xml:",chardata"`
	} `xml:"multilineLayoutFields"`
	MiniLayout *struct {
		Fields []struct {
			Text string `xml:",chardata"`
		} `xml:"fields"`
	} `xml:"miniLayout"`
	PlatformActionList *struct {
		ActionListContext struct {
			Text string `xml:",chardata"`
		} `xml:"actionListContext"`
		PlatformActionListItems []struct {
			ActionName struct {
				Text string `xml:",chardata"`
			} `xml:"actionName"`
			ActionType struct {
				Text string `xml:",chardata"`
			} `xml:"actionType"`
			SortOrder struct {
				Text string `xml:",chardata"`
			} `xml:"sortOrder"`
		} `xml:"platformActionListItems"`
	} `xml:"platformActionList"`
	QuickActionList *struct {
		QuickActionListItems []struct {
			QuickActionName struct {
				Text string `xml:",chardata"`
			} `xml:"quickActionName"`
		} `xml:"quickActionListItems"`
	} `xml:"quickActionList"`
	RelatedContent *struct {
		RelatedContentItems []struct {
			LayoutItem struct {
				Behavior struct {
					Text string `xml:",chardata"`
				} `xml:"behavior"`
				Field struct {
					Text string `xml:",chardata"`
				} `xml:"field"`
			} `xml:"layoutItem"`
		} `xml:"relatedContentItems"`
	} `xml:"relatedContent"`
	RelatedLists []struct {
		CustomButtons []struct {
			Text string `xml:",chardata"`
		} `xml:"customButtons"`
		ExcludeButtons []struct {
			Text string `xml:",chardata"`
		} `xml:"excludeButtons"`
		Fields []struct {
			Text string `xml:",chardata"`
		} `xml:"fields"`
		RelatedList struct {
			Text string `xml:",chardata"`
		} `xml:"relatedList"`
		SortField *struct {
			Text string `xml:",chardata"`
		} `xml:"sortField"`
		SortOrder *struct {
			Text string `xml:",chardata"`
		} `xml:"sortOrder"`
	} `xml:"relatedLists"`
	RelatedObjects []struct {
		Text string `xml:",chardata"`
	} `xml:"relatedObjects"`
	RunAssignmentRulesDefault *struct {
		Text string `xml:",chardata"`
	} `xml:"runAssignmentRulesDefault"`
	ShowEmailCheckbox struct {
		Text string `xml:",chardata"`
	} `xml:"showEmailCheckbox"`
	ShowHighlightsPanel *struct {
		Text string `xml:",chardata"`
	} `xml:"showHighlightsPanel"`
	ShowInteractionLogPanel *struct {
		Text string `xml:",chardata"`
	} `xml:"showInteractionLogPanel"`
	ShowKnowledgeComponent *struct {
		Text string `xml:",chardata"`
	} `xml:"showKnowledgeComponent"`
	ShowRunAssignmentRulesCheckbox *struct {
		Text string `xml:",chardata"`
	} `xml:"showRunAssignmentRulesCheckbox"`
	ShowSolutionSection *struct {
		Text string `xml:",chardata"`
	} `xml:"showSolutionSection"`
	ShowSubmitAndAttachButton struct {
		Text string `xml:",chardata"`
	} `xml:"showSubmitAndAttachButton"`
	SummaryLayout *struct {
		MasterLabel struct {
			Text string `xml:",chardata"`
		} `xml:"masterLabel"`
		SizeX struct {
			Text string `xml:",chardata"`
		} `xml:"sizeX"`
		SizeY struct {
			Text string `xml:",chardata"`
		} `xml:"sizeY"`
		SummaryLayoutItems []struct {
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			PosX struct {
				Text string `xml:",chardata"`
			} `xml:"posX"`
			PosY struct {
				Text string `xml:",chardata"`
			} `xml:"posY"`
		} `xml:"summaryLayoutItems"`
		SummaryLayoutStyle struct {
			Text string `xml:",chardata"`
		} `xml:"summaryLayoutStyle"`
	} `xml:"summaryLayout"`
}

func (c *Layout) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Layout) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Layout, error) {
	p := &Layout{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (l *Layout) DeleteField(fieldName string) error {
	fieldDeleted := false

	// Iterate through layout sections
	for sectionIdx := range l.LayoutSections {
		section := &l.LayoutSections[sectionIdx]

		// Iterate through layout columns in each section
		for columnIdx := range section.LayoutColumns {
			column := &section.LayoutColumns[columnIdx]

			// Filter out the field from layout items
			filteredItems := column.LayoutItems[:0]
			for _, item := range column.LayoutItems {
				if item.Field != nil && item.Field.Text == fieldName {
					fieldDeleted = true
					// Skip this item (don't add to filteredItems)
					continue
				}
				filteredItems = append(filteredItems, item)
			}

			column.LayoutItems = filteredItems
		}
	}

	// Also check and remove from multilineLayoutFields
	if l.MultilineLayoutFields != nil {
		filteredMultiline := l.MultilineLayoutFields[:0]
		for _, field := range l.MultilineLayoutFields {
			if field.Text == fieldName {
				fieldDeleted = true
				continue
			}
			filteredMultiline = append(filteredMultiline, field)
		}
		l.MultilineLayoutFields = filteredMultiline
	}

	// Also check and remove from miniLayout fields
	if l.MiniLayout != nil && l.MiniLayout.Fields != nil {
		filteredMini := l.MiniLayout.Fields[:0]
		for _, field := range l.MiniLayout.Fields {
			if field.Text == fieldName {
				fieldDeleted = true
				continue
			}
			filteredMini = append(filteredMini, field)
		}
		l.MiniLayout.Fields = filteredMini
	}

	// Also check and remove from relatedLists fields
	for i := range l.RelatedLists {
		relatedList := &l.RelatedLists[i]
		filteredFields := relatedList.Fields[:0]
		for _, field := range relatedList.Fields {
			if field.Text == fieldName {
				fieldDeleted = true
				continue
			}
			filteredFields = append(filteredFields, field)
		}
		relatedList.Fields = filteredFields
	}

	// Also check and remove from summaryLayout
	if l.SummaryLayout != nil && l.SummaryLayout.SummaryLayoutItems != nil {
		filteredSummary := l.SummaryLayout.SummaryLayoutItems[:0]
		for _, item := range l.SummaryLayout.SummaryLayoutItems {
			if item.Field.Text == fieldName {
				fieldDeleted = true
				continue
			}
			filteredSummary = append(filteredSummary, item)
		}
		l.SummaryLayout.SummaryLayoutItems = filteredSummary
	}

	if !fieldDeleted {
		return fmt.Errorf("field '%s' not found in layout", fieldName)
	}

	return nil
}
