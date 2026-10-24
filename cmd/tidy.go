package cmd

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	cmdrepo "github.com/ForceCLI/force-md/cmd/repo"
	"github.com/ForceCLI/force-md/repo"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	tidyCmd.Flags().BoolP("list", "l", false, "list files that need tidying")
	tidyCmd.Flags().BoolP("recursive", "r", false, "tidy metadata files in directories and their subdirectories")

	RootCmd.AddCommand(tidyCmd)
}

var tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "Tidy Metadata",
	Example: `
$ force-md tidy sfdx/main/default/objects/*/{fields,validationRules}/* sfdx/main/default/flows/*

$ force-md tidy src/objects/*

$ force-md tidy -r force-app/main/default
`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		recursive, _ := cmd.Flags().GetBool("recursive")
		if err := openTidyArgs(cmdrepo.Metadata, args, recursive); err != nil {
			return err
		}
		changes := false
		list, _ := cmd.Flags().GetBool("list")
		for _, t := range cmdrepo.Metadata.Types() {
			for _, m := range cmdrepo.Metadata.Items(t) {
				file := m.GetMetadataInfo().Path()
				o, ok := m.(general.Tidyable)
				if !ok {
					if recursive {
						log.Debugf("file %s of type %s is not tidyable", file, m.Type())
					} else {
						log.Warnf("file %s of type %s is not tidyable", file, m.Type())
					}
					continue
				}
				if list {
					orig := m.GetMetadataInfo().Contents()
					needsTidying := checkIfChanged(o, orig)
					if needsTidying {
						fmt.Println(file)
					}
					changes = needsTidying || changes
				} else {
					if err := general.Tidy(o, file); err != nil {
						log.Warnf("tidying failed: %s", err.Error())
					}
				}
			}
		}
		if changes {
			os.Exit(1)
		}
		return nil
	},
}

// openTidyArgs opens each file in args into r.  With recursive set, a
// directory in args is walked and every metadata file below it is opened.
// The walk skips files that are not metadata and hidden directories such as
// .git, and logs a warning for a metadata file that fails to open instead of
// stopping.
func openTidyArgs(r *repo.Repo, args []string, recursive bool) error {
	for _, arg := range args {
		if recursive {
			info, err := os.Stat(arg)
			if err != nil {
				return fmt.Errorf("invalid file %s: %w", arg, err)
			}
			if info.IsDir() {
				if err := openMetadataInDir(r, arg); err != nil {
					return err
				}
				continue
			}
		}
		if _, err := r.Open(arg); err != nil {
			return fmt.Errorf("invalid file %s: %w", arg, err)
		}
	}
	return nil
}

func openMetadataInDir(r *repo.Repo, dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if path != dir && strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if !repo.IsMetadataFile(path) {
			return nil
		}
		if _, err := r.Open(path); err != nil {
			log.Warnf("invalid file %s: %s", path, err.Error())
		}
		return nil
	})
}

func checkIfChanged(o general.Tidyable, orig []byte) (changed bool) {
	o.Tidy()
	newContents, err := internal.Marshal(o)
	if err != nil {
		log.Warn("serializing failed: " + err.Error())
		return false
	}
	if !bytes.Equal(orig, newContents) {
		return true
	}
	return false
}
