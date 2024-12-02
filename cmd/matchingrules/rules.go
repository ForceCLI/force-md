package matchingrules

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/matchingrules"
)

var ruleName string

func init() {
	DeleteCmd.Flags().StringVarP(&ruleName, "rule", "r", "", "rule name")
	DeleteCmd.MarkFlagRequired("rule")
}

var ListCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List matching rules",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listRules(file)
		}
	},
}

func listRules(file string) {
	w, err := matchingrules.Open(file)
	if err != nil {
		log.Warn("parsing matching rules failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".matchingRule")
	rules := w.GetMatchingRules()
	for _, r := range rules {
		fmt.Printf("%s.%s\n", objectName, r.FullName.Text)
	}
}

var DeleteCmd = &cobra.Command{
	Use:                   "delete -r RuleName [filename]...",
	Short:                 "Delete rule",
	Long:                  "Delete matching rule",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteRule(file, ruleName)
		}
	},
}

func deleteRule(file string, ruleName string) {
	p, err := matchingrules.Open(file)
	if err != nil {
		log.Warn("parsing matching rules failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".matchingRule")
	ruleName = strings.TrimPrefix(ruleName, objectName+".")
	err = p.DeleteRule(ruleName)
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
