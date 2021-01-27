package profile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var permissionName string

func init() {
	deleteUserPermissionCmd.Flags().StringVarP(&permissionName, "permission", "p", "", "user permission name")
	deleteUserPermissionCmd.MarkFlagRequired("permission")

	addUserPermissionCmd.Flags().StringVarP(&permissionName, "permission", "p", "", "user permission name")
	addUserPermissionCmd.MarkFlagRequired("permission")

	UserPermissionsCmd.AddCommand(addUserPermissionCmd)
	UserPermissionsCmd.AddCommand(deleteUserPermissionCmd)
}

var UserPermissionsCmd = &cobra.Command{
	Use:   "user-permissions",
	Short: "Manage user permissions",
}

var addUserPermissionCmd = &cobra.Command{
	Use:   "add",
	Short: "Add user permission",
	Long:  "Add user permission in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addUserPermission(file, permissionName)
		}
	},
}

var deleteUserPermissionCmd = &cobra.Command{
	Use:   "delete",
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
