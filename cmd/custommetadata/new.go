package custommetadata

import (
	"github.com/antonmedv/expr"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/custommetadata"
)

func init() {
	NewCmd.Flags().StringP("label", "l", "", "label")
	NewCmd.Flags().StringP("values", "v", "", "object describing values")
	NewCmd.MarkFlagRequired("label")
}

var NewCmd = &cobra.Command{
	Use:                   "new [filename]...",
	Short:                 "Create new custom metadata record",
	DisableFlagsInUseLine: false,
	Example: `
$ force-md custommetadata new src/customMetadata/My_Metadata.Example.md -l 'My Example' -v '{My_Field__c: "My Value", Default__c: true}'
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		label, _ := cmd.Flags().GetString("label")
		values, _ := cmd.Flags().GetString("values")
		for _, file := range args {
			createFile(file, label, values)
		}
	},
}

func createFile(file string, label, values string) {
	out, err := expr.Eval(values, nil)
	if err != nil {
		log.Fatal(err)
	}
	parsed, ok := out.(map[string]any)
	if !ok {
		log.Fatal("Could not parse values")
	}
	m := custommetadata.CustomMetadata{
		Xmlns:     "http://soap.sforce.com/2006/04/metadata",
		Xsi:       "http://www.w3.org/2001/XMLSchema-instance",
		Xsd:       "http://www.w3.org/2001/XMLSchema",
		Label:     label,
		Protected: FalseText,
	}
	for k, v := range parsed {
		m.AddValue(k, v)
	}
	m.Tidy()
	err = internal.WriteToFile(m, file)
	if err != nil {
		log.Warn("create failed: " + err.Error())
		return
	}
}
