package globalvalueset

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/globalvalueset"
)

var (
	active bool
	labels bool
)

func init() {
	ListCmd.Flags().BoolVarP(&active, "active", "a", false, "active only")
	ListCmd.Flags().BoolVarP(&labels, "labels", "l", false, "show labels instead of API value")
}

var ListCmd = &cobra.Command{
	Use:   "list [filename]...",
	Short: "List global value set values",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listValues(file)
		}
	},
}

func listValues(file string) {
	a, err := globalvalueset.Open(file)
	if err != nil {
		log.Warn("parsing global value set failed: " + err.Error())
		return
	}
	var filters []globalvalueset.ValueFilter
	if active {
		filters = append(filters, func(f globalvalueset.CustomValue) bool { return f.IsActive == nil || f.IsActive.ToBool() })
	}
	fields := a.GetValues(filters...)
	for _, f := range fields {
		if labels {
			fmt.Println(f.Label)
		} else {
			fmt.Println(f.FullName)
		}
	}
}
