package permissionset

import (
	"bytes"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/permissionset"
)

var (
	wide         bool
	ignoreErrors bool
)

func init() {
	TidyCmd.Flags().BoolVarP(&wide, "wide", "w", false, "flatten into wide format")
	TidyCmd.Flags().BoolP("list", "l", false, "list files that need tidying")
	TidyCmd.Flags().BoolVarP(&ignoreErrors, "ignore-errors", "i", false, "ignore errors")
}

var TidyCmd = &cobra.Command{
	Use:   "tidy [flags] [filename]...",
	Short: "Tidy Permission Set metadata",
	Long: `
Tidy permission set metadata.

	The --wide and --ignore-errors flags can be used to help manage
	Permission Set metadata stored in a git repository.

	Configure clean and smudge git filters to use force-md:
	$ git config --local filter.salesforce-permissionset.clean 'force-md permissionset tidy --wide --ignore-errors -'
	$ git config --local filter.salesforce-permissionset.smudge 'force-md permissionset tidy --ignore-errors -'

	Update .gitattributes to use the salesforce-permissionset filter:
	*.permissionset-meta.xml filter=salesforce-permissionset

	The --wide flag will cause the Permission Set metadata to be stored in a
	flattened format that makes it easier to resolve merge conflicts.  If a child
	of a fieldPermissions element changes, for example, the entire
	fieldPermissions element will show up as changed because it's stored on a single line.

	The smudge filter will cause the metadata to be unflattened so it's available
	in the normal "long" format in the working copy.

`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			list, _ := cmd.Flags().GetBool("list")
			if list {
				checkIfChanged(file)
			} else {
				tidy(file)
			}
		}
	},
}

func checkIfChanged(file string) {
	p := &permissionset.PermissionSet{}
	contents, err := internal.ParseMetadataXmlIfPossible(p, file)
	if err != nil {
		log.Warn("parse failure:" + err.Error())
		return
	}
	p.Tidy()
	newContents, err := internal.Marshal(p)
	if err != nil {
		log.Warn("serializing failed: " + err.Error())
		return
	}
	if !bytes.Equal(contents, newContents) {
		fmt.Println(file)
	}
}

func tidy(file string) {
	var p *permissionset.PermissionSet
	var err error
	if ignoreErrors {
		p = &permissionset.PermissionSet{}
		contents, err := internal.ParseMetadataXmlIfPossible(p, file)
		if err != nil {
			log.Warn("parse failure. leaving content unchanged.")
			if _, err = os.Stdout.Write(contents); err != nil {
				log.Warn("failed to write contents")
			}
			return
		}
	} else {
		p, err = permissionset.Open(file)
		if err != nil {
			log.Warn("parsing permission set failed: " + err.Error())
			return
		}
	}
	if wide {
		internal.MarshalWide = true
	}
	if err := general.Tidy(p, file); err != nil {
		log.Warn("tidying failed: " + err.Error())
	}
}
