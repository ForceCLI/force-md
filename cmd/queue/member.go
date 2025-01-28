package queue

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"

	"github.com/ForceCLI/force-md/internal"
	queue "github.com/ForceCLI/force-md/metadata/queues"
)

var (
	memberType MemberType
	member     string
)

type MemberType enumflag.Flag

const (
	None MemberType = iota
	Role
	RoleAndSubordinate
	PublicGroup
	User
)

var MemberTypeIds = map[MemberType][]string{
	None:               {"None"},
	Role:               {"Role"},
	RoleAndSubordinate: {"RoleAndSubordinate"},
	PublicGroup:        {"PublicGroup"},
	User:               {"User"},
}

func init() {
	addMemberCmd.Flags().StringVarP(&member, "member", "m", "", "member")
	addMemberCmd.Flags().VarP(enumflag.New(&memberType, "membertype", MemberTypeIds, enumflag.EnumCaseInsensitive),
		"membertype", "t", "member type; can be 'Role', 'RoleAndSubordinate', 'PublicGroup', or 'User'")

	addMemberCmd.MarkFlagRequired("member")
	addMemberCmd.MarkFlagRequired("membertype")

	MemberCmd.AddCommand(addMemberCmd)
}

var MemberCmd = &cobra.Command{
	Use:   "member",
	Short: "Manage queue members",
}

var addMemberCmd = &cobra.Command{
	Use:   "add -t MemberType -m Member [filename]...",
	Short: "Add queue member",
	Long:  "Add queue member",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addMember(file, memberType, member)
		}
	},
}

func addMember(file string, memberType MemberType, member string) {
	o, err := queue.Open(file)
	if err != nil {
		log.Warn("parsing queue failed: " + err.Error())
		return
	}
	switch memberType {
	case Role:
		err = o.AddRole(member)
	case RoleAndSubordinate:
		err = o.AddRoleAndSubordinate(member)
	case PublicGroup:
		err = o.AddPublicGroup(member)
	case User:
		err = o.AddUser(member)
	}
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
