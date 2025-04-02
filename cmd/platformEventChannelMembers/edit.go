package platformEventChannelMembers

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/platformEventChannelMembers"
)

var (
	channel string
	entity  string
	filter  string
)

func init() {
	NewCmd.Flags().StringVarP(&channel, "channel", "c", "", "Channel")
	NewCmd.Flags().StringVarP(&entity, "entity", "e", "", "Entity")
	NewCmd.Flags().StringVarP(&filter, "filter", "f", "", "Filter")

	NewCmd.MarkFlagRequired("channel")
	NewCmd.MarkFlagRequired("entity")
	NewCmd.MarkFlagRequired("filter")
}

var NewCmd = &cobra.Command{
	Use:                   "new [filename]...",
	Short:                 "Create new platform event channel member file",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			createFile(file)
		}
	},
}

func createFile(file string) {
	p := platformEventChannelMembers.NewPlatformEventChannelMember(channel, entity, filter)
	err := internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("create failed: " + err.Error())
		return
	}
}
