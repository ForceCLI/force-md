package reportFolder

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/reportFolder"
)

type ShareType enumflag.Flag
type AccessLevel enumflag.Flag

var shareType ShareType
var accessLevel AccessLevel

const (
	NoneShareType ShareType = iota
	User
	Role
	RoleAndSubordinates
	Organization
	Group
)

const (
	NoneAccessLevel AccessLevel = iota
	View
	Manage
	EditAllContents
)

var ShareTypeIds = map[ShareType][]string{
	NoneShareType:       {"None"},
	User:                {"User"},
	Role:                {"Role"},
	RoleAndSubordinates: {"RoleAndSubordinates"},
	Organization:        {"Organization"},
	Group:               {"Group"},
}

var AccessLevelIds = map[AccessLevel][]string{
	NoneAccessLevel: {"None"},
	View:            {"View"},
	Manage:          {"Manage"},
	EditAllContents: {"EditAllContents"},
}

func init() {
	deleteShareCmd.Flags().VarP(enumflag.New(&shareType, "type", ShareTypeIds, enumflag.EnumCaseInsensitive),
		"type", "t", "type; can be 'User', 'Role', 'RoleAndSubordinates', 'Organization', or 'Group'")
	deleteShareCmd.Flags().VarP(enumflag.New(&accessLevel, "access", AccessLevelIds, enumflag.EnumCaseInsensitive),
		"access", "l", "access level; can be 'View', 'Manage', or 'EditAllContents'")

	listSharesCmd.Flags().VarP(enumflag.New(&shareType, "type", ShareTypeIds, enumflag.EnumCaseInsensitive),
		"type", "t", "type; can be 'User', 'Role', 'RoleAndSubordinates', 'Organization', or 'Group'")
	listSharesCmd.Flags().VarP(enumflag.New(&accessLevel, "access", AccessLevelIds, enumflag.EnumCaseInsensitive),
		"access", "a", "access level; can be 'View', 'Manage', or 'EditAllContents'")
	FolderSharesCmd.AddCommand(listSharesCmd)
	FolderSharesCmd.AddCommand(deleteShareCmd)
}

var FolderSharesCmd = &cobra.Command{
	Use:   "shares",
	Short: "Manage Folder Sharing",
}

var listSharesCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List folder shares",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listShares(file)
		}
	},
}

var deleteShareCmd = &cobra.Command{
	Use:   "delete [flags] [filename]...",
	Short: "Delete folder shares",
	Long:  "Delete folder shares",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteShares(file)
		}
	},
}

func listShares(file string) {
	w, err := reportFolder.Open(file)
	if err != nil {
		log.Warn("parsing report folder failed: " + err.Error())
		return
	}
	var filters []reportFolder.FolderShareFilter
	switch accessLevel {
	case View, Manage, EditAllContents:
		filters = append(filters, func(a reportFolder.FolderShare) bool {
			return strings.ToLower(a.AccessLevel) == strings.ToLower(AccessLevelIds[accessLevel][0])
		})
	}
	switch shareType {
	case User, Role, RoleAndSubordinates, Organization, Group:
		filters = append(filters, func(a reportFolder.FolderShare) bool {
			return strings.ToLower(a.SharedToType) == strings.ToLower(ShareTypeIds[shareType][0])
		})
	}
	shares := w.GetShares(filters...)
	for _, r := range shares {
		fmt.Printf("%s: %s\n", w.Name, r.SharedTo)
	}
}

func deleteShares(file string) {
	w, err := reportFolder.Open(file)
	if err != nil {
		log.Warn("parsing report folder failed: " + err.Error())
		return
	}
	var filters []reportFolder.FolderShareFilter
	switch accessLevel {
	case View, Manage, EditAllContents:
		filters = append(filters, func(a reportFolder.FolderShare) bool {
			return strings.ToLower(a.AccessLevel) == strings.ToLower(AccessLevelIds[accessLevel][0])
		})
	}
	switch shareType {
	case User, Role, RoleAndSubordinates, Organization, Group:
		filters = append(filters, func(a reportFolder.FolderShare) bool {
			return strings.ToLower(a.SharedToType) == strings.ToLower(ShareTypeIds[shareType][0])
		})
	}
	w.DeleteShares(filters...)
	err = internal.WriteToFile(w, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
