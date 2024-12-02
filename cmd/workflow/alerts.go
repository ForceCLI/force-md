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
	"github.com/ForceCLI/force-md/metadata/workflow"
)

var (
	recipient  string
	group      string
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

	editAlertCmd.Flags().StringP("alert", "a", "", "alert name")
	editAlertCmd.Flags().StringP("sender", "s", "", "sender address")
	editAlertCmd.MarkFlagRequired("alert")

	deleteAlertCmd.Flags().StringP("alert", "a", "", "alert name")
	deleteAlertCmd.MarkFlagRequired("alert")

	AlertsCmd.AddCommand(listAlertsCmd)
	AlertsCmd.AddCommand(editAlertCmd)
	AlertsCmd.AddCommand(deleteAlertCmd)
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
		alertName, _ := cmd.Flags().GetString("alert")
		for _, file := range args {
			updateAlert(file, alertName, alertUpdates)
		}
	},
}

var deleteAlertCmd = &cobra.Command{
	Use:                   "delete -a AlertName [filename]...",
	Short:                 "Delete workflow alert",
	Long:                  "Delete workflow alert",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		alertName, _ := cmd.Flags().GetString("alert")
		for _, file := range args {
			deleteAlert(file, alertName)
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
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".workflow")
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

func updateAlert(file string, alertName string, alertUpdates workflow.Alert) {
	a, err := workflow.Open(file)
	if err != nil {
		log.Warn("parsing workflow failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".workflow")
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

func deleteAlert(file string, alertName string) {
	a, err := workflow.Open(file)
	if err != nil {
		log.Warn("parsing workflow failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".workflow")
	alertName = strings.ToLower(strings.TrimPrefix(alertName, objectName+"."))
	err = a.DeleteAlert(alertName)
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
