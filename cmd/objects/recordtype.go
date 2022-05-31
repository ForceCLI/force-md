package objects

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/objects"
)

func init() {
	RecordTypeCmd.AddCommand(listRecordTypesCmd)
	RecordTypeCmd.AddCommand(recordtypePicklistCmd)

	recordtypePicklistCmd.AddCommand(recordtypePicklistTableCmd)
}

var RecordTypeCmd = &cobra.Command{
	Use:                   "recordtype",
	Short:                 "Manage object record type metadata",
	DisableFlagsInUseLine: true,
}

var recordtypePicklistCmd = &cobra.Command{
	Use:                   "picklist",
	Short:                 "Manage record type picklist options",
	DisableFlagsInUseLine: true,
}

var recordtypePicklistTableCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "Display record type picklist options",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tableRecordTypePicklistOptions(file)
		}
	},
}

var listRecordTypesCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List object record types",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listRecordType(file)
		}
	},
}

func listRecordType(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	recordTypes := o.GetRecordTypes()
	for _, f := range recordTypes {
		fmt.Printf("%s.%s\n", objectName, f.FullName)
	}
}

func tableRecordTypePicklistOptions(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	recordTypes := o.GetRecordTypes()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Record Type", "Field", "Value", "Default"})
	table.SetAutoMergeCells(true)
	table.SetAutoMergeCellsByColumnIndex([]int{1})
	table.SetRowLine(true)
	for _, r := range recordTypes {
		for _, p := range r.PicklistValues {
			for _, v := range p.Values {
				if s, err := url.QueryUnescape(v.FullName); err == nil {
					table.Append([]string{r.FullName, objectName + "." + p.Picklist, s, v.Default.Text})
				} else {
					panic(err.Error())
				}
			}
		}
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}
