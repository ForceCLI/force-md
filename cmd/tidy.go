package cmd

import (
	"bytes"
	"fmt"
	"os"

	// Register types that don't have any other commands yet
	_ "github.com/ForceCLI/force-md/appMenus"
	_ "github.com/ForceCLI/force-md/assignmentRules"
	_ "github.com/ForceCLI/force-md/aura"
	_ "github.com/ForceCLI/force-md/autoResponseRules"
	_ "github.com/ForceCLI/force-md/classes"
	_ "github.com/ForceCLI/force-md/communities"
	_ "github.com/ForceCLI/force-md/components"
	_ "github.com/ForceCLI/force-md/connectedApps"
	_ "github.com/ForceCLI/force-md/contentassets"
	_ "github.com/ForceCLI/force-md/customHelpMenuSections"
	_ "github.com/ForceCLI/force-md/customPermissions"
	_ "github.com/ForceCLI/force-md/dashboardFolder"
	_ "github.com/ForceCLI/force-md/delegateGroups"
	_ "github.com/ForceCLI/force-md/documentFolder"
	_ "github.com/ForceCLI/force-md/documents"
	_ "github.com/ForceCLI/force-md/duplicateRules"
	_ "github.com/ForceCLI/force-md/emailFolder"
	_ "github.com/ForceCLI/force-md/emailTemplate"
	_ "github.com/ForceCLI/force-md/externalCredentials"
	_ "github.com/ForceCLI/force-md/flexipages"
	_ "github.com/ForceCLI/force-md/flow"
	_ "github.com/ForceCLI/force-md/groups"
	_ "github.com/ForceCLI/force-md/homePageComponents"
	_ "github.com/ForceCLI/force-md/homePageLayouts"
	_ "github.com/ForceCLI/force-md/installedPackages"
	_ "github.com/ForceCLI/force-md/layouts"
	_ "github.com/ForceCLI/force-md/letterhead"
	_ "github.com/ForceCLI/force-md/lwc"
	_ "github.com/ForceCLI/force-md/messageChannels"
	_ "github.com/ForceCLI/force-md/namedCredentials"
	_ "github.com/ForceCLI/force-md/notificationTypeConfig"
	_ "github.com/ForceCLI/force-md/notificationtypes"
	_ "github.com/ForceCLI/force-md/objectTranslations"
	_ "github.com/ForceCLI/force-md/omniDataTransforms"
	_ "github.com/ForceCLI/force-md/omniIntegrationProcedures"
	_ "github.com/ForceCLI/force-md/omniInteractionConfig"
	_ "github.com/ForceCLI/force-md/omniScripts"
	_ "github.com/ForceCLI/force-md/omniUiCard"
	_ "github.com/ForceCLI/force-md/pages"
	_ "github.com/ForceCLI/force-md/platformEventSubscriberConfigs"
	_ "github.com/ForceCLI/force-md/queues"
	_ "github.com/ForceCLI/force-md/quickActions"
	_ "github.com/ForceCLI/force-md/remoteSiteSettings"
	_ "github.com/ForceCLI/force-md/reports"
	_ "github.com/ForceCLI/force-md/restrictionrules"
	_ "github.com/ForceCLI/force-md/role"
	_ "github.com/ForceCLI/force-md/settings"
	_ "github.com/ForceCLI/force-md/sites"
	_ "github.com/ForceCLI/force-md/staticresources"
	_ "github.com/ForceCLI/force-md/tabs"
	_ "github.com/ForceCLI/force-md/triggers"
	_ "github.com/ForceCLI/force-md/wave"
	_ "github.com/ForceCLI/force-md/weblinks"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	tidyCmd.Flags().BoolP("list", "l", false, "list files that need tidying")

	RootCmd.AddCommand(tidyCmd)
}

var tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "Tidy Metadata",
	Example: `
$ force-md tidy sfdx/main/default/objects/*/{fields,validationRules}/* sfdx/main/default/flows/*

$ force-md tidy src/objects/*
`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, file := range args {
			_, err := internal.Metadata.Open(file)
			if err != nil {
				return fmt.Errorf("invalid file %s: %w", file, err)
			}
		}
		changes := false
		list, _ := cmd.Flags().GetBool("list")
		for _, t := range internal.Metadata.Types() {
			for _, m := range internal.Metadata.Items(t) {
				file := m.GetMetadataInfo().Path()
				o, ok := m.(general.Tidyable)
				if !ok {
					log.Warnf("file %s of type %s is not tidyable", file, m.Type())
					continue
				}
				if list {
					orig := m.GetMetadataInfo().Contents()
					needsTidying := checkIfChanged(o, orig)
					if needsTidying {
						fmt.Println(file)
					}
					changes = needsTidying || changes
				} else {
					if err := general.Tidy(o, file); err != nil {
						log.Warnf("tidying failed: %s", err.Error())
					}
				}
			}
		}
		if changes {
			os.Exit(1)
		}
		return nil
	},
}

func checkIfChanged(o general.Tidyable, orig []byte) (changed bool) {
	o.Tidy()
	newContents, err := internal.Marshal(o)
	if err != nil {
		log.Warn("serializing failed: " + err.Error())
		return false
	}
	if !bytes.Equal(orig, newContents) {
		return true
	}
	return false
}
