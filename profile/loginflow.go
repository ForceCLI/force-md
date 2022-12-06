package profile

import (
	"fmt"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

func (p *Profile) UpdateLoginFlow(updates LoginFlow) error {
	loginFlow := p.LoginFlows
	if loginFlow == nil {
		return fmt.Errorf("No Login Flow defined")
	}
	if err := mergo.Merge(&updates, loginFlow); err != nil {
		return errors.Wrap(err, "merging settings")
	}
	p.LoginFlows = &updates
	return nil
}
