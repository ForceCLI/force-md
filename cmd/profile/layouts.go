package profile

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/profile"
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
	deleteLayoutCmd.Flags().StringVarP(&layoutName, "layout", "l", "", "layout name")
	deleteLayoutCmd.Flags().StringVarP(&recordType, "recordtype", "r", "", "record type")
	deleteLayoutCmd.MarkFlagRequired("object")

	tableLayoutsCmd.Flags().StringVarP(&objectName, "object", "o", "", "object name")
	tableLayoutsCmd.Flags().StringVarP(&layoutName, "layout", "l", "", "layout name")
	tableLayoutsCmd.Flags().StringVarP(&recordType, "recordtype", "r", "", "record type")

	LayoutCmd.AddCommand(showLayoutCmd)
	LayoutCmd.AddCommand(editLayoutCmd)
	LayoutCmd.AddCommand(deleteLayoutCmd)
	LayoutCmd.AddCommand(tableLayoutsCmd)
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
	Use:   "delete -o SObject [filename]...",
	Short: "Delete page layout assignment",
	Long:  "Delete page layout assignment for object from profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteLayout(file)
		}
	},
}

var tableLayoutsCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List Page Layouts in a table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tableLayouts(args)
	},
}

func showLayout(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	objectFilter := func(f profile.LayoutAssignment) bool {
		pieces := strings.SplitN(f.Layout, "-", 2)
		if len(pieces) != 2 {
			return false
		}
		return strings.ToLower(pieces[0]) == strings.ToLower(objectName)
	}
	layouts := p.GetLayouts(objectFilter)
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
		layout = strings.TrimPrefix(layout, object+"-")
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

func deleteLayout(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	var filters []profile.LayoutFilter
	if layoutName != "" {
		layoutName = strings.TrimPrefix(layoutName, objectName+"-")
		fullLayoutName := strings.ToLower(objectName + "-" + layoutName)
		filters = append(filters, func(l profile.LayoutAssignment) bool {
			return strings.ToLower(l.Layout) == fullLayoutName
		})
	}
	if recordType != "" {
		recordType = strings.TrimPrefix(recordType, objectName+".")
		fullRecordTypeName := strings.ToLower(objectName + "." + recordType)
		filters = append(filters, func(l profile.LayoutAssignment) bool {
			return l.RecordType != nil && strings.ToLower(l.RecordType.Text) == fullRecordTypeName
		})
	}
	err = p.DeleteObjectLayout(objectName, filters...)
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

func tableLayouts(files []string) {
	var filters []profile.LayoutFilter
	if objectName != "" {
		filters = append(filters, func(f profile.LayoutAssignment) bool {
			pieces := strings.SplitN(f.Layout, "-", 2)
			if len(pieces) != 2 {
				return false
			}
			return strings.ToLower(pieces[0]) == strings.ToLower(objectName)
		})
	}
	if recordType != "" {
		fullRecordTypeName := strings.ToLower(recordType)
		if objectName != "" {
			recordType = strings.TrimPrefix(strings.ToLower(recordType), strings.ToLower(objectName)+".")
			fullRecordTypeName = objectName + "." + recordType
		}
		filters = append(filters, func(f profile.LayoutAssignment) bool {
			if f.RecordType == nil {
				return false
			}
			return strings.ToLower(f.RecordType.Text) == strings.ToLower(fullRecordTypeName)
		})
	}
	if layoutName != "" {
		fullLayoutName := strings.ToLower(layoutName)
		if objectName != "" {
			layoutName = strings.TrimPrefix(strings.ToLower(layoutName), strings.ToLower(objectName)+"-")
			fullLayoutName = strings.ToLower(objectName) + "-" + layoutName
		}
		filters = append(filters, func(f profile.LayoutAssignment) bool {
			return strings.ToLower(f.Layout) == strings.ToLower(fullLayoutName)
		})
	}
	type profileLayouts struct {
		layouts profile.LayoutAssignmentList
		profile string
	}
	var layouts []profileLayouts
	for _, file := range files {
		p, err := profile.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
		profileName := internal.TrimSuffixToEnd(path.Base(file), ".profile")
		layouts = append(layouts, profileLayouts{layouts: p.GetLayouts(filters...), profile: profileName})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Profile", "Object", "RecordType", "Layout"})
	table.SetRowLine(true)
	for _, vis := range layouts {
		for _, o := range vis.layouts {
			pieces := strings.SplitN(o.Layout, "-", 2)
			if len(pieces) != 2 {
				log.Warn("Unexpected Layout: " + o.Layout)
				continue
			}
			obj := pieces[0]
			layout := pieces[1]
			rt := ""
			if o.RecordType != nil {
				rt = o.RecordType.Text
			}
			table.Append([]string{vis.profile, obj, rt, layout})
		}
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}
