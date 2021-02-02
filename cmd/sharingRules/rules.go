package sharingRules

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/sharingRules"
)

func init() {
	CriteriaRulesCmd.AddCommand(listCriteriaRulesCmd)
	OwnerRulesCmd.AddCommand(listOwnerRulesCmd)
}

var CriteriaRulesCmd = &cobra.Command{
	Use:   "criteria-rules",
	Short: "Manage criteria-based sharing rules",
}

var OwnerRulesCmd = &cobra.Command{
	Use:   "owner-rules",
	Short: "Manage owner-based sharing rules",
}

var listCriteriaRulesCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List workflow rules",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listCriteriaRules(file)
		}
	},
}

var listOwnerRulesCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List workflow rules",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listOwnerRules(file)
		}
	},
}

func listCriteriaRules(file string) {
	w, err := sharingRules.Open(file)
	if err != nil {
		log.Warn("parsing sharing rules failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".sharingRules")
	rules := w.GetCriteriaRules()
	for _, r := range rules {
		fmt.Printf("%s.%s\n", objectName, r.FullName.Text)
	}
}

func listOwnerRules(file string) {
	w, err := sharingRules.Open(file)
	if err != nil {
		log.Warn("parsing sharing rules failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".sharingRules")
	rules := w.GetOwnerRules()
	for _, r := range rules {
		fmt.Printf("%s.%s\n", objectName, r.FullName.Text)
	}
}
