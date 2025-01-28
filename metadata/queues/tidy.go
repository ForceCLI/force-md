package queue

import "sort"

func (roleAndSubordinates RoleAndSubordinates) Tidy() {
	sort.Slice(roleAndSubordinates.RoleAndSubordinate, func(i, j int) bool {
		return roleAndSubordinates.RoleAndSubordinate[i] < roleAndSubordinates.RoleAndSubordinate[j]
	})
}

func (roles Roles) Tidy() {
	sort.Slice(roles.Role, func(i, j int) bool {
		return roles.Role[i] < roles.Role[j]
	})
}

func (publicGroups PublicGroups) Tidy() {
	sort.Slice(publicGroups.PublicGroup, func(i, j int) bool {
		return publicGroups.PublicGroup[i] < publicGroups.PublicGroup[j]
	})
}

func (users Users) Tidy() {
	sort.Slice(users.User, func(i, j int) bool {
		return users.User[i] < users.User[j]
	})
}
