package permissionset

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/permissionset"
)

var (
	label       string
	description string
)

func init() {
	NewCmd.Flags().StringVarP(&label, "label", "l", "", "label")
	NewCmd.Flags().StringVarP(&description, "description", "d", "", "description")
	NewCmd.MarkFlagRequired("label")
}

var NewCmd = &cobra.Command{
	Use:   "new [flags] [filename]...",
	Short: "Create new permission set",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addPermissionSet(file)
		}
	},
}

func addPermissionSet(file string) {
	p := &permissionset.PermissionSet{
		Xmlns:                 "http://soap.sforce.com/2006/04/metadata",
		HasActivationRequired: FalseText,
		Label:                 label,
	}
	if description != "" {
		p.Description = &permissionset.Description{
			Text: description,
		}
	}
	err := internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("write failed: " + err.Error())
		return
	}
}
