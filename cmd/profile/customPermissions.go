package profile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/permissionset"
	"github.com/ForceCLI/force-md/metadata/profile"
)

func init() {
	deleteCustomPermissionCmd.Flags().StringVarP(&permissionName, "permission", "p", "", "custom permission name")
	deleteCustomPermissionCmd.MarkFlagRequired("permission")

	addCustomPermissionCmd.Flags().StringVarP(&permissionName, "permission", "p", "", "custom permission name")
	addCustomPermissionCmd.MarkFlagRequired("permission")

	editCustomPermissionCmd.Flags().StringVarP(&permissionName, "permission", "p", "", "custom permission name")
	editCustomPermissionCmd.Flags().BoolVarP(&enable, "enable", "e", false, "enable custom permission")
	editCustomPermissionCmd.Flags().BoolVarP(&disable, "disable", "d", false, "disable custom permission")
	editCustomPermissionCmd.MarkFlagRequired("permission")

	listCustomPermissionsCmd.Flags().BoolVarP(&enable, "enabled", "e", false, "enables custom permissions")
	listCustomPermissionsCmd.Flags().BoolVarP(&disable, "disabled", "d", false, "disabled custom permissions")

	CustomPermissionsCmd.AddCommand(addCustomPermissionCmd)
	CustomPermissionsCmd.AddCommand(deleteCustomPermissionCmd)
	CustomPermissionsCmd.AddCommand(editCustomPermissionCmd)
	CustomPermissionsCmd.AddCommand(listCustomPermissionsCmd)
}

var CustomPermissionsCmd = &cobra.Command{
	Use:   "custom-permissions",
	Short: "Manage custom permissions",
}

var addCustomPermissionCmd = &cobra.Command{
	Use:   "add -p Permission [flags] [filename]...",
	Short: "Add custom permission",
	Long:  "Add custom permission in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addCustomPermission(file, permissionName)
		}
	},
}

var listCustomPermissionsCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List custom permissions",
	Long:  "List custom permissions defined in profile",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listCustomPermissions(file)
		}
	},
}

var editCustomPermissionCmd = &cobra.Command{
	Use:   "edit [flags] [filename]...",
	Short: "Edit custom permissions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			editCustomPermission(file, permissionName)
		}
	},
}

var deleteCustomPermissionCmd = &cobra.Command{
	Use:   "delete -p Permission [flags] [filename]...",
	Short: "Delete custom permission",
	Long:  "Delete custom permission in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteCustomPermission(file, permissionName)
		}
	},
}

func addCustomPermission(file string, permissionName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.AddCustomPermission(permissionName)
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

func deleteCustomPermission(file string, permissionName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.DeleteCustomPermission(permissionName)
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

func editCustomPermission(file string, permissionName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	if enable {
		err = p.EnableCustomPermission(permissionName)
	} else if disable {
		err = p.DisableCustomPermission(permissionName)
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

func listCustomPermissions(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	var filters []profile.CustomPermissionFilter
	if enable {
		filters = append(filters, func(u permissionset.CustomPermission) bool {
			return u.Enabled.ToBool()
		})
	}
	if disable {
		filters = append(filters, func(u permissionset.CustomPermission) bool {
			return !u.Enabled.ToBool()
		})
	}
	permissions := p.GetCustomPermissions(filters...)
	for _, perm := range permissions {
		enabledText := "disabled"
		if perm.Enabled.Text == "true" {
			enabledText = "enabled"
		}
		fmt.Printf("%s: %s\n", perm.Name, enabledText)
	}
}
