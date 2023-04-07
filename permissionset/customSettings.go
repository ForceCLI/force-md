package permissionset

import (
	"github.com/pkg/errors"

	. "github.com/ForceCLI/force-md/general"
)

var CustomSettingExistsError = errors.New("custom setting already exists")

func (p *PermissionSet) AddCustomSetting(setting string) error {
	for _, c := range p.CustomSettingAccesses {
		if c.Name == setting {
			return CustomSettingExistsError
		}
	}
	p.CustomSettingAccesses = append(p.CustomSettingAccesses, CustomSetting{
		Name:    setting,
		Enabled: TrueText,
	})
	p.CustomSettingAccesses.Tidy()
	return nil
}

func (p *PermissionSet) DeleteCustomSettings(setting string) error {
	found := false
	newSettings := p.CustomSettingAccesses[:0]
	for _, f := range p.CustomSettingAccesses {
		if f.Name == setting {
			found = true
		} else {
			newSettings = append(newSettings, f)
		}
	}
	if !found {
		return errors.New("setting not found")
	}
	p.CustomSettingAccesses = newSettings
	return nil
}

func (p *PermissionSet) GetCustomSettings() CustomSettingList {
	return p.CustomSettingAccesses
}
