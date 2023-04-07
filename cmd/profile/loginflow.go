package profile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/profile"
)

func init() {
	editLoginFlowCmd.Flags().BoolP("lightning", "l", false, "use lightning runtime")
	editLoginFlowCmd.Flags().BoolP("no-lightning", "L", false, "do not use lightning runtime")

	LoginFlowCmd.AddCommand(editLoginFlowCmd)
}

var LoginFlowCmd = &cobra.Command{
	Use:   "loginflow",
	Short: "Manage login flow",
}

var editLoginFlowCmd = &cobra.Command{
	Use:   "edit [flags] [filename]...",
	Short: "Update login flow",
	Long:  "Update login flow in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		settings := loginFlowSettingsToUpdate(cmd)
		for _, file := range args {
			updateLoginFlow(file, settings)
		}
	},
}

func loginFlowSettingsToUpdate(cmd *cobra.Command) profile.LoginFlow {
	settings := profile.LoginFlow{}
	settings.UseLightningRuntime = textValue(cmd, "lightning")
	return settings
}

func updateLoginFlow(file string, settings profile.LoginFlow) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.UpdateLoginFlow(settings)
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
