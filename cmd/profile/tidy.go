package profile

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/profile"
)

var TidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "Tidy profile metadata",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tidy(file)
		}
	},
}

func tidy(file string) {
	p, err := profile.ParseProfile(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	p.Tidy()
	err = p.Write(file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
