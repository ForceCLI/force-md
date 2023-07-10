package application

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/application"
	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

var (
	wide         bool
	ignoreErrors bool
)

func init() {
	TidyCmd.Flags().BoolVarP(&wide, "wide", "w", false, "flatten into wide format")
	TidyCmd.Flags().BoolVarP(&ignoreErrors, "ignore-errors", "i", false, "ignore errors")
}

var TidyCmd = &cobra.Command{
	Use:   "tidy [filename]...",
	Short: "Tidy custom application",
	Long: `
Tidy custom application metadata.

	The --wide and --ignore-errors flags can be used to help manage
	CustomApplication metadata stored in a git repository.

	Configure clean and smudge git filters to use force-md:
	$ git config --local filter.salesforce-application.clean 'force-md application tidy --wide --ignore-errors -'
	$ git config --local filter.salesforce-application.smudge 'force-md application tidy --ignore-errors -'

	Update .gitattributes to use the salesforce-application filter:
	*.app-meta.xml filter=salesforce-application

	The --wide flag will cause the CustomApplication metadata to be stored in a
	flattened format that makes it easier to resolve merge conflicts.  If a child
	of a profileActionOverrides element changes, for example, the entire
	profileActionOverrides element will show up as changed because it's stored on a single line.

	The smudge filter will cause the metadata to be unflattened so it's available
	in the normal "long" format in the working copy.

`,
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tidy(file)
		}
	},
}

func tidy(file string) {
	var p *application.CustomApplication
	var err error
	if ignoreErrors {
		p = &application.CustomApplication{}
		contents, err := internal.ParseMetadataXmlIfPossible(p, file)
		if err != nil {
			log.Warn("parse failure. leaving content unchanged.")
			if _, err = os.Stdout.Write(contents); err != nil {
				log.Warn("failed to write contents")
			}
		}
	} else {
		p, err = application.Open(file)
		if err != nil {
			log.Warn("parsing application failed: " + err.Error())
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
