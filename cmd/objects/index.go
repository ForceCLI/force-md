package objects

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/objects"
	"github.com/ForceCLI/force-md/metadata/objects/index"
)

var (
	indexesDir string
)

func init() {
	writeIndexesCmd.Flags().StringVarP(&indexesDir, "directory", "d", "", "directory where indexes should be output")
	writeIndexesCmd.MarkFlagRequired("directory")

	IndexCmd.AddCommand(writeIndexesCmd)
}

var IndexCmd = &cobra.Command{
	Use:                   "index",
	Short:                 "Manage big object index metadata",
	DisableFlagsInUseLine: true,
}

var writeIndexesCmd = &cobra.Command{
	Use:                   "write -d directory [filename]...",
	Short:                 "Split object indexes into separate files",
	Long:                  "Split object indexes into separate metadata files to match sfdx's source format",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			writeIndexes(file, indexesDir)
		}
	},
}

func writeIndexes(file string, indexesDir string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	indexes := o.GetIndexes()
	for _, i := range indexes {
		idx := index.Index{
			BigObjectIndex: i,
			Xmlns:          o.Xmlns,
		}
		err = internal.WriteToFile(idx, indexesDir+"/"+i.FullName+".index-meta.xml")
		if err != nil {
			log.Warn("write failed: " + err.Error())
			return
		}
	}
}
