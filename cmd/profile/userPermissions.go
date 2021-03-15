package profile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var (
	permissionName  string
	enable, disable bool
)

func init() {
	deleteUserPermissionCmd.Flags().StringVarP(&permissionName, "permission", "p", "", "user permission name")
	deleteUserPermissionCmd.MarkFlagRequired("permission")

	addUserPermissionCmd.Flags().StringVarP(&permissionName, "permission", "p", "", "user permission name")
	addUserPermissionCmd.MarkFlagRequired("permission")

	editUserPermissionCmd.Flags().StringVarP(&permissionName, "permission", "p", "", "user permission name")
	editUserPermissionCmd.Flags().BoolVarP(&enable, "enable", "e", false, "enable user permission")
	editUserPermissionCmd.Flags().BoolVarP(&disable, "disable", "d", false, "disable user permission")
	editUserPermissionCmd.MarkFlagRequired("permission")

	UserPermissionsCmd.AddCommand(addUserPermissionCmd)
	UserPermissionsCmd.AddCommand(deleteUserPermissionCmd)
	UserPermissionsCmd.AddCommand(editUserPermissionCmd)
	UserPermissionsCmd.AddCommand(listUserPermissionsCmd)
}

var UserPermissionsCmd = &cobra.Command{
	Use:   "user-permissions",
	Short: "Manage user permissions",
}

var addUserPermissionCmd = &cobra.Command{
	Use:   "add -p Permission [flags] [filename]...",
	Short: "Add user permission",
	Long:  "Add user permission in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addUserPermission(file, permissionName)
		}
	},
}

var listUserPermissionsCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List user permissions",
	Long:  "List user permissions defined in profile",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listUserPermissions(file)
		}
	},
}

var editUserPermissionCmd = &cobra.Command{
	Use:   "edit [flags] [filename]...",
	Short: "Edit user permissions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			editUserPermission(file, permissionName)
		}
	},
}

var deleteUserPermissionCmd = &cobra.Command{
	Use:   "delete -p Permission [flags] [filename]...",
	Short: "Delete user permission",
	Long:  "Delete user permission in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteUserPermission(file, permissionName)
		}
	},
}

func addUserPermission(file string, permissionName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
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
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
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

func editUserPermission(file string, permissionName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	if enable {
		err = p.EnableUserPermission(permissionName)
	} else if disable {
		err = p.DisableUserPermission(permissionName)
	}
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
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	permissions := p.GetUserPermissions()
	for _, perm := range permissions {
		enabledText := "disabled"
		if perm.Enabled.Text == "true" {
			enabledText = "enabled"
		}
		fmt.Printf("%s: %s\n", perm.Name, enabledText)
	}
}
