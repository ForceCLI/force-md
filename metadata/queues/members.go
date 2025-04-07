package queue

import (
	"github.com/pkg/errors"
)

var MemberExistsError = errors.New("member already exists")
var SobjectExistsError = errors.New("sobject already exists")

func (q *Queue) AddRole(memberName string) error {
	if q.QueueMembers.Roles == nil {
		q.QueueMembers.Roles = &Roles{}
	}
	for _, c := range q.QueueMembers.Roles.Role {
		if c == memberName {
			return MemberExistsError
		}
	}
	q.QueueMembers.Roles.Role = append(q.QueueMembers.Roles.Role, memberName)
	q.QueueMembers.Roles.Tidy()
	return nil
}

func (q *Queue) AddRoleAndSubordinate(memberName string) error {
	if q.QueueMembers.RoleAndSubordinates == nil {
		q.QueueMembers.RoleAndSubordinates = &RoleAndSubordinates{}
	}
	for _, c := range q.QueueMembers.RoleAndSubordinates.RoleAndSubordinate {
		if c == memberName {
			return MemberExistsError
		}
	}
	q.QueueMembers.RoleAndSubordinates.RoleAndSubordinate = append(q.QueueMembers.RoleAndSubordinates.RoleAndSubordinate, memberName)
	q.QueueMembers.RoleAndSubordinates.Tidy()
	return nil
}

func (q *Queue) AddPublicGroup(memberName string) error {
	if q.QueueMembers.PublicGroups == nil {
		q.QueueMembers.PublicGroups = &PublicGroups{}
	}
	for _, c := range q.QueueMembers.PublicGroups.PublicGroup {
		if c == memberName {
			return MemberExistsError
		}
	}
	q.QueueMembers.PublicGroups.PublicGroup = append(q.QueueMembers.PublicGroups.PublicGroup, memberName)
	q.QueueMembers.PublicGroups.Tidy()
	return nil
}

func (q *Queue) AddUser(memberName string) error {
	if q.QueueMembers.Users == nil {
		q.QueueMembers.Users = &Users{}
	}
	for _, c := range q.QueueMembers.Users.User {
		if c == memberName {
			return MemberExistsError
		}
	}
	q.QueueMembers.Users.User = append(q.QueueMembers.Users.User, memberName)
	q.QueueMembers.Users.Tidy()
	return nil
}

func (q *Queue) AddSobject(sobject string) error {
	if q.QueueSobject == nil {
		q.QueueMembers.Users = &Users{}
	}
	for _, c := range q.QueueSobject {
		if c.SobjectType == sobject {
			return SobjectExistsError
		}
	}
	q.QueueSobject = append(q.QueueSobject, SobjectType{SobjectType: sobject})
	q.QueueSobject.Tidy()
	return nil
}
