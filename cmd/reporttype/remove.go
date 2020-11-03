package reporttype

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/reporttype"
)

var fieldName string

func init() {
	RemoveCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	RemoveCmd.MarkFlagRequired("field")
}

var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove field",
	Long:  "Remove a field from Report Types",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			removeField(file)
		}
	},
}

func removeField(file string) {
	t, err := reporttype.Open(file)
	if err != nil {
		log.Warn("parsing report type failed: " + err.Error())
		return
	}
	if err = t.RemoveField(fieldName); err != nil {
		log.Warn(fmt.Sprintf("remove failed for %s: %s", file, err.Error()))
		return
	}
	if err = internal.WriteToFile(t, file); err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
