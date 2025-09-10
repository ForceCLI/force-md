package report

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	reports "github.com/ForceCLI/force-md/metadata/reports"
)

var fieldName string

func init() {
	deleteFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	deleteFieldCmd.MarkFlagRequired("field")

	FieldCmd.AddCommand(deleteFieldCmd)
}

var FieldCmd = &cobra.Command{
	Use:   "field",
	Short: "Manage fields in reports",
}

var deleteFieldCmd = &cobra.Command{
	Use:   "delete -f FieldName [flags] [filename]...",
	Short: "Delete field from reports",
	Long:  "Delete field from reports",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteField(file, fieldName)
		}
	},
}

func deleteField(file string, fieldName string) {
	report, err := reports.Open(file)
	if err != nil {
		log.Warn("parsing report failed: " + err.Error())
		return
	}

	err = report.DeleteField(fieldName)
	if err != nil {
		log.Warn(fmt.Sprintf("delete failed for %s: %s", file, err.Error()))
		return
	}

	err = internal.WriteToFile(report, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
