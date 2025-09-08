package layout

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	layouts "github.com/ForceCLI/force-md/metadata/layouts"
)

var fieldName string

func init() {
	deleteFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	deleteFieldCmd.MarkFlagRequired("field")

	FieldCmd.AddCommand(deleteFieldCmd)
}

var FieldCmd = &cobra.Command{
	Use:   "field",
	Short: "Manage fields in page layouts",
}

var deleteFieldCmd = &cobra.Command{
	Use:   "delete -f FieldName [flags] [filename]...",
	Short: "Delete field from page layouts",
	Long:  "Delete field from page layouts",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteField(file, fieldName)
		}
	},
}

func deleteField(file string, fieldName string) {
	layout, err := layouts.Open(file)
	if err != nil {
		log.Warn("parsing layout failed: " + err.Error())
		return
	}

	err = layout.DeleteField(fieldName)
	if err != nil {
		log.Warn(fmt.Sprintf("delete failed for %s: %s", file, err.Error()))
		return
	}

	err = internal.WriteToFile(layout, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
