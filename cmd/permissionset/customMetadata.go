package permissionset

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/permissionset"
)

var customMetadataType string

func init() {
	deleteCustomMetadataTypeCmd.Flags().StringVarP(&customMetadataType, "type", "t", "", "custom metadata type")
	deleteCustomMetadataTypeCmd.MarkFlagRequired("type")

	addCustomMetadataTypeCmd.Flags().StringVarP(&customMetadataType, "type", "t", "", "custom metadata type")
	addCustomMetadataTypeCmd.MarkFlagRequired("type")

	CustomMetadataTypesCmd.AddCommand(addCustomMetadataTypeCmd)
	CustomMetadataTypesCmd.AddCommand(deleteCustomMetadataTypeCmd)
	CustomMetadataTypesCmd.AddCommand(listCustomMetadataTypesCmd)
}

var CustomMetadataTypesCmd = &cobra.Command{
	Use:   "custom-metadata",
	Short: "Manage custom metadata types",
}

var addCustomMetadataTypeCmd = &cobra.Command{
	Use:   "add -p MetadataType [flags] [filename]...",
	Short: "Add custom metadata type",
	Long:  "Add custom metadata type in permission sets",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addCustomMetadataType(file, customMetadataType)
		}
	},
}

var listCustomMetadataTypesCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List custom metadata types enabled",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listCustomMetadataTypes(file)
		}
	},
}

var deleteCustomMetadataTypeCmd = &cobra.Command{
	Use:   "delete -p MetadataType [flags] [filename]...",
	Short: "Delete custom metadata type",
	Long:  "Delete custom metadata type in permissionset",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteCustomMetadataType(file, customMetadataType)
		}
	},
}

func addCustomMetadataType(file string, customMetadataType string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	err = p.AddCustomMetadataType(customMetadataType)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func deleteCustomMetadataType(file string, customMetadataType string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	err = p.DeleteCustomMetadataType(customMetadataType)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func listCustomMetadataTypes(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	metadataTypes := p.GetCustomMetadataTypes()
	for _, md := range metadataTypes {
		fmt.Printf("%s\n", md.Name)
	}
}
