package profile

func (p *Profile) GetEnabledCustomPermissions() []string {
	var permissions []string
	for _, v := range p.CustomPermissions {
		if v.Enabled.ToBool() {
			permissions = append(permissions, v.Name)
		}
	}
	return permissions
}
