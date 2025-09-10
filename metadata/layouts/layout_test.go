package layout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteField(t *testing.T) {
	// Create a test layout with field instances in various places
	testLayout := &Layout{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
		LayoutSections: []struct {
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
		}{
			{
				CustomLabel: struct {
					Text string `xml:",chardata"`
				}{Text: "true"},
				DetailHeading: struct {
					Text string `xml:",chardata"`
				}{Text: "false"},
				EditHeading: struct {
					Text string `xml:",chardata"`
				}{Text: "true"},
				Label: &struct {
					Text string `xml:",chardata"`
				}{Text: "Information"},
				LayoutColumns: []struct {
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
				}{
					{
						LayoutItems: []struct {
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
						}{
							{
								Field: &struct {
									Text string `xml:",chardata"`
								}{Text: "Account.Name"},
								Behavior: &struct {
									Text string `xml:",chardata"`
								}{Text: "Required"},
							},
							{
								Field: &struct {
									Text string `xml:",chardata"`
								}{Text: "TestField__c"},
								Behavior: &struct {
									Text string `xml:",chardata"`
								}{Text: "Edit"},
							},
						},
					},
					{
						LayoutItems: []struct {
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
						}{
							{
								Field: &struct {
									Text string `xml:",chardata"`
								}{Text: "Account.Type"},
								Behavior: &struct {
									Text string `xml:",chardata"`
								}{Text: "Edit"},
							},
						},
					},
				},
				Style: struct {
					Text string `xml:",chardata"`
				}{Text: "TwoColumnsTopToBottom"},
			},
		},
		MultilineLayoutFields: []struct {
			Text string `xml:",chardata"`
		}{
			{Text: "TestField__c"},
			{Text: "Description"},
		},
		MiniLayout: &struct {
			Fields []struct {
				Text string `xml:",chardata"`
			} `xml:"fields"`
		}{
			Fields: []struct {
				Text string `xml:",chardata"`
			}{
				{Text: "Account.Name"},
				{Text: "TestField__c"},
			},
		},
		RelatedLists: []struct {
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
		}{
			{
				RelatedList: struct {
					Text string `xml:",chardata"`
				}{Text: "Contacts"},
				Fields: []struct {
					Text string `xml:",chardata"`
				}{
					{Text: "FULL_NAME"},
					{Text: "TestField__c"},
					{Text: "CONTACT.EMAIL"},
				},
			},
		},
		SummaryLayout: &struct {
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
		}{
			SummaryLayoutItems: []struct {
				Field struct {
					Text string `xml:",chardata"`
				} `xml:"field"`
				PosX struct {
					Text string `xml:",chardata"`
				} `xml:"posX"`
				PosY struct {
					Text string `xml:",chardata"`
				} `xml:"posY"`
			}{
				{
					Field: struct {
						Text string `xml:",chardata"`
					}{Text: "Account.Name"},
					PosX: struct {
						Text string `xml:",chardata"`
					}{Text: "0"},
					PosY: struct {
						Text string `xml:",chardata"`
					}{Text: "0"},
				},
				{
					Field: struct {
						Text string `xml:",chardata"`
					}{Text: "TestField__c"},
					PosX: struct {
						Text string `xml:",chardata"`
					}{Text: "1"},
					PosY: struct {
						Text string `xml:",chardata"`
					}{Text: "0"},
				},
			},
		},
	}

	t.Run("delete_existing_field_from_all_locations", func(t *testing.T) {
		// Delete TestField__c which appears in multiple places
		err := testLayout.DeleteField("TestField__c")
		assert.NoError(t, err)

		// Verify the field was deleted from layoutSections
		fieldFoundInLayout := false
		for _, section := range testLayout.LayoutSections {
			for _, column := range section.LayoutColumns {
				for _, item := range column.LayoutItems {
					if item.Field != nil && item.Field.Text == "TestField__c" {
						fieldFoundInLayout = true
					}
				}
			}
		}
		assert.False(t, fieldFoundInLayout, "TestField__c should have been deleted from layout sections")

		// Verify the field was deleted from multilineLayoutFields
		fieldFoundInMultiline := false
		for _, field := range testLayout.MultilineLayoutFields {
			if field.Text == "TestField__c" {
				fieldFoundInMultiline = true
			}
		}
		assert.False(t, fieldFoundInMultiline, "TestField__c should have been deleted from multiline fields")

		// Verify the field was deleted from miniLayout
		fieldFoundInMini := false
		if testLayout.MiniLayout != nil {
			for _, field := range testLayout.MiniLayout.Fields {
				if field.Text == "TestField__c" {
					fieldFoundInMini = true
				}
			}
		}
		assert.False(t, fieldFoundInMini, "TestField__c should have been deleted from mini layout")

		// Verify the field was deleted from relatedLists
		fieldFoundInRelated := false
		for _, relatedList := range testLayout.RelatedLists {
			for _, field := range relatedList.Fields {
				if field.Text == "TestField__c" {
					fieldFoundInRelated = true
				}
			}
		}
		assert.False(t, fieldFoundInRelated, "TestField__c should have been deleted from related lists")

		// Verify the field was deleted from summaryLayout
		fieldFoundInSummary := false
		if testLayout.SummaryLayout != nil {
			for _, item := range testLayout.SummaryLayout.SummaryLayoutItems {
				if item.Field.Text == "TestField__c" {
					fieldFoundInSummary = true
				}
			}
		}
		assert.False(t, fieldFoundInSummary, "TestField__c should have been deleted from summary layout")

		// Verify other fields are still present
		assert.Equal(t, "Account.Name", testLayout.LayoutSections[0].LayoutColumns[0].LayoutItems[0].Field.Text)
		assert.Equal(t, "Account.Type", testLayout.LayoutSections[0].LayoutColumns[1].LayoutItems[0].Field.Text)
		assert.Equal(t, 1, len(testLayout.MultilineLayoutFields))
		assert.Equal(t, "Description", testLayout.MultilineLayoutFields[0].Text)
	})

	t.Run("delete_nonexistent_field", func(t *testing.T) {
		// Try to delete a field that doesn't exist
		err := testLayout.DeleteField("NonExistentField__c")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'NonExistentField__c' not found")
	})
}
