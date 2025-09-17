package objects

import (
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/objects"
)

var (
	webLinkName string
	urlFilter   string
)

func init() {
	deleteWebLinkCmd.Flags().StringVarP(&webLinkName, "weblink", "w", "", "web link name")
	deleteWebLinkCmd.MarkFlagRequired("weblink")

	listWebLinksCmd.Flags().StringVarP(&urlFilter, "url", "u", "", "filter weblinks containing this string in the URL")

	WebLinkCmd.AddCommand(listWebLinksCmd)
	WebLinkCmd.AddCommand(deleteWebLinkCmd)
}

var WebLinkCmd = &cobra.Command{
	Use:                   "weblink",
	Short:                 "Manage object web link metadata",
	DisableFlagsInUseLine: true,
}

var deleteWebLinkCmd = &cobra.Command{
	Use:                   "delete -s WebLink [filename]...",
	Short:                 "Delete object web link",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteWebLink(file, webLinkName)
		}
	},
}

var listWebLinksCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List object web links",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listWebLinks(file)
		}
	},
}

func listWebLinks(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
	webLinks := o.GetWebLinks()
	for _, f := range webLinks {
		// Apply URL filter if specified
		if urlFilter != "" {
			if f.URL == nil || !strings.Contains(f.URL.Text, urlFilter) {
				continue
			}
		}
		fmt.Printf("%s.%s\n", objectName, f.FullName)
	}
}

func deleteWebLink(file string, webLinkName string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
	webLinkName = strings.TrimPrefix(webLinkName, objectName+".")
	err = o.DeleteWebLink(webLinkName)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(o, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
