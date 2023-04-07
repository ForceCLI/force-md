package standardvalueset

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/standardvalueset"
)

var ListCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List values",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listValues(file)
		}
	},
}

func listValues(file string) {
	w, err := standardvalueset.Open(file)
	if err != nil {
		log.Warn("parsing value set failed: " + err.Error())
		return
	}
	valueSet := strings.TrimSuffix(path.Base(file), ".standardValueSet")
	rules := w.GetValues()
	for _, r := range rules {
		fmt.Printf("%s: %s\n", valueSet, r.FullName.Text)
	}
}
