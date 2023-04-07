package permissionset

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/permissionset"
)

var customPermission string

func init() {
	deleteCustomPermissionCmd.Flags().StringVarP(&customPermission, "permission", "p", "", "custom permission name")
	deleteCustomPermissionCmd.MarkFlagRequired("permission")

	addCustomPermissionCmd.Flags().StringVarP(&customPermission, "permission", "p", "", "custom permission name")
	addCustomPermissionCmd.MarkFlagRequired("permission")

	CustomPermissionsCmd.AddCommand(addCustomPermissionCmd)
	CustomPermissionsCmd.AddCommand(deleteCustomPermissionCmd)
	CustomPermissionsCmd.AddCommand(listCustomPermissionsCmd)
}

var CustomPermissionsCmd = &cobra.Command{
	Use:   "custom-permissions",
	Short: "Manage custom permissions",
}

var addCustomPermissionCmd = &cobra.Command{
	Use:   "add -p Permission [flags] [filename]...",
	Short: "Add custom permission",
	Long:  "Add custom permission in permission sets",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addCustomPermission(file, customPermission)
		}
	},
}

var listCustomPermissionsCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List custom permissions enabled",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listCustomPermissions(file)
		}
	},
}

var deleteCustomPermissionCmd = &cobra.Command{
	Use:   "delete -p Permission [flags] [filename]...",
	Short: "Delete custom permission",
	Long:  "Delete custom permission in permissionset",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteCustomPermission(file, customPermission)
		}
	},
}

func addCustomPermission(file string, customPermission string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	err = p.AddCustomPermission(customPermission)
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

func deleteCustomPermission(file string, customPermission string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	err = p.DeleteCustomPermission(customPermission)
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

func listCustomPermissions(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	permissions := p.GetCustomPermissions()
	for _, perm := range permissions {
		fmt.Printf("%s\n", perm.Name)
	}
}
