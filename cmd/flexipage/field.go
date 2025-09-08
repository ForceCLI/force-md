package flexipage

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	flexipages "github.com/ForceCLI/force-md/metadata/flexipages"
)

var fieldName string

func init() {
	deleteFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	deleteFieldCmd.MarkFlagRequired("field")

	FieldCmd.AddCommand(deleteFieldCmd)
}

var FieldCmd = &cobra.Command{
	Use:   "field",
	Short: "Manage fields in lightning pages",
}

var deleteFieldCmd = &cobra.Command{
	Use:   "delete -f FieldName [flags] [filename]...",
	Short: "Delete field from lightning pages",
	Long:  "Delete field from lightning pages",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteField(file, fieldName)
		}
	},
}

func deleteField(file string, fieldName string) {
	flexiPage, err := flexipages.Open(file)
	if err != nil {
		log.Warn("parsing flexipage failed: " + err.Error())
		return
	}

	err = flexiPage.DeleteField(fieldName)
	if err != nil {
		log.Warn(fmt.Sprintf("delete failed for %s: %s", file, err.Error()))
		return
	}

	err = internal.WriteToFile(flexiPage, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
