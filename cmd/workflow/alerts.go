package workflow

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/workflow"
)

func init() {
	AlertsCmd.AddCommand(listAlertsCmd)
}

var AlertsCmd = &cobra.Command{
	Use:   "alerts",
	Short: "Manage workflow alerts",
}

var listAlertsCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List workflow alerts",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listAlerts(file)
		}
	},
}

func listAlerts(file string) {
	w, err := workflow.Open(file)
	if err != nil {
		log.Warn("parsing workflow failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".workflow")
	alerts := w.GetAlerts()
	for _, r := range alerts {
		fmt.Printf("%s.%s\n", objectName, r.FullName.Text)
	}
}
