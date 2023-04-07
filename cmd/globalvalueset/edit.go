package globalvalueset

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/globalvalueset"
	"github.com/ForceCLI/force-md/internal"
)

var ()

func init() {
	EditCmd.Flags().BoolP("sorted", "s", false, "sort values")
	EditCmd.Flags().BoolP("no-sorted", "S", false, "do not sort values")
}

var EditCmd = &cobra.Command{
	Use:   "edit [flags] [filename]...",
	Short: "Edit global value set",
	Long:  "Edit global value set",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		updates := setAttributes(cmd)
		for _, file := range args {
			updateGlobalValueSet(file, updates)
		}
	},
}

func setAttributes(cmd *cobra.Command) globalvalueset.GlobalValueSet {
	globalValueSet := globalvalueset.GlobalValueSet{}
	globalValueSet.Sorted = BooleanTextValue(cmd, "sorted")
	return globalValueSet
}

func updateGlobalValueSet(file string, updates globalvalueset.GlobalValueSet) {
	a, err := globalvalueset.Open(file)
	if err != nil {
		log.Warn("parsing global value set failed: " + err.Error())
		return
	}
	a.UpdateGlobalValueSet(updates)
	a.Tidy()
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(a, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
