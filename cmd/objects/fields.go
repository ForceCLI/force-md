package objects

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/objects"
)

var requiredOnly bool

func init() {
	listFieldsCmd.Flags().BoolVarP(&requiredOnly, "required", "r", false, "required fields only")

	FieldCmd.AddCommand(listFieldsCmd)
}

var FieldCmd = &cobra.Command{
	Use:   "field",
	Short: "Manage object field metadata",
}

var listFieldsCmd = &cobra.Command{
	Use:   "list",
	Short: "List object fields",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listFields(file)
		}
	},
}

var alwaysRequired map[string]bool = map[string]bool{
	"Name":    true,
	"OwnerId": true,
}

func listFields(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	var filters []objects.FieldFilter
	if requiredOnly {
		filters = append(filters, func(f objects.Field) bool {
			isRequired := alwaysRequired[f.FullName.Text] || (f.Required != nil && f.Required.Text == "true")
			isMasterDetail := f.Type != nil && f.Type.Text == "MasterDetail"
			return isRequired || isMasterDetail
		})
	}
	fields := o.GetFields(filters...)
	for _, f := range fields {
		fmt.Printf("%s.%s\n", objectName, f.FullName.Text)
	}
}
