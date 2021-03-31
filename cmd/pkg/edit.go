package pkg

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/pkg"
)

var (
	metadataType string
	name         string
)

func init() {
	AddCmd.Flags().StringVarP(&metadataType, "type", "t", "", "metadata type")
	AddCmd.Flags().StringVarP(&name, "name", "n", "", "metadata item name")

	AddCmd.MarkFlagRequired("type")
	AddCmd.MarkFlagRequired("name")
}

var AddCmd = &cobra.Command{
	Use:                   "add -t Type -n Name [filename]...",
	Short:                 "Add metadata item to package.xml",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			add(file, metadataType, name)
		}
	},
}

var DeleteCmd = &cobra.Command{
	Use:                   "delete -t Type -n Name [filename]...",
	Short:                 "Remove metadata item from package.xml",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteMember(file, metadataType, name)
		}
	},
}

var TidyCmd = &cobra.Command{
	Use:                   "tidy [filename]...",
	Short:                 "Tidy package.xml",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tidy(file)
		}
	},
}

func add(file string, metadataType string, member string) {
	p, err := pkg.Open(file)
	if err != nil {
		log.Warn("parsing package.xml failed: " + err.Error())
		return
	}
	err = p.Add(metadataType, member)
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

func deleteMember(file string, metadataType string, member string) {
	p, err := pkg.Open(file)
	if err != nil {
		log.Warn("parsing package.xml failed: " + err.Error())
		return
	}
	err = p.Delete(metadataType, member)
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

func tidy(file string) {
	p, err := pkg.Open(file)
	if err != nil {
		log.Warn("parsing package.xml failed: " + err.Error())
		return
	}
	if err := general.Tidy(p, file); err != nil {
		log.Warn("tidying failed: " + err.Error())
	}
}
