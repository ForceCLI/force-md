package permissionset

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/permissionset"
)

var TidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "Tidy Permission Set metadata",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tidy(file)
		}
	},
}

func tidy(file string) {
	p, err := permissionset.ParsePermissionSet(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
		return
	}
	p.Tidy()
	err = p.Write(file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
