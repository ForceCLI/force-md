package custommetadata

import (
	"sort"
)

func (p *CustomMetadata) Tidy() {
	sort.Slice(p.Values, func(i, j int) bool {
		return p.Values[i].Field < p.Values[j].Field
	})
}
