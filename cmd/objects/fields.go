package objects

import (
	"encoding/xml"
	"fmt"
	"path"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/objects"
)

var (
	externalId   bool
	requiredOnly bool
	withHistory  bool
	withTrending bool
	unique       bool
	formulaField bool
	fieldName    string
	fieldType    string
	references   string
)

func init() {
	listFieldsCmd.Flags().BoolVarP(&requiredOnly, "required", "r", false, "required fields only")
	listFieldsCmd.Flags().BoolVarP(&withHistory, "history", "k", false, "with history tracking")
	listFieldsCmd.Flags().BoolVarP(&withTrending, "trending", "d", false, "with trending tracking")
	listFieldsCmd.Flags().BoolVarP(&formulaField, "formula", "u", false, "formula fields only")
	listFieldsCmd.Flags().BoolVarP(&externalId, "external-id", "x", false, "external id fields only")
	listFieldsCmd.Flags().StringVarP(&fieldType, "type", "t", "", "field type")
	listFieldsCmd.Flags().StringVarP(&references, "references", "R", "", "references object")

	addFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	addFieldCmd.MarkFlagRequired("field")

	editFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	editFieldCmd.Flags().StringP("label", "l", "", "label")
	editFieldCmd.Flags().StringP("type", "t", "", "field type")
	editFieldCmd.Flags().StringP("description", "d", "", "description")
	editFieldCmd.Flags().StringP("default", "v", "", "default value")
	editFieldCmd.Flags().StringP("inline-help", "i", "", "inline help")
	editFieldCmd.Flags().IntP("precision", "p", 0, "precision")
	editFieldCmd.Flags().IntP("scale", "s", 0, "scale")
	editFieldCmd.Flags().IntP("length", "n", 0, "length")
	editFieldCmd.Flags().BoolP("required", "r", false, "required")
	editFieldCmd.Flags().BoolP("no-required", "R", false, "not required")
	editFieldCmd.Flags().BoolP("unique", "u", false, "unique")
	editFieldCmd.Flags().BoolP("no-unique", "U", false, "not unique")
	editFieldCmd.Flags().BoolP("external-id", "x", false, "external id")
	editFieldCmd.Flags().BoolP("no-external-id", "X", false, "not external id")
	editFieldCmd.Flags().BoolP("history-tracking", "k", false, "history tracking")
	editFieldCmd.Flags().BoolP("no-history-tracking", "K", false, "no history tracking")
	editFieldCmd.MarkFlagRequired("field")

	deleteFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	deleteFieldCmd.MarkFlagRequired("field")

	showFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	showFieldCmd.MarkFlagRequired("field")

	FieldCmd.AddCommand(listFieldsCmd)
	FieldCmd.AddCommand(addFieldCmd)
	FieldCmd.AddCommand(editFieldCmd)
	FieldCmd.AddCommand(showFieldCmd)
	FieldCmd.AddCommand(deleteFieldCmd)
}

var FieldCmd = &cobra.Command{
	Use:                   "fields",
	Short:                 "Manage object field metadata",
	DisableFlagsInUseLine: true,
}

var listFieldsCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List object fields",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listFields(file)
		}
	},
}

var addFieldCmd = &cobra.Command{
	Use:                   "add -f Field [filename]...",
	Short:                 "Add field",
	Long:                  "Add field to object",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addField(file, fieldName)
		}
	},
}

var editFieldCmd = &cobra.Command{
	Use:   "edit -f Field [flags] [filename]...",
	Short: "Edit object fields",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fieldUpdates := fieldUpdates(cmd)
		for _, file := range args {
			updateField(file, fieldUpdates)
		}
	},
}

var deleteFieldCmd = &cobra.Command{
	Use:   "delete -f Field [flags] [filename]...",
	Short: "Delete object field",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteField(file, fieldName)
		}
	},
}

var showFieldCmd = &cobra.Command{
	Use:                   "show -f Field [filename]...",
	Short:                 "Show object field",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			showField(file, fieldName)
		}
	},
}

var alwaysRequired map[string]bool = map[string]bool{
	"Name":    true,
	"OwnerId": true,
}

func listFields(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	var filters []objects.FieldFilter
	if requiredOnly {
		filters = append(filters, func(f objects.Field) bool {
			isRequired := alwaysRequired[f.FullName] || (f.Required != nil && f.Required.Text == "true")
			isMasterDetail := f.Type != nil && f.Type.Text == "MasterDetail"
			return isRequired || isMasterDetail
		})
	}
	if withHistory {
		filters = append(filters, func(f objects.Field) bool {
			return f.TrackHistory.ToBool()
		})
	}
	if withTrending {
		filters = append(filters, func(f objects.Field) bool {
			return f.TrackTrending.Text == "true"
		})
	}
	if formulaField {
		filters = append(filters, func(f objects.Field) bool {
			return f.Formula != nil
		})
	}
	if externalId {
		filters = append(filters, func(f objects.Field) bool {
			return f.ExternalId.ToBool()
		})
	}
	if fieldType != "" {
		filters = append(filters, func(f objects.Field) bool {
			t := strings.ToLower(fieldType)
			return f.Type != nil && strings.ToLower(f.Type.Text) == t
		})
	}
	if references != "" {
		filters = append(filters, func(f objects.Field) bool {
			r := strings.ToLower(references)
			return f.ReferenceTo != nil && strings.ToLower(f.ReferenceTo.Text) == r
		})
	}
	fields := o.GetFields(filters...)
	for _, f := range fields {
		fmt.Printf("%s.%s\n", objectName, f.FullName)
	}
}

func addField(file string, fieldName string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	err = o.AddField(fieldName)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(o, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func updateField(file string, fieldUpdates objects.Field) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	fieldName = strings.ToLower(strings.TrimPrefix(fieldName, objectName+"."))
	err = o.UpdateField(fieldName, fieldUpdates)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(o, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func showField(file string, fieldName string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	fieldName = strings.ToLower(strings.TrimPrefix(fieldName, objectName+"."))
	fields := o.GetFields(func(f objects.Field) bool {
		return strings.ToLower(f.FullName) == fieldName
	})
	if len(fields) == 0 {
		log.Warn(fmt.Sprintf("field not found in %s", file))
		return
	}
	b, err := xml.MarshalIndent(fields[0], "", "    ")
	if err != nil {
		log.Warn("marshal failed: " + err.Error())
		return
	}
	fmt.Println(string(b))
}

func deleteField(file string, fieldName string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	fieldName = strings.TrimPrefix(fieldName, objectName+".")
	err = o.DeleteField(fieldName)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(o, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func fieldUpdates(cmd *cobra.Command) objects.Field {
	field := objects.Field{}
	field.Label = textValue(cmd, "label")
	field.Unique = booleanTextValue(cmd, "unique")
	field.ExternalId = booleanTextValue(cmd, "external-id")
	field.TrackHistory = booleanTextValue(cmd, "history-tracking")
	field.Required = booleanTextValue(cmd, "required")
	field.Description = textValue(cmd, "description")
	field.Type = textValue(cmd, "type")
	field.InlineHelpText = textValue(cmd, "inline-help")
	field.DefaultValue = textValue(cmd, "default")
	field.Precision = integerValue(cmd, "precision")
	field.Scale = integerValue(cmd, "scale")
	field.Length = integerValue(cmd, "length")
	return field
}

func textValue(cmd *cobra.Command, flag string) (t *objects.TextLiteral) {
	if cmd.Flags().Changed(flag) {
		val, _ := cmd.Flags().GetString(flag)
		t = &objects.TextLiteral{
			Text: val,
		}
	}
	return t
}

func integerValue(cmd *cobra.Command, flag string) (t *IntegerText) {
	if cmd.Flags().Changed(flag) {
		val, _ := cmd.Flags().GetInt(flag)
		t = &IntegerText{
			Text: strconv.Itoa(val),
		}
	}
	return t
}

func booleanTextValue(cmd *cobra.Command, flag string) (t *BooleanText) {
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
