package dashboard

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/dashboard"
	"github.com/octoberswimmer/force-md/internal"
)

var (
	runningUser string
)

func init() {
	EditCmd.Flags().StringVarP(&runningUser, "running-user", "r", "", "user dashboard runs as")
}

var EditCmd = &cobra.Command{
	Use:   "edit [flags] [filename]...",
	Short: "Edit dashboard",
	Long:  "Edit dashboard",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			updateDashboard(file)
		}
	},
}

func updateDashboard(file string) {
	a, err := dashboard.Open(file)
	if err != nil {
		log.Warn("parsing dashboard failed: " + err.Error())
		return
	}
	if runningUser != "" {
		a.UpdateRunningUser(runningUser)
	}
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(a, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
