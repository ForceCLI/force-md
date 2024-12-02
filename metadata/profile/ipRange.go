package profile

import (
	"github.com/pkg/errors"
)

var DuplicateIPRangeError = errors.New("login IP range already exists")

func (p *Profile) AddLoginIPRange(start string, end string, description string) error {
	for _, f := range p.LoginIPRanges {
		if f.StartAddress == start && f.EndAddress == end {
			return DuplicateIPRangeError
		}
	}

	p.LoginIPRanges = append(p.LoginIPRanges, LoginIpRange{
		StartAddress: start,
		EndAddress:   end,
		Description:  description,
	})
	p.LoginIPRanges.Tidy()
	return nil
}
