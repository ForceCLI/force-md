package labels

import (
	"os"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/metadata/labels"
)

func init() {
}

var TableCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List custom labels in a table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tableCustomLabels(args)
	},
}

func tableCustomLabels(files []string) {
	var rows labels.CustomLabelList
	for _, file := range files {
		m, err := labels.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
		for _, v := range m.GetLabels() {
			rows = append(rows, v)
		}
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{"Name", "Categories", "Language", "Protected", "Description", "Value"})
	table.SetRowLine(true)
	for _, r := range rows {
		row := []string{r.FullName, r.Categories, r.Language, r.Protected, r.ShortDescription, r.Value}
		table.Append(row)
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}
