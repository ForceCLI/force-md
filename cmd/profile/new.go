package profile

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/profile"
)

var licenseType string

func init() {
	NewCmd.Flags().StringVarP(&licenseType, "license", "l", "Salesforce", "license type")
}

var NewCmd = &cobra.Command{
	Use:   "new [flags] [filename]...",
	Short: "Create new profile",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addProfile(file)
		}
	},
}

func addProfile(file string) {
	p := &profile.Profile{
		UserLicense: licenseType,
		Custom:      TrueText,
	}
	err := internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("write failed: " + err.Error())
		return
	}
}
