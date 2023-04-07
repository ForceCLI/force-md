package workflow

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/workflow"
)

var (
	recipient  string
	group      string
	alertName  string
	senderType SenderType
)

type SenderType enumflag.Flag

const (
	None SenderType = iota
	CurrentUser
	DefaultWorkflowUser
	OrgWideEmailAddress
)

var SenderTypeIds = map[SenderType][]string{
	None:                {"None"},
	CurrentUser:         {"CurrentUser"},
	DefaultWorkflowUser: {"DefaultWorkflowUser"},
	OrgWideEmailAddress: {"OrgWideEmailAddress"},
}

func init() {
	listAlertsCmd.Flags().StringVarP(&recipient, "recipient", "r", "", "recipient or CC email address")
	listAlertsCmd.Flags().StringVarP(&group, "group", "g", "", "group recipient (DeveloperName)")
	listAlertsCmd.Flags().VarP(enumflag.New(&senderType, "sendertype", SenderTypeIds, enumflag.EnumCaseInsensitive),
		"sendertype", "t", "sender type; can be 'CurrentUser', 'DefaultWorkflowUser', or 'OrgWideEmailAddress'")

	editAlertCmd.Flags().StringVarP(&alertName, "alert", "a", "", "alert name")
	editAlertCmd.Flags().StringP("sender", "s", "", "sender address")
	editAlertCmd.MarkFlagRequired("alert")

	AlertsCmd.AddCommand(listAlertsCmd)
	AlertsCmd.AddCommand(editAlertCmd)
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

var editAlertCmd = &cobra.Command{
	Use:   "edit -a AlertName [flags] [filename]...",
	Short: "Edit workflow alert",
	Long:  "Edit workflow alert",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		alertUpdates := setFields(cmd)
		for _, file := range args {
			updateAlert(file, alertUpdates)
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
	if senderType != 0 {
		filters = append(filters, func(a workflow.Alert) bool {
			return strings.ToLower(a.SenderType.Text) == strings.ToLower(SenderTypeIds[senderType][0])
		})
	}
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
		fmt.Printf("%s.%s\n", objectName, r.FullName)
	}
}

func setFields(cmd *cobra.Command) workflow.Alert {
	alert := workflow.Alert{}
	alert.SenderAddress = TextValue(cmd, "sender")
	return alert
}

func updateAlert(file string, alertUpdates workflow.Alert) {
	a, err := workflow.Open(file)
	if err != nil {
		log.Warn("parsing workflow failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".workflow")
	alertName = strings.ToLower(strings.TrimPrefix(alertName, objectName+"."))
	err = a.UpdateAlert(alertName, alertUpdates)
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
