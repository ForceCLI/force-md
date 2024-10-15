package cmd

import (
	"bytes"
	"fmt"
	"os"

	_ "github.com/ForceCLI/force-md/flow"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	tidyCmd.Flags().BoolP("list", "l", false, "list files that need tidying")

	RootCmd.AddCommand(tidyCmd)
}

var tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "Tidy Metadata",
	Example: `
$ force-md tidy sfdx/main/default/objects/*/{fields,validationRules}/* sfdx/main/default/flows/*

$ force-md tidy src/objects/*
`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, file := range args {
			_, err := internal.Metadata.Open(file)
			if err != nil {
				return fmt.Errorf("invalid file %s: %w", file, err)
			}
		}
		changes := false
		list, _ := cmd.Flags().GetBool("list")
		for _, t := range internal.Metadata.Types() {
			for _, m := range internal.Metadata.Items(t) {
				file := m.GetMetadataInfo().Path()
				o, ok := m.(general.Tidyable)
				if !ok {
					log.Warnf("file %s of type %T is not tidyable", file, m)
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
