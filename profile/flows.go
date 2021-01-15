package profile

import (
	"github.com/pkg/errors"
)

func (p *Profile) DeleteFlowAccess(flowName string) error {
	found := false
	newFlows := p.FlowAccesses[:0]
	for _, f := range p.FlowAccesses {
		if f.Flow.Text == flowName {
			found = true
		} else {
			newFlows = append(newFlows, f)
		}
	}
	if !found {
		return errors.New("flow not found")
	}
	p.FlowAccesses = newFlows
	return nil
}
