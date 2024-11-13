package objects

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects"
	"github.com/ForceCLI/force-md/objects/field"
)

var warn bool

func init() {
	TidyCmd.Flags().BoolP("list", "l", false, "list files that need tidying")
	TidyCmd.Flags().Bool("fix-missing", false, "fix missing configuration (record type picklist options)")
	TidyCmd.Flags().BoolVar(&warn, "warn", false, "warn about possibly bad metadata (unassiged record type picklist options)")
}

type fixes struct {
	recordTypePicklistOptions bool
}

var TidyCmd = &cobra.Command{
	Use:   "tidy [flags] [filename]...",
	Short: "Tidy object metadata",
	Long: `
Tidy object metadata.

	The --fix-missing flag can be used to add missing object metadata.  This includes:
	* picklist fields missing from Record Types
`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		changes := false
		list, _ := cmd.Flags().GetBool("list")
		fixMissing, _ := cmd.Flags().GetBool("fix-missing")
		fix := fixes{
			recordTypePicklistOptions: fixMissing,
		}
		for _, file := range args {
			_, err := internal.Metadata.Open(file)
			if err != nil {
				return fmt.Errorf("invalid file %s: %w", file, err)
			}
		}

		items := internal.Metadata.Items(objects.NAME)
		if len(items) == 0 {
			log.Warn("No objects to tidy")
		}
		for _, item := range items {
			o := item.(*objects.CustomObject)
			if list {
				needsTidying := checkIfChanged(o, fix)
				changes = needsTidying || changes
			} else {
				tidy(o, fix)
			}
		}
		if changes {
			os.Exit(1)
		}
		return nil
	},
}

func checkUnassignedPicklistOptions(o *objects.CustomObject) {
	if len(o.RecordTypes) == 0 {
		return
	}
	filter := func(f field.Field) bool {
		if f.Type != nil && (strings.ToLower(f.Type.Text) == "picklist" || strings.ToLower(f.Type.Text) == "multiselectpicklist") {
			return true
		}
		return false
	}
	picklists := o.GetFields(filter)
PICKLIST:
	for _, field := range picklists {
		if strings.ToLower(o.Name()) == "account" && strings.HasPrefix(strings.ToLower(field.FullName), "person") && !strings.HasSuffix(strings.ToLower(field.FullName), "__c") {
			// Person Record Types are configured in the PersonAccount object
			continue
		}
		if field.ValueSet == nil {
			log.Debug(fmt.Sprintf("Value set not set for %s: checking not yet supported", field.FullName))
			continue PICKLIST
		}
		if field.ValueSet.ValueSetDefinition == nil && field.ValueSet.ValueSetName != nil {
			log.Debug(fmt.Sprintf("Values defined in value set %s: checking not yet supported", field.ValueSet.ValueSetName.Text))
			continue PICKLIST
		}
	VALUE:
		for _, value := range field.ValueSet.ValueSetDefinition.Value {
			if value.IsActive != nil && !value.IsActive.ToBool() {
				// Inactive values shouldn't be assigned to record types
				continue VALUE
			}
			foundValue := false
			for _, recordType := range o.RecordTypes {
				for _, recordTypePicklist := range recordType.PicklistValues {
					if strings.ToLower(recordTypePicklist.Picklist) != strings.ToLower(field.FullName) {
						continue
					}
					for _, recordTypePicklistValue := range recordTypePicklist.Values {
						v1, err := url.PathUnescape(recordTypePicklistValue.FullName)
						if err != nil {
							log.Warn(fmt.Sprintf("Could not decode value %s: %s", recordTypePicklistValue.FullName, err.Error()))
						}
						v2, err := url.PathUnescape(value.FullName)
						if err != nil {
							log.Warn(fmt.Sprintf("Could not decode value %s: %s", value.FullName, err.Error()))
						}
						if strings.ToLower(v1) == strings.ToLower(v2) {
							foundValue = true
							continue VALUE
						}

					}
				}
			}
			if !foundValue {
				log.Warn(fmt.Sprintf("%s.%s (%s): value %s not assigned to any record types", o.Name(), field.FullName, o.Path(), value.FullName))
			}
		}
	}
}

func addMissingRecordTypePicklistFields(o *objects.CustomObject) {
	filter := func(f field.Field) bool {
		if f.Type != nil && (strings.ToLower(f.Type.Text) == "picklist" || strings.ToLower(f.Type.Text) == "multiselectpicklist") {
			return true
		}
		return false
	}
	picklists := o.GetFields(filter)
	for _, field := range picklists {
		if strings.ToLower(o.Name()) == "account" && strings.HasPrefix(strings.ToLower(field.FullName), "person") && !strings.HasSuffix(strings.ToLower(field.FullName), "__c") {
			// Person Record Types are configured in the PersonAccount object
			continue
		}
		for _, recordType := range o.RecordTypes {
			hasPicklist := false
			for _, recordTypePicklist := range recordType.PicklistValues {
				if strings.ToLower(recordTypePicklist.Picklist) == strings.ToLower(field.FullName) {
					hasPicklist = true
					break
				}
			}
			if !hasPicklist {
				log.Warn(fmt.Sprintf("%s (%s): adding %s picklist field to record type %s", o.Name(), o.Path(), field.FullName, recordType.FullName))
				err := o.AddBlankPicklistOptionsToRecordType(field.FullName, recordType.FullName)
				if err != nil {
					log.Warn(err.Error())
					return
				}
			}
		}
	}
}

func checkIfChanged(o *objects.CustomObject, fix fixes) (changed bool) {
	contents := o.Contents()
	if fix.recordTypePicklistOptions {
		addMissingRecordTypePicklistFields(o)
	}
	if warn {
		checkUnassignedPicklistOptions(o)
	}
	o.Tidy()
	newContents, err := internal.Marshal(o)
	if err != nil {
		log.Warn("serializing failed: " + err.Error())
		return
	}
	if !bytes.Equal(contents, newContents) {
		fmt.Println(o.Path())
		return true
	}
	return false
}

func tidy(o *objects.CustomObject, fix fixes) {
	if warn {
		checkUnassignedPicklistOptions(o)
	}
	if fix.recordTypePicklistOptions {
		addMissingRecordTypePicklistFields(o)
	}
	if err := general.Tidy(o, o.Path()); err != nil {
		log.Warn("tidying failed: " + err.Error())
	}
}
