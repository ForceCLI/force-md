package custommetadata

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/custommetadata"
)

func init() {
	TableCmd.Flags().StringP("filter", "f", "true", "expr boolean expression to filter records")
	TableCmd.Flags().StringArrayP("fields", "i", []string{}, "field(s) to include in table")
	ListCmd.Flags().StringP("filter", "f", "true", "expr boolean expression to filter records")
}

var ListCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List custom metadata",
	Args:  cobra.MinimumNArgs(1),
	Example: `
$ force-md custommetadata list -f 'dlrs__CalculationMode__c != "Realtime"' src/customMetadata/dlrs__LookupRollupSummary2.*
`,
	Run: func(cmd *cobra.Command, args []string) {
		filter, _ := cmd.Flags().GetString("filter")
		listCustomMetadata(args, filter)
	},
}

var TableCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List custom metadata in a table",
	Args:  cobra.MinimumNArgs(1),
	Example: `
$ force-md custommetadata table -f 'dlrs__CalculationMode__c != "Realtime"' src/customMetadata/dlrs__LookupRollupSummary2.*
`,
	Run: func(cmd *cobra.Command, args []string) {
		filter, _ := cmd.Flags().GetString("filter")
		fields, _ := cmd.Flags().GetStringArray("fields")
		fieldMap := make(map[string]struct{})
		for _, f := range fields {
			fieldMap[strings.ToLower(f)] = struct{}{}
		}
		tableCustomMetadata(args, filter, fieldMap)
	},
}

func tableCustomMetadata(files []string, filter string, wantedFields map[string]struct{}) {
	var program *vm.Program
	fieldNames := []string{"Label"}
	customFields := make(map[string]struct{})
	var records []map[string]string
	for _, file := range files {
		m, err := custommetadata.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
		fields := make(map[string]string)
		fields["Label"] = m.Label
		for _, v := range m.Values {
			if _, w := wantedFields[strings.ToLower(v.Field)]; !w && len(wantedFields) > 0 {
				// Skip fields not wanted in output
				continue
			}
			if _, ok := customFields[v.Field]; !ok {
				fieldNames = append(fieldNames, v.Field)
				customFields[v.Field] = struct{}{}
			}
			fields[v.Field] = v.Value.Text
		}
		if program == nil {
			program, err = expr.Compile(filter, expr.Env(fields))
			if err != nil {
				log.Fatalln("Invalid expression:", err)
			}
		}
		out, err := expr.Run(program, fields)
		if err != nil {
			panic(err)
		}
		include, _ := out.(bool)
		if include {
			records = append(records, fields)
		}

	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader(fieldNames)
	table.SetRowLine(true)
	for _, r := range records {
		row := []string{}
		for _, f := range fieldNames {
			row = append(row, r[f])
		}
		table.Append(row)
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}

func listCustomMetadata(files []string, filter string) {
	var program *vm.Program
	for _, file := range files {
		m, err := custommetadata.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
		fields := make(map[string]string)
		for _, v := range m.Values {
			fields[v.Field] = v.Value.Text
		}
		if program == nil {
			program, err = expr.Compile(filter, expr.Env(fields))
			if err != nil {
				log.Fatalln("Invalid expression:", err)
			}
		}
		out, err := expr.Run(program, fields)
		if err != nil {
			panic(err)
		}
		include, _ := out.(bool)
		if include {
			fmt.Println(strings.TrimSuffix(filepath.Base(file), filepath.Ext(file)))
		}
	}
}
