package objects

import (
	"bytes"
	"fmt"
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
		for _, file := range args {
			list, _ := cmd.Flags().GetBool("list")
			fixMissing, _ := cmd.Flags().GetBool("fix-missing")
			fix := fixes{
				recordTypePicklistOptions: fixMissing,
			}
			if list {
				checkIfChanged(file, fix)
			} else {
				tidy(file, fix)
			}
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
		for _, recordType := range o.RecordTypes {
			hasPicklist := false
			for _, recordTypePicklist := range recordType.PicklistValues {
				if strings.ToLower(recordTypePicklist.Picklist) == strings.ToLower(field.FullName) {
					hasPicklist = true
					break
				}
			}
			if !hasPicklist {
				log.Warn(fmt.Sprintf("adding %s to record type %s", field.FullName, recordType.FullName))
				err := o.AddBlankPicklistOptionsToRecordType(field.FullName, recordType.FullName)
				if err != nil {
					log.Warn(err.Error())
					return
				}
			}
		}
	}
}

func checkIfChanged(file string, fix fixes) {
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
	}
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
