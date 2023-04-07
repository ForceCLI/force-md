package workflow

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/workflow"
)

var (
	active bool
)

func init() {
	listRulesCmd.Flags().BoolVarP(&active, "active", "a", false, "active")
	RulesCmd.AddCommand(listRulesCmd)
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

func listRules(file string) {
	w, err := workflow.Open(file)
	if err != nil {
		log.Warn("parsing workflow failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".workflow")
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
		fmt.Printf("%s.%s: %s\n", objectName, r.FullName.Text, active)
	}
}
