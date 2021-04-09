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
	addClassCmd.Flags().StringP("class", "c", "", "class name")
	addClassCmd.MarkFlagRequired("class")

	deleteClassCmd.Flags().StringVarP(&apexClassName, "class", "c", "", "apex classname")
	deleteClassCmd.MarkFlagRequired("class")

	ApexCmd.AddCommand(addClassCmd)
	ApexCmd.AddCommand(deleteClassCmd)
	ApexCmd.AddCommand(listClassesCmd)
}

var ApexCmd = &cobra.Command{
	Use:   "apex",
	Short: "Manage apex class visibility",
}

var addClassCmd = &cobra.Command{
	Use:                   "add -c ClassName [flags] [filename]...",
	Short:                 "Add Apex Class to Profile",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		className, _ := cmd.Flags().GetString("class")
		for _, file := range args {
			addClass(file, className)
		}
	},
}

var deleteClassCmd = &cobra.Command{
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

var listClassesCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List apex classes",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listClasses(file)
		}
	},
}

func addClass(file, className string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	p.AddClass(className)
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
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

func listClasses(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	classes := p.GetApexClasses()
	for _, a := range classes {
		if a.Enabled.Text == "true" {
			fmt.Println(a.ApexClass)
		}
	}
}
