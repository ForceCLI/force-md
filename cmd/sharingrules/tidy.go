package sharingrules

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/sharingrules"
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
	Use:   "tidy [flags] [filename]...",
	Short: "Tidy sharing rules",
	Long: `
Tidy sharing rules metadata.

	The --wide and --ignore-errors flags can be used to help manage
	Sharing Rule metadata stored in a git repository.

	Configure clean and smudge git filters to use force-md:
	$ git config --local filter.salesforce-sharingrules.clean 'force-md sharingrules tidy --wide --ignore-errors -'
	$ git config --local filter.salesforce-sharingrules.smudge 'force-md sharingrules tidy --ignore-errors -'

	Update .gitattributes to use the salesforce-sharingrules filter:
	*.sharingRules-meta.xml filter=salesforce-sharingrules

	The --wide flag will cause the Sharing Rule metadata to be stored in a
	flattened format that makes it easier to resolve merge conflicts.  If a child
	of a fieldPermissions element changes, for example, the entire
	fieldPermissions element will show up as changed because it's stored on a single line.

	The smudge filter will cause the metadata to be unflattened so it's available
	in the normal "long" format in the working copy.

`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tidy(file)
		}
	},
}

func tidy(file string) {
	var p *sharingrules.SharingRules
	var err error
	if ignoreErrors {
		p = &sharingrules.SharingRules{}
		contents, err := metadata.ParseMetadataXmlIfPossible(p, file)
		if err != nil {
			log.Warn("parse failure. leaving content unchanged.")
			if _, err = os.Stdout.Write(contents); err != nil {
				log.Warn("failed to write contents")
			}
			return
		}
	} else {
		p, err = sharingrules.Open(file)
		if err != nil {
			log.Warn("parsing sharing rules failed: " + err.Error())
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
