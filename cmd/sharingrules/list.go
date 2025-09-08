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

var criteriaField string

func init() {
	ListCmd.Flags().StringVarP(&criteriaField, "criteria-field", "f", "", "filter by sharing rules that use the specified field")
}

var ListCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List both criteria-based and owner-based sharing rules",
	Long:  "List both criteria-based and owner-based sharing rules. Use --criteria-field to filter by rules that use a specific field.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listAllRules(file, criteriaField)
		}
	},
}

func listAllRules(file string, criteriaFieldFilter string) {
	w, err := sharingrules.Open(file)
	if err != nil {
		log.Warn("parsing sharing rules failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".sharingRules")

	// List criteria-based rules
	criteriaRules := w.GetCriteriaRules()
	for _, r := range criteriaRules {
		// Apply field filter if specified
		if criteriaFieldFilter != "" {
			if !r.UsesField(criteriaFieldFilter) {
				continue
			}
		}

		// Format criteria rule output
		criteriaInfo := ""
		if len(r.CriteriaItems) > 0 {
			var conditions []string
			for _, item := range r.CriteriaItems {
				conditions = append(conditions, fmt.Sprintf("%s %s %s", item.Field.Text, item.Operation.Text, item.Value.Text))
			}
			criteriaInfo = fmt.Sprintf(" [%s]", strings.Join(conditions, " AND "))
		}

		sharedToInfo := getSharedToInfo(r.SharedTo.Role, r.SharedTo.Group, r.SharedTo.RoleAndSubordinates, r.SharedTo.AllInternalUsers != nil)

		fmt.Printf("CRITERIA: %s.%s -> %s (AccessLevel: %s)%s\n",
			objectName, r.FullName, sharedToInfo, r.AccessLevel.Text, criteriaInfo)
	}

	// List owner-based rules
	ownerRules := w.GetOwnerRules()
	for _, r := range ownerRules {
		// Owner rules don't have criteria fields, so skip if filtering by field
		if criteriaFieldFilter != "" {
			continue
		}

		sharedFromInfo := getSharedFromInfo(r.SharedFrom.Role, r.SharedFrom.Group, r.SharedFrom.RoleAndSubordinates, r.SharedFrom.Queue, r.SharedFrom.AllInternalUsers != nil)
		sharedToInfo := getSharedToInfo(r.SharedTo.Role, r.SharedTo.Group, r.SharedTo.RoleAndSubordinates, false)

		fmt.Printf("OWNER: %s.%s -> %s to %s (AccessLevel: %s)\n",
			objectName, r.FullName, sharedFromInfo, sharedToInfo, r.AccessLevel.Text)
	}
}

func getSharedToInfo(role *sharingrules.Role, group *sharingrules.Group, roleAndSubs *sharingrules.RoleAndSubordinates, allInternalUsers bool) string {
	if role != nil {
		return fmt.Sprintf("Role(%s)", role.Text)
	}
	if group != nil {
		return fmt.Sprintf("Group(%s)", group.Text)
	}
	if roleAndSubs != nil {
		return fmt.Sprintf("RoleAndSubs(%s)", roleAndSubs.Text)
	}
	if allInternalUsers {
		return "AllInternalUsers"
	}
	return "Unknown"
}

func getSharedFromInfo(role *sharingrules.Role, group *sharingrules.Group, roleAndSubs *sharingrules.RoleAndSubordinates, queue *sharingrules.Queue, allInternalUsers bool) string {
	if role != nil {
		return fmt.Sprintf("Role(%s)", role.Text)
	}
	if group != nil {
		return fmt.Sprintf("Group(%s)", group.Text)
	}
	if roleAndSubs != nil {
		return fmt.Sprintf("RoleAndSubs(%s)", roleAndSubs.Text)
	}
	if queue != nil {
		return fmt.Sprintf("Queue(%s)", queue.Text)
	}
	if allInternalUsers {
		return "AllInternalUsers"
	}
	return "Unknown"
}
