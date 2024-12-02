package objects

import (
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/objects"
)

var actionType Type
var formFactor FormFactor

type Type enumflag.Flag
type FormFactor enumflag.Flag

const (
	NoneType Type = iota
	Default
	Flexipage
	LightningComponent
	Scontrol
	Standard
	Visualforce
)

const (
	NoneFormFactor FormFactor = iota
	Large
	Small
)

var TypeIds = map[Type][]string{
	NoneType:           {"None"},
	Default:            {"Default"},
	Flexipage:          {"FlexiPage"},
	LightningComponent: {"LightningComponent"},
	Scontrol:           {"Scontrol"},
	Standard:           {"Standard"},
	Visualforce:        {"Visualforce"},
}

var FormFactorIds = map[FormFactor][]string{
	NoneFormFactor: {"None"},
	Large:          {"Large"},
	Small:          {"Small"},
}

func init() {
	tableActionCmd.Flags().VarP(enumflag.New(&actionType, "type", TypeIds, enumflag.EnumCaseInsensitive),
		"type", "t", "type; can be 'Default', 'FlexiPage', 'LightningComponent', 'Scontrol', 'Standard', or 'Visualforce'")

	tableActionCmd.Flags().VarP(enumflag.New(&formFactor, "formfactor", FormFactorIds, enumflag.EnumCaseInsensitive),
		"formfactor", "f", "form factor; can be 'Large' or 'Small'")

	ActionCmd.AddCommand(tableActionCmd)
}

var ActionCmd = &cobra.Command{
	Use:   "action",
	Short: "Manage Action Overrides ",
}

var tableActionCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List Action Overrides in a table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tableActionOverrides(file)
		}
	},
}

func tableActionOverrides(file string) {
	w, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing applications failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
	var filters []objects.ActionOverrideFilter
	switch actionType {
	case Default, Flexipage, LightningComponent, Scontrol, Standard, Visualforce:
		filters = append(filters, func(a objects.ActionOverride) bool {
			return strings.ToLower(a.Type) == strings.ToLower(TypeIds[actionType][0])
		})
	}
	switch formFactor {
	case Large, Small:
		filters = append(filters, func(a objects.ActionOverride) bool {
			return a.FormFactor != nil && strings.ToLower(*a.FormFactor) == strings.ToLower(FormFactorIds[formFactor][0])
		})
	}
	actions := w.GetActionOverrides(filters...)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Object", "Action", "Form Factor", "Type", "Page/Component"})
	table.SetAutoMergeCells(true)
	table.SetAutoMergeCellsByColumnIndex([]int{1, 2, 3})
	table.SetRowLine(true)
	for _, r := range actions {
		var formFactor, content string
		if r.FormFactor != nil {
			formFactor = *r.FormFactor
		}
		if r.Content != nil {
			content = *r.Content
		}
		table.Append([]string{objectName, r.ActionName, formFactor, r.Type, content})
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}
