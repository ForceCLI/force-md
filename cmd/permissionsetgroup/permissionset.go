package permissionsetgroup

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/permissionsetgroup"
)

var permissionSetName string

func init() {
	addPermissionSetCmd.Flags().StringVarP(&permissionSetName, "permissionset", "p", "", "permission set name")
	addPermissionSetCmd.MarkFlagRequired("permissionset")

	deletePermissionSetCmd.Flags().StringVarP(&permissionSetName, "permissionset", "p", "", "permission set name")
	deletePermissionSetCmd.MarkFlagRequired("permissionset")

	PermissionSetCmd.AddCommand(addPermissionSetCmd)
	PermissionSetCmd.AddCommand(deletePermissionSetCmd)
	PermissionSetCmd.AddCommand(listPermissionSetsCmd)
}

var PermissionSetCmd = &cobra.Command{
	Use:   "permissionset",
	Short: "Manage Permission Sets",
}

var addPermissionSetCmd = &cobra.Command{
	Use:                   "add -p PermissionSetName [flags] [filename]...",
	Short:                 "Add Permission Set to Permission Set",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addPermissionSet(file, permissionSetName)
		}
	},
}

var deletePermissionSetCmd = &cobra.Command{
	Use:                   "delete -p PermissionSetName [filename]...",
	Short:                 "Delete permission set",
	Long:                  "Delete permission set in permission set groups",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deletePermissionSet(file, permissionSetName)
		}
	},
}

var listPermissionSetsCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List permission sets",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listPermissionSets(file)
		}
	},
}

func addPermissionSet(file, permissionSetName string) {
	p, err := permissionsetgroup.Open(file)
	if err != nil {
		log.Warn("parsing permission set group failed: " + err.Error())
		return
	}
	p.AddPermissionSet(permissionSetName)
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func deletePermissionSet(file string, permissionSetName string) {
	p, err := permissionsetgroup.Open(file)
	if err != nil {
		log.Warn("parsing permission set group failed: " + err.Error())
		return
	}
	err = p.DeletePermissionSet(permissionSetName)
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

func listPermissionSets(file string) {
	p, err := permissionsetgroup.Open(file)
	if err != nil {
		log.Warn("parsing permission set group group failed: " + err.Error())
		return
	}
	permissionSets := p.GetPermissionSets()
	for _, a := range permissionSets {
		fmt.Println(a.Text)
	}
}
