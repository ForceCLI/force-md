package objects

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects"
)

var (
	ruleName string
	rulesDir string
)

func init() {
	deleteRuleCmd.Flags().StringVarP(&ruleName, "rule", "r", "", "rule name")
	deleteRuleCmd.MarkFlagRequired("rule")

	writeRulesCmd.Flags().StringVarP(&fieldsDir, "directory", "d", "", "directory where rules should be output")
	writeRulesCmd.MarkFlagRequired("directory")

	ValidationRuleCmd.AddCommand(deleteRuleCmd)
	ValidationRuleCmd.AddCommand(writeRulesCmd)
	ValidationRuleCmd.AddCommand(listRulesCmd)
}

var ValidationRuleCmd = &cobra.Command{
	Use:                   "validationrule",
	Short:                 "Manage validation rule metadata",
	DisableFlagsInUseLine: true,
}

var deleteRuleCmd = &cobra.Command{
	Use:   "delete -r ValidationRule [flags] [filename]...",
	Short: "Delete object validation rule",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteRule(file, ruleName)
		}
	},
}

var listRulesCmd = &cobra.Command{
	Use:   "list [filename]...",
	Short: "List validation rules",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listRules(file)
		}
	},
}

var writeRulesCmd = &cobra.Command{
	Use:                   "write -d directory [filename]...",
	Short:                 "Split object validation rules into separate files",
	Long:                  "Split object validation rules into separate metadata files to match sfdx's source format",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			writeRules(file, rulesDir)
		}
	},
}

func deleteRule(file string, ruleName string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
	ruleName = strings.TrimPrefix(ruleName, objectName+".")
	err = o.DeleteRule(ruleName)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(o, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func listRules(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
	rules := o.GetValidationRules()
	for _, r := range rules {
		fmt.Printf("%s.%s\n", objectName, r.FullName)
	}
}

func writeRules(file string, rulesDir string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	rules := o.GetValidationRules()
	for _, f := range rules {
		rule := objects.ValidationRule{
			Rule:  f,
			Xmlns: o.Xmlns,
		}
		err = internal.WriteToFile(rule, rulesDir+"/"+f.FullName+".validationRule-meta.xml")
		if err != nil {
			log.Warn("write failed: " + err.Error())
			return
		}
	}
}
