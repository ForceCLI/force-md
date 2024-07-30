package custommetadata

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/custommetadata"
	"github.com/ForceCLI/force-md/internal"
)

var (
	field string
	value string
	label string
)

func init() {
	EditCmd.Flags().StringVarP(&field, "field", "f", "", "field")
	EditCmd.Flags().StringVarP(&value, "value", "v", "", "value")

	EditCmd.Flags().StringVarP(&label, "label", "l", "", "label")

	EditCmd.MarkFlagsRequiredTogether("field", "value")
}

var EditCmd = &cobra.Command{
	Use:   "edit [filename]...",
	Short: "Edit custom metadata",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			if field != "" {
				editCustomMetadataValue(file, field, value)
			}
			if label != "" {
				editCustomMetadataLabel(file, label)
			}
		}
	},
}

func editCustomMetadataValue(file string, field string, value string) {
	p, err := custommetadata.Open(file)
	if err != nil {
		log.Warn("parsing custom metadata failed: " + err.Error())
		return
	}
	err = p.UpdateFieldValue(field, value)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func editCustomMetadataLabel(file string, label string) {
	p, err := custommetadata.Open(file)
	if err != nil {
		log.Warn("parsing custom metadata failed: " + err.Error())
		return
	}
	// Create new CustomMetadata element to deal with bugs
	// marshaling/unmarshaling XML with namespaces
	m := custommetadata.CustomMetadata{
		Xmlns:     "http://soap.sforce.com/2006/04/metadata",
		Xsi:       "http://www.w3.org/2001/XMLSchema-instance",
		Xsd:       "http://www.w3.org/2001/XMLSchema",
		Label:     label,
		Protected: p.Protected,
	}
	for _, v := range p.Values {
		m.AddValue(v.Field, v.Value.Text)
	}
	m.Tidy()
	err = internal.WriteToFile(m, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
