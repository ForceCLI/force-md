package profile

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var licenseType string

var trueBooleanText = profile.BooleanText{
	Text: "true",
}

func init() {
	NewCmd.Flags().StringVarP(&licenseType, "license", "l", "Salesforce", "license type")
}

var NewCmd = &cobra.Command{
	Use:   "new",
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
		Custom:      trueBooleanText,
	}
	err := internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("write failed: " + err.Error())
		return
	}
}
