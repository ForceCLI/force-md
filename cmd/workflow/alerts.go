package workflow

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/workflow"
)

var (
	recipient string
	group     string
)

func init() {
	listAlertsCmd.Flags().StringVarP(&recipient, "recipient", "r", "", "recipient or CC email address")
	listAlertsCmd.Flags().StringVarP(&group, "group", "g", "", "group recipient (DeveloperName)")
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
	var filters []workflow.AlertFilter
	objectName := strings.TrimSuffix(path.Base(file), ".workflow")
	if recipient != "" {
		filters = append(filters, func(a workflow.Alert) bool {
			t := strings.ToLower(recipient)
			for _, r := range a.Recipients {
				if strings.ToLower(r.Type.Text) == "user" && strings.ToLower(r.Recipient.Text) == t {
					return true
				}
			}
			for _, r := range a.CcEmails {
				if strings.ToLower(r.Text) == t {
					return true
				}
			}
			return false
		})
	}
	if group != "" {
		filters = append(filters, func(a workflow.Alert) bool {
			t := strings.ToLower(group)
			for _, r := range a.Recipients {
				if strings.ToLower(r.Type.Text) == "group" && strings.ToLower(r.Recipient.Text) == t {
					return true
				}
			}
			return false
		})
	}
	alerts := w.GetAlerts(filters...)
	for _, r := range alerts {
		fmt.Printf("%s.%s\n", objectName, r.FullName.Text)
	}
}
