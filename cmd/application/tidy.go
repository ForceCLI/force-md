package application

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/application"
	"github.com/octoberswimmer/force-md/general"
)

var TidyCmd = &cobra.Command{
	Use:                   "tidy [filename]...",
	Short:                 "Tidy custom application",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tidy(file)
		}
	},
}

func tidy(file string) {
	p, err := application.Open(file)
	if err != nil {
		log.Warn("parsing application failed: " + err.Error())
		return
	}
	if err := general.Tidy(p, file); err != nil {
		log.Warn("tidying failed: " + err.Error())
	}
}
