package cmd

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"

	cmdrepo "github.com/ForceCLI/force-md/cmd/repo"
	"github.com/ForceCLI/force-md/repo"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"

	"github.com/octoberswimmer/apexfmt/formatter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	tidyCmd.Flags().BoolP("list", "l", false, "list files that need tidying")
	tidyCmd.Flags().BoolP("recursive", "r", false, "tidy metadata files, and format Apex classes and triggers with apexfmt, in directories and their subdirectories")

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
		apexFiles, err := openTidyArgs(cmdrepo.Metadata, args, recursive)
		if err != nil {
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
		if list {
			unformatted := listUnformattedApex(apexFiles)
			for _, file := range unformatted {
				fmt.Println(file)
			}
			changes = len(unformatted) > 0 || changes
		} else {
			formatApex(apexFiles)
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
// stopping.  Also with recursive set, the Apex class and trigger source files
// found, whether named in args or below a directory in args, are returned
// for formatting with apexfmt.
func openTidyArgs(r *repo.Repo, args []string, recursive bool) ([]string, error) {
	var apexFiles []string
	for _, arg := range args {
		if recursive {
			info, err := os.Stat(arg)
			if err != nil {
				return nil, fmt.Errorf("invalid file %s: %w", arg, err)
			}
			if info.IsDir() {
				found, err := openMetadataInDir(r, arg)
				if err != nil {
					return nil, err
				}
				apexFiles = append(apexFiles, found...)
				continue
			}
			if isApexSource(arg) {
				apexFiles = append(apexFiles, arg)
				continue
			}
		}
		if _, err := r.Open(arg); err != nil {
			return nil, fmt.Errorf("invalid file %s: %w", arg, err)
		}
	}
	return apexFiles, nil
}

func openMetadataInDir(r *repo.Repo, dir string) ([]string, error) {
	var apexFiles []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if path != dir && strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if isApexSource(path) {
			apexFiles = append(apexFiles, path)
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
	return apexFiles, err
}

// isApexSource reports whether path is the source of an Apex class or
// trigger, as opposed to its -meta.xml file.
func isApexSource(path string) bool {
	return strings.HasSuffix(path, ".cls") || strings.HasSuffix(path, ".trigger")
}

// formatApex formats each file with apexfmt and writes the files whose
// formatting changed.  A file that fails to format, such as one with a
// syntax error, is logged and left unchanged.  apexfmt's errors start with
// the file name, so the log message does not repeat it.
func formatApex(files []string) {
	eachApexFile(files, func(f *formatter.Formatter) {
		changed, err := f.Changed()
		if err != nil {
			log.Warnf("formatting failed: %s", err.Error())
			return
		}
		if !changed {
			return
		}
		if err := f.Write(); err != nil {
			log.Warnf("writing %s failed: %s", f.SourceName(), err.Error())
		}
	})
}

// listUnformattedApex returns, sorted, the files whose formatting differs
// from apexfmt's.
func listUnformattedApex(files []string) []string {
	var mu sync.Mutex
	var unformatted []string
	eachApexFile(files, func(f *formatter.Formatter) {
		changed, err := f.Changed()
		if err != nil {
			log.Warnf("formatting failed: %s", err.Error())
			return
		}
		if changed {
			mu.Lock()
			unformatted = append(unformatted, f.SourceName())
			mu.Unlock()
		}
	})
	sort.Strings(unformatted)
	return unformatted
}

// eachApexFile calls fn with a formatter for each file, running one call per
// CPU at a time.
func eachApexFile(files []string, fn func(*formatter.Formatter)) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, runtime.GOMAXPROCS(0))
	for _, file := range files {
		wg.Add(1)
		sem <- struct{}{}
		go func(file string) {
			defer wg.Done()
			defer func() { <-sem }()
			fn(formatter.NewFormatter(file, nil))
		}(file)
	}
	wg.Wait()
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
