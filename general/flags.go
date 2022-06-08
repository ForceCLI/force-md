package general

import (
	"strconv"

	"github.com/spf13/cobra"
)

func TextValue(cmd *cobra.Command, flag string) (t *TextLiteral) {
	if cmd.Flags().Changed(flag) {
		val, _ := cmd.Flags().GetString(flag)
		t = &TextLiteral{
			Text: val,
		}
	}
	return t
}

func IntegerValue(cmd *cobra.Command, flag string) (t *IntegerText) {
	if cmd.Flags().Changed(flag) {
		val, _ := cmd.Flags().GetInt(flag)
		t = &IntegerText{
			Text: strconv.Itoa(val),
		}
	}
	return t
}

func BooleanTextValue(cmd *cobra.Command, flag string) (t *BooleanText) {
	if cmd.Flags().Changed(flag) {
		val, _ := cmd.Flags().GetBool(flag)
		t = &BooleanText{
			Text: strconv.FormatBool(val),
		}
	}
	antiFlag := "no-" + flag
	if cmd.Flags().Changed(antiFlag) {
		val, _ := cmd.Flags().GetBool(antiFlag)
		t = &BooleanText{
			Text: strconv.FormatBool(!val),
		}
	}
	return t
}
