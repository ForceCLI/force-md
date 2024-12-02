package permissionset

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/permissionset"
)

var permissionName string

func init() {
	deleteUserPermissionCmd.Flags().StringVarP(&permissionName, "permission", "p", "", "user permission name")
	deleteUserPermissionCmd.MarkFlagRequired("permission")

	addUserPermissionCmd.Flags().StringVarP(&permissionName, "permission", "p", "", "user permission name")
	addUserPermissionCmd.MarkFlagRequired("permission")

	UserPermissionsCmd.AddCommand(addUserPermissionCmd)
	UserPermissionsCmd.AddCommand(deleteUserPermissionCmd)
	UserPermissionsCmd.AddCommand(listUserPermissionsCmd)
}

var UserPermissionsCmd = &cobra.Command{
	Use:   "user-permissions",
	Short: "Manage user permissions",
}

var addUserPermissionCmd = &cobra.Command{
	Use:   "add -p Permission [flags] [filename]...",
	Short: "Add user permission",
	Long:  "Add user permission in permission sets",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addUserPermission(file, permissionName)
		}
	},
}

var listUserPermissionsCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List user permissions enabled",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listUserPermissions(file)
		}
	},
}

var deleteUserPermissionCmd = &cobra.Command{
	Use:   "delete -p Permission [flags] [filename]...",
	Short: "Delete user permission",
	Long:  "Delete user permission in permissionset",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteUserPermission(file, permissionName)
		}
	},
}

func addUserPermission(file string, permissionName string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	err = p.AddUserPermission(permissionName)
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

func deleteUserPermission(file string, permissionName string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	err = p.DeleteUserPermission(permissionName)
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

func listUserPermissions(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	permissions := p.GetUserPermissions()
	for _, perm := range permissions {
		fmt.Printf("%s\n", perm.Name)
	}
}
