package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/workflow"
)

func init() {
	workflowCmd.AddCommand(workflow.RulesCmd)
	workflowCmd.AddCommand(workflow.AlertsCmd)
	workflowCmd.AddCommand(workflow.FieldUpdatesCmd)
	RootCmd.AddCommand(workflowCmd)
}

var workflowCmd = &cobra.Command{
	Use:   "workflows",
	Short: "Manage Workflow",
}
