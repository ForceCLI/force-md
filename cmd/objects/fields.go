package objects

import (
	"encoding/xml"
	"fmt"
	"html"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/objects"
)

var (
	formulaField   bool
	filteredLookup bool
	fieldName      string
	fieldType      string
	references     string
	fieldsDir      string
	label          string
)

func init() {
	listFieldsCmd.Flags().BoolP("required", "r", false, "required fields")
	listFieldsCmd.Flags().BoolP("no-required", "R", false, "not required fields")
	listFieldsCmd.Flags().BoolP("history-tracking", "k", false, "with history tracking")
	listFieldsCmd.Flags().BoolP("no-history-tracking", "K", false, "without history tracking")
	listFieldsCmd.Flags().BoolP("trending", "d", false, "with trending tracking")
	listFieldsCmd.Flags().BoolP("no-trending", "D", false, "without trending tracking")
	listFieldsCmd.Flags().BoolVarP(&formulaField, "formula", "m", false, "formula fields only")
	listFieldsCmd.Flags().BoolVarP(&filteredLookup, "filtered-lookup", "f", false, "filtered lookup fields only")
	listFieldsCmd.Flags().BoolP("external-id", "x", false, "external id fields only")
	listFieldsCmd.Flags().BoolP("no-external-id", "X", false, "non-external id fields only")
	listFieldsCmd.Flags().BoolP("unique", "u", false, "unique fields only")
	listFieldsCmd.Flags().BoolP("no-unique", "U", false, "non-unique fields only")
	listFieldsCmd.Flags().StringVarP(&fieldType, "type", "t", "", "field type")
	listFieldsCmd.Flags().StringVarP(&label, "label", "l", "", "label")
	listFieldsCmd.Flags().StringVarP(&references, "references", "L", "", "references object")

	tableFieldsCmd.Flags().BoolP("required", "r", false, "required fields")
	tableFieldsCmd.Flags().BoolP("no-required", "R", false, "not required fields")
	tableFieldsCmd.Flags().BoolP("history-tracking", "k", false, "with history tracking")
	tableFieldsCmd.Flags().BoolP("no-history-tracking", "K", false, "without history tracking")
	tableFieldsCmd.Flags().BoolP("trending", "d", false, "with trending tracking")
	tableFieldsCmd.Flags().BoolP("no-trending", "D", false, "without trending tracking")
	tableFieldsCmd.Flags().BoolVarP(&formulaField, "formula", "m", false, "formula fields only")
	tableFieldsCmd.Flags().BoolVarP(&filteredLookup, "filtered-lookup", "f", false, "filtered lookup fields only")
	tableFieldsCmd.Flags().BoolP("external-id", "x", false, "external id fields only")
	tableFieldsCmd.Flags().BoolP("no-external-id", "X", false, "non-external id fields only")
	tableFieldsCmd.Flags().BoolP("unique", "u", false, "unique fields only")
	tableFieldsCmd.Flags().BoolP("no-unique", "U", false, "non-unique fields only")
	tableFieldsCmd.Flags().StringVarP(&fieldType, "type", "t", "", "field type")
	tableFieldsCmd.Flags().StringVarP(&label, "label", "l", "", "label")
	tableFieldsCmd.Flags().StringVarP(&references, "references", "L", "", "references object")

	addFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	addFieldCmd.MarkFlagRequired("field")

	editFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	editFieldCmd.Flags().StringP("label", "l", "", "label")
	editFieldCmd.Flags().StringP("type", "t", "", "field type")
	editFieldCmd.Flags().StringP("references", "L", "", "references object")
	editFieldCmd.Flags().StringP("relationship-name", "c", "", "relationship name")
	editFieldCmd.Flags().StringP("description", "d", "", "description")
	editFieldCmd.Flags().StringP("default", "v", "", "default value")
	editFieldCmd.Flags().StringP("inline-help", "i", "", "inline help")
	editFieldCmd.Flags().StringP("business-owner-user", "o", "", "business owner username")
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
	showFieldCmd.Flags().BoolVarP(&formulaField, "formula", "m", false, "show formula only")
	showFieldCmd.MarkFlagRequired("field")

	writeFieldsCmd.Flags().StringVarP(&fieldsDir, "directory", "d", "", "directory where fields should be output")
	writeFieldsCmd.MarkFlagRequired("directory")

	FieldCmd.AddCommand(listFieldsCmd)
	FieldCmd.AddCommand(tableFieldsCmd)
	FieldCmd.AddCommand(addFieldCmd)
	FieldCmd.AddCommand(editFieldCmd)
	FieldCmd.AddCommand(showFieldCmd)
	FieldCmd.AddCommand(deleteFieldCmd)
	FieldCmd.AddCommand(writeFieldsCmd)
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
		filterAttributes := setFields(cmd)
		for _, file := range args {
			listFields(file, filterAttributes)
		}
	},
}

var tableFieldsCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List object fields in a table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filterAttributes := setFields(cmd)
		tableFields(args, filterAttributes)
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
		fieldUpdates := setFields(cmd)
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

var writeFieldsCmd = &cobra.Command{
	Use:                   "write -d directory [filename]...",
	Short:                 "Split object fields into separate files",
	Long:                  "Split object fields into separate metadata files to match sfdx's source format",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			writeFields(file, fieldsDir)
		}
	},
}

var alwaysRequired map[string]bool = map[string]bool{
	"Name":    true,
	"OwnerId": true,
}

func listFields(file string, attributes objects.Field) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	var filters []objects.FieldFilter
	requiredFilter := func(f objects.Field) bool {
		isRequired := alwaysRequired[f.FullName] || (f.Required != nil && f.Required.Text == "true")
		isMasterDetail := f.Type != nil && f.Type.Text == "MasterDetail"
		return isRequired || isMasterDetail
	}
	if attributes.Required.IsTrue() {
		filters = append(filters, requiredFilter)
	}
	if attributes.Required.IsFalse() {
		filters = append(filters, func(f objects.Field) bool { return !requiredFilter(f) })
	}
	if attributes.TrackHistory.IsTrue() {
		filters = append(filters, func(f objects.Field) bool { return f.TrackHistory.ToBool() })
	}
	if attributes.TrackHistory.IsFalse() {
		filters = append(filters, func(f objects.Field) bool { return !f.TrackHistory.ToBool() })
	}
	if attributes.TrackTrending.IsTrue() {
		filters = append(filters, func(f objects.Field) bool { return f.TrackTrending.ToBool() })
	}
	if attributes.TrackTrending.IsFalse() {
		filters = append(filters, func(f objects.Field) bool { return !f.TrackTrending.ToBool() })
	}
	if formulaField {
		filters = append(filters, func(f objects.Field) bool { return f.Formula != nil })
	}
	if filteredLookup {
		filters = append(filters, func(f objects.Field) bool { return f.LookupFilter != nil })
	}
	if attributes.ExternalId.IsTrue() {
		filters = append(filters, func(f objects.Field) bool { return f.ExternalId.ToBool() })
	}
	if attributes.ExternalId.IsFalse() {
		filters = append(filters, func(f objects.Field) bool { return !f.ExternalId.ToBool() })
	}
	if attributes.Unique.IsTrue() {
		filters = append(filters, func(f objects.Field) bool { return f.Unique.ToBool() })
	}
	if attributes.Unique.IsFalse() {
		filters = append(filters, func(f objects.Field) bool { return !f.Unique.ToBool() })
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
	if label != "" {
		filters = append(filters, func(f objects.Field) bool {
			l := strings.ToLower(label)
			return f.Label != nil && strings.ToLower(f.Label.Text) == l
		})
	}
	fields := o.GetFields(filters...)
	for _, f := range fields {
		fmt.Printf("%s.%s\n", objectName, f.FullName)
	}
}

func tableFields(files []string, attributes objects.Field) {
	var filters []objects.FieldFilter
	requiredFilter := func(f objects.Field) bool {
		isRequired := alwaysRequired[f.FullName] || (f.Required != nil && f.Required.Text == "true")
		isMasterDetail := f.Type != nil && f.Type.Text == "MasterDetail"
		return isRequired || isMasterDetail
	}
	if attributes.Required.IsTrue() {
		filters = append(filters, requiredFilter)
	}
	if attributes.Required.IsFalse() {
		filters = append(filters, func(f objects.Field) bool { return !requiredFilter(f) })
	}
	if attributes.TrackHistory.IsTrue() {
		filters = append(filters, func(f objects.Field) bool { return f.TrackHistory.ToBool() })
	}
	if attributes.TrackHistory.IsFalse() {
		filters = append(filters, func(f objects.Field) bool { return !f.TrackHistory.ToBool() })
	}
	if attributes.TrackTrending.IsTrue() {
		filters = append(filters, func(f objects.Field) bool { return f.TrackTrending.ToBool() })
	}
	if attributes.TrackTrending.IsFalse() {
		filters = append(filters, func(f objects.Field) bool { return !f.TrackTrending.ToBool() })
	}
	if formulaField {
		filters = append(filters, func(f objects.Field) bool { return f.Formula != nil })
	}
	if filteredLookup {
		filters = append(filters, func(f objects.Field) bool { return f.LookupFilter != nil })
	}
	if attributes.ExternalId.IsTrue() {
		filters = append(filters, func(f objects.Field) bool { return f.ExternalId.ToBool() })
	}
	if attributes.ExternalId.IsFalse() {
		filters = append(filters, func(f objects.Field) bool { return !f.ExternalId.ToBool() })
	}
	if attributes.Unique.IsTrue() {
		filters = append(filters, func(f objects.Field) bool { return f.Unique.ToBool() })
	}
	if attributes.Unique.IsFalse() {
		filters = append(filters, func(f objects.Field) bool { return !f.Unique.ToBool() })
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
	if label != "" {
		filters = append(filters, func(f objects.Field) bool {
			l := strings.ToLower(label)
			return f.Label != nil && strings.ToLower(f.Label.Text) == l
		})
	}
	type field struct {
		fields objects.FieldList
		object string
	}
	var fields []field
	for _, file := range files {
		o, err := objects.Open(file)
		if err != nil {
			log.Warn("parsing object failed: " + err.Error())
			return
		}
		objectName := strings.TrimSuffix(path.Base(file), ".object")
		fields = append(fields, field{object: objectName, fields: o.GetFields(filters...)})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Object", "Field", "Type", "Label", "Required", "Description", "Help Text"})
	table.SetRowLine(true)
	for _, o := range fields {
		for _, f := range o.fields {
			table.Append([]string{o.object, f.FullName,
				f.Type.String(),
				f.Label.String(),
				strconv.FormatBool(f.Required.ToBool()),
				f.Description.String(),
				f.InlineHelpText.String(),
			})
		}
	}
	if table.NumLines() > 0 {
		table.Render()
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
	if formulaField {
		if fields[0].Formula != nil {
			fmt.Println(html.UnescapeString(fields[0].Formula.Text))
		}
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

func writeFields(file string, fieldsDir string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	fields := o.GetFields()
	for _, f := range fields {
		customField := objects.CustomField{
			Field: f,
			Xmlns: o.Xmlns,
		}
		err = internal.WriteToFile(customField, fieldsDir+"/"+f.FullName+".field-meta.xml")
		if err != nil {
			log.Warn("write failed: " + err.Error())
			return
		}
	}
}

func setFields(cmd *cobra.Command) objects.Field {
	field := objects.Field{}
	field.Label = TextValue(cmd, "label")
	field.Unique = BooleanTextValue(cmd, "unique")
	field.ExternalId = BooleanTextValue(cmd, "external-id")
	field.TrackHistory = BooleanTextValue(cmd, "history-tracking")
	field.TrackTrending = BooleanTextValue(cmd, "trending")
	field.Required = BooleanTextValue(cmd, "required")
	field.Description = TextValue(cmd, "description")
	field.Type = TextValue(cmd, "type")
	field.InlineHelpText = TextValue(cmd, "inline-help")
	field.BusinessOwnerUser = TextValue(cmd, "business-owner-user")
	field.ReferenceTo = TextValue(cmd, "references")
	field.RelationshipName = TextValue(cmd, "relationship-name")
	field.DefaultValue = TextValue(cmd, "default")
	field.Precision = IntegerValue(cmd, "precision")
	field.Scale = IntegerValue(cmd, "scale")
	field.Length = IntegerValue(cmd, "length")
	return field
}
