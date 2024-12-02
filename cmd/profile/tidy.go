package profile

import (
	"bytes"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/profile"
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
	Short: "Tidy profile metadata",
	Long: `
Tidy profile metadata.

	The --wide and --ignore-errors flags can be used to help manage
	Profile metadata stored in a git repository.

	Configure clean and smudge git filters to use force-md:
	$ git config --local filter.salesforce-profile.clean 'force-md profile tidy --wide --ignore-errors -'
	$ git config --local filter.salesforce-profile.smudge 'force-md profile tidy --ignore-errors -'

	Update .gitattributes to use the salesforce-profile filter:
	*.profile-meta.xml filter=salesforce-profile

	The --wide flag will cause the Profile metadata to be stored in a
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
	p := &profile.Profile{}
	contents, err := metadata.ParseMetadataXmlIfPossible(p, file)
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
	var p *profile.Profile
	var err error
	if ignoreErrors {
		p = &profile.Profile{}
		contents, err := metadata.ParseMetadataXmlIfPossible(p, file)
		if err != nil {
			log.Warn("parse failure. leaving content unchanged.")
			if _, err = os.Stdout.Write(contents); err != nil {
				log.Warn("failed to write contents")
			}
			return
		}
	} else {
		p, err = profile.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
	}
	if wide {
		internal.MarshalWide = true
	}
	if err := general.Tidy(p, metadata.MetadataFilePath(file)); err != nil {
		log.Warn("tidying failed: " + err.Error())
	}
}
