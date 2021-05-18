package reportType

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/reportType"
)

func init() {
	SectionCmd.AddCommand(listSectionsCmd)
}

var SectionCmd = &cobra.Command{
	Use:                   "sections",
	Short:                 "Manage report type sections",
	DisableFlagsInUseLine: true,
}

var listSectionsCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List report type sections",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listSections(file)
		}
	},
}

func listSections(file string) {
	o, err := reportType.Open(file)
	if err != nil {
		log.Warn("parsing report type failed: " + err.Error())
		return
	}
	sections := o.GetSections()
	for _, s := range sections {
		fmt.Printf("%s\n", s.MasterLabel)
	}
}
