package sharingrules

// AccessLevel represents an access level field
type AccessLevel struct {
	Text string `xml:",chardata"`
}

// AccountSettings represents account-specific sharing settings
type AccountSettings struct {
	CaseAccessLevel        AccessLevel `xml:"caseAccessLevel"`
	ContactAccessLevel     AccessLevel `xml:"contactAccessLevel"`
	OpportunityAccessLevel AccessLevel `xml:"opportunityAccessLevel"`
}

// Description represents a description field
type Description struct {
	Text string `xml:",innerxml"`
}

// Label represents a label field
type Label struct {
	Text string `xml:",innerxml"`
}

// Group represents a group reference
type Group struct {
	Text string `xml:",chardata"`
}

// Role represents a role reference
type Role struct {
	Text string `xml:",chardata"`
}

// RoleAndSubordinates represents a role and subordinates reference
type RoleAndSubordinates struct {
	Text string `xml:",chardata"`
}

// Queue represents a queue reference
type Queue struct {
	Text string `xml:",chardata"`
}

// AllInternalUsers represents all internal users
type AllInternalUsers struct{}

// GuestUser represents a guest user reference
type GuestUser struct {
	Text string `xml:",chardata"`
}

// SharedTo represents who a rule shares to
type SharedTo struct {
	Group                       *Group               `xml:"group"`
	Role                        *Role                `xml:"role"`
	AllCustomerPortalUsers      *AllInternalUsers    `xml:"allCustomerPortalUsers"`
	AllInternalUsers            *AllInternalUsers    `xml:"allInternalUsers"`
	AllPartnerUsers             *AllInternalUsers    `xml:"allPartnerUsers"`
	PortalRole                  *Role                `xml:"portalRole"`
	PortalRoleAndSubordinates   *RoleAndSubordinates `xml:"portalRoleAndSubordinates"`
	Queue                       *Queue               `xml:"queue"`
	RoleAndSubordinates         *RoleAndSubordinates `xml:"roleAndSubordinates"`
	RoleAndSubordinatesInternal *RoleAndSubordinates `xml:"roleAndSubordinatesInternal"`
	GuestUser                   *GuestUser           `xml:"guestUser"`
}

// SharedFrom represents who a rule shares from
type SharedFrom struct {
	RoleAndSubordinates         *RoleAndSubordinates `xml:"roleAndSubordinates"`
	RoleAndSubordinatesInternal *RoleAndSubordinates `xml:"roleAndSubordinatesInternal"`
	Group                       *Group               `xml:"group"`
	Queue                       *Queue               `xml:"queue"`
	Role                        *Role                `xml:"role"`
	PortalRole                  *Role                `xml:"portalRole"`
	PortalRoleAndSubordinates   *RoleAndSubordinates `xml:"portalRoleAndSubordinates"`
	AllCustomerPortalUsers      *AllInternalUsers    `xml:"allCustomerPortalUsers"`
	AllInternalUsers            *AllInternalUsers    `xml:"allInternalUsers"`
	AllPartnerUsers             *AllInternalUsers    `xml:"allPartnerUsers"`
}

// BooleanFilter represents a boolean filter expression
type BooleanFilter struct {
	Text string `xml:",chardata"`
}

// Field represents a field reference
type Field struct {
	Text string `xml:",chardata"`
}

// Operation represents an operation
type Operation struct {
	Text string `xml:",chardata"`
}

// Value represents a value
type Value struct {
	Text string `xml:",chardata"`
}

// CriteriaItem represents a single criteria item in a rule
type CriteriaItem struct {
	Field     Field     `xml:"field"`
	Operation Operation `xml:"operation"`
	Value     Value     `xml:"value"`
}
