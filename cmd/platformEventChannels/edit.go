package platformEventChannels

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/platformEventChannels"
)

var (
	label string
)

func init() {
	NewCmd.Flags().StringVarP(&label, "label", "l", "", "Label")

	NewCmd.MarkFlagRequired("label")
}

var NewCmd = &cobra.Command{
	Use:                   "new [filename]...",
	Short:                 "Create new platform event channel file",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			createFile(file)
		}
	},
}

func createFile(file string) {
	p := platformEventChannels.NewPlatformEventChannel(label)
	err := internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("create failed: " + err.Error())
		return
	}
}
