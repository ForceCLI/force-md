package workflow

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/workflow"
)

var (
	active bool
)

func init() {
	deleteRuleCmd.Flags().StringP("rule", "r", "", "rule name")
	deleteRuleCmd.MarkFlagRequired("alert")

	listRulesCmd.Flags().BoolVarP(&active, "active", "a", false, "active")
	RulesCmd.AddCommand(listRulesCmd)
	RulesCmd.AddCommand(deleteRuleCmd)
}

var RulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "Manage workflow rules",
}

var listRulesCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List workflow rules",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listRules(file)
		}
	},
}

var deleteRuleCmd = &cobra.Command{
	Use:                   "delete -a RuleName [filename]...",
	Short:                 "Delete workflow alert",
	Long:                  "Delete workflow alert",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ruleName, _ := cmd.Flags().GetString("rule")
		for _, file := range args {
			deleteRule(file, ruleName)
		}
	},
}

func listRules(file string) {
	w, err := workflow.Open(file)
	if err != nil {
		log.Warn("parsing workflow failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".workflow")
	var filters []workflow.RuleFilter
	if active {
		filters = append(filters, func(r workflow.Rule) bool {
			return r.Active.Text == "true"
		})
	}
	rules := w.GetRules(filters...)
	for _, r := range rules {
		active := "inactive"
		if r.Active.ToBool() {
			active = "active"
		}
		fmt.Printf("%s.%s: %s\n", objectName, r.FullName, active)
	}
}

func deleteRule(file string, ruleName string) {
	a, err := workflow.Open(file)
	if err != nil {
		log.Warn("parsing workflow failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".workflow")
	ruleName = strings.ToLower(strings.TrimPrefix(ruleName, objectName+"."))
	err = a.DeleteRule(ruleName)
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
