package cmd

import (
	"fmt"
	"os"

	"github.com/ForceCLI/force-md/internal"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	version     = "dev"
	silent      bool
	verbose     bool
	xmlEntities bool
)

func init() {
	cobra.OnInitialize(globalConfig)
	RootCmd.PersistentFlags().BoolVarP(&silent, "silent", "", false, "show errors only")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "", false, "show debugging output")
	RootCmd.PersistentFlags().BoolVarP(&internal.ConvertNumericXMLEntities, "convert-xml-entities", "", true, "convert numeric xml entities to character entities")

	RootCmd.MarkFlagsMutuallyExclusive("silent", "verbose")
}

var RootCmd = &cobra.Command{
	Use:   "force-md",
	Short: "force-md manipulate Salesforce metadata",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
	DisableFlagsInUseLine: true,
}

func globalConfig() {
	if silent {
		log.SetLevel(log.ErrorLevel)
	}
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
