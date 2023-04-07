package application

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"

	"github.com/ForceCLI/force-md/application"
	"github.com/ForceCLI/force-md/internal"
)

var action Action
var formFactor FormFactor
var profile string
var pageObject string
var recordType string
var content string

type Action enumflag.Flag
type FormFactor enumflag.Flag

const (
	NoneAction Action = iota
	View
	Tab
)

const (
	NoneFormFactor FormFactor = iota
	Large
	Small
)

var ActionIds = map[Action][]string{
	NoneAction: {"None"},
	View:       {"View"},
	Tab:        {"Tab"},
}

var FormFactorIds = map[FormFactor][]string{
	NoneFormFactor: {"None"},
	Large:          {"Large"},
	Small:          {"Small"},
}

func init() {
	tableActionCmd.Flags().VarP(enumflag.New(&action, "action", ActionIds, enumflag.EnumCaseInsensitive),
		"action", "a", "action; can be 'Tab' or 'View'")

	tableActionCmd.Flags().VarP(enumflag.New(&formFactor, "formfactor", FormFactorIds, enumflag.EnumCaseInsensitive),
		"formfactor", "f", "form factor; can be 'Large' or 'Small'")

	tableActionCmd.Flags().StringVarP(&profile, "profile", "p", "", "profile name")
	tableActionCmd.Flags().StringVarP(&pageObject, "object", "o", "", "sobject or page name")
	tableActionCmd.Flags().StringVarP(&recordType, "recordType", "r", "", "record type")
	tableActionCmd.Flags().StringVarP(&content, "content", "c", "", "content")

	deleteActionCmd.Flags().StringVarP(&profile, "profile", "p", "", "profile name")
	deleteActionCmd.Flags().StringVarP(&pageObject, "object", "o", "", "sobject or page name")
	deleteActionCmd.Flags().StringVarP(&content, "content", "c", "", "content")
	deleteActionCmd.Flags().VarP(enumflag.New(&formFactor, "formfactor", FormFactorIds, enumflag.EnumCaseInsensitive),
		"formfactor", "f", "form factor; can be 'Large' or 'Small'")

	ActionCmd.AddCommand(tableActionCmd)
	ActionCmd.AddCommand(deleteActionCmd)
}

var ActionCmd = &cobra.Command{
	Use:   "action",
	Short: "Manage Profile Action Overrides ",
}

var tableActionCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List Profile Action Overrides in a table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tableProfileActionOverrides(file)
		}
	},
}

var deleteActionCmd = &cobra.Command{
	Use:   "delete [flags] [filename]...",
	Short: "Delete action overrides",
	Long:  "Delete action overrides from applications",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteActionOverride(file)
		}
	},
}

func deleteActionOverride(file string) {
	o, err := application.Open(file)
	if err != nil {
		log.Warn("parsing application failed: " + err.Error())
		return
	}
	var filters []application.ProfileActionOverrideFilter
	if content != "" {
		filters = append(filters, func(a application.ProfileActionOverride) bool {
			return strings.ToLower(a.Content) == strings.ToLower(content)
		})
	}
	switch formFactor {
	case Large, Small:
		filters = append(filters, func(a application.ProfileActionOverride) bool {
			return a.FormFactor == FormFactorIds[formFactor][0]
		})
	}
	if profile != "" {
		filters = append(filters, func(a application.ProfileActionOverride) bool {
			return strings.ToLower(a.Profile) == strings.ToLower(profile)
		})
	}
	if pageObject != "" {
		filters = append(filters, func(a application.ProfileActionOverride) bool {
			return strings.ToLower(a.PageOrSobjectType) == strings.ToLower(pageObject)
		})
	}
	err = o.DeleteActionOverrides(filters...)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(o, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func tableProfileActionOverrides(file string) {
	w, err := application.Open(file)
	if err != nil {
		log.Warn("parsing applications failed: " + err.Error())
		return
	}
	applicationName := strings.TrimSuffix(path.Base(file), ".app")
	var filters []application.ProfileActionOverrideFilter
	switch action {
	case Tab, View:
		filters = append(filters, func(a application.ProfileActionOverride) bool {
			return a.ActionName == ActionIds[action][0]
		})
	}
	switch formFactor {
	case Large, Small:
		filters = append(filters, func(a application.ProfileActionOverride) bool {
			return a.FormFactor == FormFactorIds[formFactor][0]
		})
	}
	if profile != "" {
		filters = append(filters, func(a application.ProfileActionOverride) bool {
			return strings.ToLower(a.Profile) == strings.ToLower(profile)
		})
	}
	if pageObject != "" {
		filters = append(filters, func(a application.ProfileActionOverride) bool {
			return strings.ToLower(a.PageOrSobjectType) == strings.ToLower(pageObject)
		})
	}
	if recordType != "" {
		if !strings.Contains(recordType, ".") && pageObject != "" {
			recordType = pageObject + "." + recordType
		}
		filters = append(filters, func(a application.ProfileActionOverride) bool {
			return a.RecordType != nil && strings.ToLower(*a.RecordType) == strings.ToLower(recordType)
		})
	}
	if content != "" {
		filters = append(filters, func(a application.ProfileActionOverride) bool {
			return strings.ToLower(a.Content) == strings.ToLower(content)
		})
	}
	actions := w.GetProfileActionOverrides(filters...)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Application", "Page/Object", "Record Type", "Profile", "Action", "Form Factor", "Lightning Page"})
	table.SetAutoMergeCells(true)
	table.SetAutoMergeCellsByColumnIndex([]int{1, 2})
	table.SetRowLine(true)
	for _, r := range actions {
		recordType := ""
		if r.RecordType != nil {
			recordType = *r.RecordType
		}
		table.Append([]string{applicationName, r.PageOrSobjectType, recordType, r.Profile, r.ActionName, r.FormFactor, r.Content})
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}
