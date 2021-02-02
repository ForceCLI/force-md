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
	FieldUpdatesCmd.AddCommand(listFieldUpdatesCmd)
}

var FieldUpdatesCmd = &cobra.Command{
	Use:   "field-updates",
	Short: "Manage workflow field updates",
}

var listFieldUpdatesCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List workflow field updates",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listFieldUpdates(file)
		}
	},
}

func listFieldUpdates(file string) {
	w, err := workflow.Open(file)
	if err != nil {
		log.Warn("parsing workflow failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".workflow")
	fieldUpdates := w.GetFieldUpdates()
	for _, r := range fieldUpdates {
		fmt.Printf("%s.%s\n", objectName, r.FullName.Text)
	}
}
