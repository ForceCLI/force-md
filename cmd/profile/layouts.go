package profile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/profile"
)

func init() {
	showLayoutCmd.Flags().StringVarP(&objectName, "object", "o", "", "object name")
	showLayoutCmd.MarkFlagRequired("object")

	LayoutCmd.AddCommand(showLayoutCmd)
}

var LayoutCmd = &cobra.Command{
	Use:   "layout",
	Short: "Manage page layouts",
}

var showLayoutCmd = &cobra.Command{
	Use:   "show",
	Short: "Show page layout assignment",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			showLayout(file)
		}
	},
}

func showLayout(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
		return
	}
	layouts := p.GetLayouts(objectName)
	for _, l := range layouts {
		if l.RecordType != nil {
			fmt.Printf("%s (%s)\n", l.Layout.Text, l.RecordType.Text)
		} else {
			fmt.Println(l.Layout.Text)
		}
	}
}
