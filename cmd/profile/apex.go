package profile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var apexClassName string

func init() {
	deleteApexClassCmd.Flags().StringVarP(&apexClassName, "class", "c", "", "apex classname")
	deleteApexClassCmd.MarkFlagRequired("class")

	ApexCmd.AddCommand(deleteApexClassCmd)
}

var ApexCmd = &cobra.Command{
	Use:   "apex",
	Short: "Manage apex class visibility",
}

var deleteApexClassCmd = &cobra.Command{
	Use:   "delete -c ClassName [flags] [filename]...",
	Short: "Delete apex class visibility",
	Long:  "Delete apex class visibility in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteApexClassVisibility(file, apexClassName)
		}
	},
}

func deleteApexClassVisibility(file string, apexClassName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.DeleteApexClassAccess(apexClassName)
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
