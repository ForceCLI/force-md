package sharingrules

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/sharingrules"
)

var deleteRuleName string
var ruleType string

func init() {
	DeleteCmd.Flags().StringVarP(&deleteRuleName, "rule", "r", "", "rule name to delete")
	DeleteCmd.Flags().StringVarP(&ruleType, "type", "t", "", "rule type (criteria, owner) - if not specified, will try to delete from both")
	DeleteCmd.MarkFlagRequired("rule")
}

var DeleteCmd = &cobra.Command{
	Use:   "delete -r RuleName [flags] [filename]...",
	Short: "Delete both criteria-based and owner-based sharing rules",
	Long:  "Delete both criteria-based and owner-based sharing rules. Use --type to specify the rule type, or omit to try deleting from both types.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteRule(file, deleteRuleName, ruleType)
		}
	},
}

func deleteRule(file string, ruleName string, ruleType string) {
	p, err := sharingrules.Open(file)
	if err != nil {
		log.Warn("parsing sharing rules failed: " + err.Error())
		return
	}

	objectName := internal.TrimSuffixToEnd(path.Base(file), ".sharingRules")
	ruleName = strings.TrimPrefix(ruleName, objectName+".")

	deleted := false
	var deleteErrors []string

	// If type is specified, only try that type
	if ruleType != "" {
		switch strings.ToLower(ruleType) {
		case "criteria", "c":
			err = p.DeleteCriteriaRule(ruleName)
			if err != nil {
				deleteErrors = append(deleteErrors, fmt.Sprintf("criteria rule deletion failed: %s", err.Error()))
			} else {
				deleted = true
			}
		case "owner", "o":
			err = p.DeleteOwnerRule(ruleName)
			if err != nil {
				deleteErrors = append(deleteErrors, fmt.Sprintf("owner rule deletion failed: %s", err.Error()))
			} else {
				deleted = true
			}
		default:
			log.Warn(fmt.Sprintf("Invalid rule type '%s'. Use 'criteria' or 'owner'", ruleType))
			return
		}
	} else {
		// Try both types if no type specified
		// First try criteria rules
		err = p.DeleteCriteriaRule(ruleName)
		if err == nil {
			deleted = true
		} else {
			deleteErrors = append(deleteErrors, fmt.Sprintf("criteria rule deletion failed: %s", err.Error()))
		}

		// If criteria deletion failed, try owner rules
		if !deleted {
			err = p.DeleteOwnerRule(ruleName)
			if err == nil {
				deleted = true
				// Clear the error since we found it in owner rules
				deleteErrors = []string{}
			} else {
				deleteErrors = append(deleteErrors, fmt.Sprintf("owner rule deletion failed: %s", err.Error()))
			}
		}
	}

	if !deleted {
		if len(deleteErrors) > 0 {
			log.Warn(fmt.Sprintf("Failed to delete rule '%s' from %s: %s", ruleName, file, strings.Join(deleteErrors, "; ")))
		} else {
			log.Warn(fmt.Sprintf("Rule '%s' not found in %s", ruleName, file))
		}
		return
	}

	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
