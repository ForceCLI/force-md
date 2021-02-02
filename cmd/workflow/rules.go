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
	rules := w.GetRules()
	for _, r := range rules {
		fmt.Printf("%s.%s\n", objectName, r.FullName.Text)
	}
}
