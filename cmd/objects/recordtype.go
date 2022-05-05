package objects

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/objects"
)

func init() {
	RecordTypeCmd.AddCommand(listRecordTypesCmd)
}

var RecordTypeCmd = &cobra.Command{
	Use:                   "recordtype",
	Short:                 "Manage object record type metadata",
	DisableFlagsInUseLine: true,
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
