package platformEventSubscriberConfig

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/platformEventSubscriberConfig"
)

var (
	user string
)

func init() {
	EditCmd.Flags().StringVarP(&user, "user", "u", "", "user trigger runs as")
}

var EditCmd = &cobra.Command{
	Use:   "edit [flags] [filename]...",
	Short: "Edit Platform Event Subscriber Config",
	Long:  "Edit Platform Event Subscriber Config",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			updatePlatformEventSubscriberConfig(file)
		}
	},
}

func updatePlatformEventSubscriberConfig(file string) {
	a, err := platformEventSubscriberConfig.Open(file)
	if err != nil {
		log.Warn("parsing PlatformEventSubscriberConfig failed: " + err.Error())
		return
	}
	if user == "" {
		a.DeleteUser()
	} else {
		a.UpdateUser(user)
	}
	err = internal.WriteToFile(a, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
