package objects

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects"
	"github.com/ForceCLI/force-md/objects/field"
)

func init() {
	TidyCmd.Flags().BoolP("list", "l", false, "list files that need tidying")
	TidyCmd.Flags().Bool("fix-missing", false, "fix missing configuration (record type picklist options)")
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
	Run: func(cmd *cobra.Command, args []string) {
		changes := false
		for _, file := range args {
			list, _ := cmd.Flags().GetBool("list")
			fixMissing, _ := cmd.Flags().GetBool("fix-missing")
			fix := fixes{
				recordTypePicklistOptions: fixMissing,
			}
			if list {
				needsTidying := checkIfChanged(file, fix)
				changes = needsTidying || changes
			} else {
				tidy(file, fix)
			}
		}
		if changes {
			os.Exit(1)
		}
	},
}

func addMissingRecordTypePicklistFields(o *objects.CustomObject) {
	filter := func(f field.Field) bool {
		if f.Type != nil && strings.ToLower(f.Type.Text) == "picklist" {
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

func checkIfChanged(file string, fix fixes) (changed bool) {
	o := &objects.CustomObject{}
	contents, err := internal.ParseMetadataXmlIfPossible(o, file)
	if err != nil {
		log.Warn("parse failure:" + err.Error())
		return
	}
	if fix.recordTypePicklistOptions {
		addMissingRecordTypePicklistFields(o)
	}
	o.Tidy()
	newContents, err := internal.Marshal(o)
	if err != nil {
		log.Warn("serializing failed: " + err.Error())
		return
	}
	if !bytes.Equal(contents, newContents) {
		fmt.Println(file)
		return true
	}
	return false
}

func tidy(file string, fix fixes) {
	p, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	if fix.recordTypePicklistOptions {
		addMissingRecordTypePicklistFields(p)
	}
	if err := general.Tidy(p, file); err != nil {
		log.Warn("tidying failed: " + err.Error())
	}
}
