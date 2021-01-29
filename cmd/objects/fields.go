package objects

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/objects"
)

var requiredOnly bool
var fieldLabel, fieldName string

func init() {
	listFieldsCmd.Flags().BoolVarP(&requiredOnly, "required", "r", false, "required fields only")

	editFieldCmd.Flags().StringVarP(&fieldLabel, "label", "l", "", "field label")
	editFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	editFieldCmd.MarkFlagRequired("field")

	FieldCmd.AddCommand(listFieldsCmd)
	FieldCmd.AddCommand(editFieldCmd)
}

var FieldCmd = &cobra.Command{
	Use:   "field",
	Short: "Manage object field metadata",
}

var listFieldsCmd = &cobra.Command{
	Use:   "list",
	Short: "List object fields",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listFields(file)
		}
	},
}

var editFieldCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit object fields",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fieldUpdates := fieldUpdates(cmd)
		for _, file := range args {
			updateField(file, fieldUpdates)
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
			isRequired := alwaysRequired[f.FullName.Text] || (f.Required != nil && f.Required.Text == "true")
			isMasterDetail := f.Type != nil && f.Type.Text == "MasterDetail"
			return isRequired || isMasterDetail
		})
	}
	fields := o.GetFields(filters...)
	for _, f := range fields {
		fmt.Printf("%s.%s\n", objectName, f.FullName.Text)
	}
}

func updateField(file string, fieldUpdates objects.Field) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
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

func fieldUpdates(cmd *cobra.Command) objects.Field {
	field := objects.Field{}
	field.Label = textValue(cmd, "label")
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
