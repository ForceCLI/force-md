package profile

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var layoutName string
var recordType string

func init() {
	showLayoutCmd.Flags().StringVarP(&objectName, "object", "o", "", "object name")
	showLayoutCmd.MarkFlagRequired("object")

	editLayoutCmd.Flags().StringVarP(&objectName, "object", "o", "", "object name")
	editLayoutCmd.Flags().StringVarP(&layoutName, "layout", "l", "", "layout name")
	editLayoutCmd.Flags().StringVarP(&recordType, "recordtype", "r", "", "record type")
	editLayoutCmd.MarkFlagRequired("layout")
	editLayoutCmd.MarkFlagRequired("object")

	deleteLayoutCmd.Flags().StringVarP(&objectName, "object", "o", "", "object name")
	deleteLayoutCmd.MarkFlagRequired("object")

	LayoutCmd.AddCommand(showLayoutCmd)
	LayoutCmd.AddCommand(editLayoutCmd)
	LayoutCmd.AddCommand(deleteLayoutCmd)
}

var LayoutCmd = &cobra.Command{
	Use:   "layout",
	Short: "Manage page layouts",
}

var showLayoutCmd = &cobra.Command{
	Use:                   "show -o SObject [filename]...",
	Short:                 "Show page layout assignment",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			showLayout(file)
		}
	},
}

var editLayoutCmd = &cobra.Command{
	Use:                   "edit -o SObject -l Layout [filename]...",
	Short:                 "Show page layout assignment",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			editLayout(file, objectName, layoutName, recordType)
		}
	},
}

var deleteLayoutCmd = &cobra.Command{
	Use:                   "delete -o SObject [filename]...",
	Short:                 "Delete page layout assignment",
	Long:                  "Delete page layout assignment for object from profiles",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteLayout(file, objectName)
		}
	},
}

func showLayout(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	layouts := p.GetLayouts(objectName)
	for _, l := range layouts {
		if l.RecordType != nil {
			fmt.Printf("%s (%s)\n", l.Layout, l.RecordType.Text)
		} else {
			fmt.Println(l.Layout)
		}
	}
}

func editLayout(file string, object, layout string, recordType string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	if recordType != "" {
		recordType = strings.TrimPrefix(recordType, object+".")
		p.SetObjectLayoutForRecordType(object, layout, recordType)
	} else {
		p.SetObjectLayout(object, layout)
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func deleteLayout(file string, object string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	p.DeleteObjectLayout(object)
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
