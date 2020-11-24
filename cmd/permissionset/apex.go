package permissionset

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/permissionset"
)

func init() {
	AddClassCmd.Flags().StringP("class", "c", "", "class name")
	AddClassCmd.MarkFlagRequired("class")
}

var AddClassCmd = &cobra.Command{
	Use:   "add-class",
	Short: "Add Apex Class to Permission Set",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		className, _ := cmd.Flags().GetString("class")
		for _, file := range args {
			addClass(file, className)
		}
	},
}

func addClass(file, className string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
		return
	}
	p.AddClass(className)
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
