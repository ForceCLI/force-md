package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	silent  bool
)

func init() {
	cobra.OnInitialize(globalConfig)
	RootCmd.PersistentFlags().BoolVarP(&silent, "silent", "", false, "show errors only")
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
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
