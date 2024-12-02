package profile

import (
	"strings"

	"github.com/ForceCLI/force-md/metadata/permissionset"
	"github.com/pkg/errors"
)

type FlowFilter func(permissionset.FlowAccess) bool

func (p *Profile) DeleteFlowAccess(flowName string) error {
	found := false
	newFlows := p.FlowAccesses[:0]
	for _, f := range p.FlowAccesses {
		if strings.ToLower(f.Flow) == strings.ToLower(flowName) {
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

func (p *Profile) GetFlows(filters ...FlowFilter) permissionset.FlowAccessList {
	var flowAccesses permissionset.FlowAccessList
FLOWS:
	for _, t := range p.FlowAccesses {
		for _, filter := range filters {
			if !filter(t) {
				continue FLOWS
			}
		}
		flowAccesses = append(flowAccesses, t)
	}
	return flowAccesses
}
