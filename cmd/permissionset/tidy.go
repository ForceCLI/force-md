package permissionset

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/permissionset"
)

var TidyCmd = &cobra.Command{
	Use:   "tidy [flags] [filename]...",
	Short: "Tidy Permission Set metadata",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tidy(file)
		}
	},
}

func tidy(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
		return
	}
	if err := general.Tidy(p, file); err != nil {
		log.Warn("tidying failed: " + err.Error())
	}
}
