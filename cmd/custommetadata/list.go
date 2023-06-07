package custommetadata

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/antonmedv/expr"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/custommetadata"
)

func init() {
	TableCmd.Flags().StringP("filter", "f", "true", "expr boolean expression to filter records")
	ListCmd.Flags().StringP("filter", "f", "true", "expr boolean expression to filter records")
}

var ListCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List custom metadata in a table",
	Args:  cobra.MinimumNArgs(1),
	Example: `
$ force-md custommetadata list -f 'record.dlrs__CalculationMode__c != "Realtime"' src/customMetadata/dlrs__LookupRollupSummary2.*
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
$ force-md custommetadata table -f 'record.dlrs__CalculationMode__c != "Realtime"' src/customMetadata/dlrs__LookupRollupSummary2.*
`,
	Run: func(cmd *cobra.Command, args []string) {
		filter, _ := cmd.Flags().GetString("filter")
		tableCustomMetadata(args, filter)
	},
}

func tableCustomMetadata(files []string, filter string) {
	env := map[string]interface{}{
		"record": make(map[string]string),
	}
	program, err := expr.Compile(filter, expr.Env(env))
	if err != nil {
		log.Fatalln("Invalid expression:", err)
	}
	var fieldNames []string
	allFields := make(map[string]bool)
	type record struct {
		label  string
		fields map[string]string
	}
	var records []record
	for _, file := range files {
		m, err := custommetadata.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
		// file := strings.TrimSuffix(path.Base(file), ".md")
		fields := make(map[string]string)
		for _, v := range m.Values {
			if _, ok := allFields[v.Field]; !ok {
				fieldNames = append(fieldNames, v.Field)
				allFields[v.Field] = true
			}
			fields[v.Field] = v.Value.Text
		}
		env := map[string]interface{}{
			"record": fields,
		}
		out, err := expr.Run(program, env)
		if err != nil {
			panic(err)
		}
		include, _ := out.(bool)
		if include {
			records = append(records, record{m.Label, fields})
		}

	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader(append([]string{"Label"}, fieldNames...))
	table.SetRowLine(true)
	for _, r := range records {
		row := []string{r.label}
		for _, f := range fieldNames {
			row = append(row, r.fields[f])
		}
		table.Append(row)
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}

func listCustomMetadata(files []string, filter string) {
	env := map[string]interface{}{
		"record": make(map[string]string),
	}
	program, err := expr.Compile(filter, expr.Env(env))
	if err != nil {
		log.Fatalln("Invalid expression:", err)
	}
	var fieldNames []string
	allFields := make(map[string]bool)
	type record struct {
		label  string
		fields map[string]string
	}
	for _, file := range files {
		m, err := custommetadata.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
		// file := strings.TrimSuffix(path.Base(file), ".md")
		fields := make(map[string]string)
		for _, v := range m.Values {
			if _, ok := allFields[v.Field]; !ok {
				fieldNames = append(fieldNames, v.Field)
				allFields[v.Field] = true
			}
			fields[v.Field] = v.Value.Text
		}
		env := map[string]interface{}{
			"record": fields,
		}
		out, err := expr.Run(program, env)
		if err != nil {
			panic(err)
		}
		include, _ := out.(bool)
		if include {
			fmt.Println(strings.TrimSuffix(filepath.Base(file), filepath.Ext(file)))
		}
	}
}
