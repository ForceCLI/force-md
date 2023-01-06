package custommetadata

import (
	"os"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/custommetadata"
)

func init() {
}

var TableCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List custom metadata in a table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tableCustomMetadata(args)
	},
}

func tableCustomMetadata(files []string) {
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
		records = append(records, record{m.Label, fields})
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
