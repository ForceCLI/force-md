package custompermission

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/custompermission"
	"github.com/ForceCLI/force-md/internal"
)

var (
	label       string
	description string
)

var defaultVersion = "51.0"

func init() {
	NewCmd.Flags().StringVarP(&label, "label", "l", "", "label")

	NewCmd.MarkFlagRequired("label")
}

var NewCmd = &cobra.Command{
	Use:                   "new -l <label> [filename]...",
	Short:                 "Create new custom permission",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			createFile(file)
		}
	},
}

func createFile(file string) {
	p := custompermission.New(label)
	err := internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("create failed: " + err.Error())
		return
	}
}
