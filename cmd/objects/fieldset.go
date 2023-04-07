package objects

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects"
)

var (
	fieldSetName string
)

func init() {
	deleteFieldSetCmd.Flags().StringVarP(&fieldSetName, "fieldset", "s", "", "field set name")
	deleteFieldSetCmd.MarkFlagRequired("fieldset")

	FieldSetCmd.AddCommand(listFieldSetsCmd)
	FieldSetCmd.AddCommand(deleteFieldSetCmd)
}

var FieldSetCmd = &cobra.Command{
	Use:                   "fieldset",
	Short:                 "Manage object field set metadata",
	DisableFlagsInUseLine: true,
}

var deleteFieldSetCmd = &cobra.Command{
	Use:                   "delete -s FieldSet [filename]...",
	Short:                 "Delete object field set",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteFieldSet(file, fieldSetName)
		}
	},
}

var listFieldSetsCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List object field sets",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listFieldSets(file)
		}
	},
}

func listFieldSets(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	fieldSets := o.GetFieldSets()
	for _, f := range fieldSets {
		fmt.Printf("%s.%s\n", objectName, f.FullName)
	}
}

func deleteFieldSet(file string, fieldSetName string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	fieldSetName = strings.TrimPrefix(fieldSetName, objectName+".")
	err = o.DeleteFieldSet(fieldSetName)
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
