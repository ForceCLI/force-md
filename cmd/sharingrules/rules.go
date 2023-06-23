package sharingrules

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/sharingrules"
)

var ruleName string

func init() {
	deleteCriteriaRulesCmd.Flags().StringVarP(&ruleName, "rule", "r", "", "rule name")
	deleteCriteriaRulesCmd.MarkFlagRequired("rule")
	deleteOwnerRulesCmd.Flags().StringVarP(&ruleName, "rule", "r", "", "rule name")
	deleteOwnerRulesCmd.MarkFlagRequired("rule")

	CriteriaCmd.AddCommand(listCriteriaRulesCmd)
	CriteriaCmd.AddCommand(deleteCriteriaRulesCmd)

	OwnerCmd.AddCommand(listOwnerRulesCmd)
	OwnerCmd.AddCommand(deleteOwnerRulesCmd)
}

var CriteriaCmd = &cobra.Command{
	Use:   "criteria",
	Short: "Manage criteria-based sharing rules",
}

var OwnerCmd = &cobra.Command{
	Use:   "owner",
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
	w, err := sharingrules.Open(file)
	if err != nil {
		log.Warn("parsing sharing rules failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".sharingRules")
	rules := w.GetCriteriaRules()
	for _, r := range rules {
		fmt.Printf("%s.%s\n", objectName, r.FullName.Text)
	}
}

func listOwnerRules(file string) {
	w, err := sharingrules.Open(file)
	if err != nil {
		log.Warn("parsing sharing rules failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".sharingRules")
	rules := w.GetOwnerRules()
	for _, r := range rules {
		fmt.Printf("%s.%s\n", objectName, r.FullName.Text)
	}
}

var deleteCriteriaRulesCmd = &cobra.Command{
	Use:                   "delete -r RuleName [filename]...",
	Short:                 "Delete rule",
	Long:                  "Delete criteria-based sharing rule",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteCriteriaRule(file, ruleName)
		}
	},
}

var deleteOwnerRulesCmd = &cobra.Command{
	Use:                   "delete -r RuleName [filename]...",
	Short:                 "Delete rule",
	Long:                  "Delete owner-based sharing rule",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteOwnerRule(file, ruleName)
		}
	},
}

func deleteCriteriaRule(file string, ruleName string) {
	p, err := sharingrules.Open(file)
	if err != nil {
		log.Warn("parsing sharing rules failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".sharingRules")
	ruleName = strings.TrimPrefix(ruleName, objectName+".")
	err = p.DeleteCriteriaRule(ruleName)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func deleteOwnerRule(file string, ruleName string) {
	p, err := sharingrules.Open(file)
	if err != nil {
		log.Warn("parsing sharing rules failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".sharingRules")
	ruleName = strings.TrimPrefix(ruleName, objectName+".")
	err = p.DeleteOwnerRule(ruleName)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
