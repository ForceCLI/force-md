package profile

import (
	"strings"
)

func (p *Profile) GetLayouts(objectName string) LayoutAssignmentList {
	layoutPrefix := objectName + "-"
	var layouts LayoutAssignmentList
	for _, layout := range p.LayoutAssignments {
		if strings.HasPrefix(layout.Layout, layoutPrefix) {
			layouts = append(layouts, layout)
		}

	}
	return layouts
}

func (p *Profile) SetObjectLayout(objectName, layoutName string) {
	layoutPrefix := objectName + "-"
	for i, f := range p.LayoutAssignments {
		if strings.HasPrefix(f.Layout, layoutPrefix) {
			p.LayoutAssignments[i].Layout = layoutPrefix + layoutName
			return
		}
	}
	p.LayoutAssignments = append(p.LayoutAssignments, LayoutAssignment{
		Layout: layoutPrefix + layoutName,
	})
	p.LayoutAssignments.Tidy()
}
