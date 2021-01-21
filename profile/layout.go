package profile

import (
	"strings"
)

func (p *Profile) GetLayouts(objectName string) LayoutAssignmentList {
	layoutPrefix := objectName + "-"
	var layouts LayoutAssignmentList
	for _, layout := range p.LayoutAssignments {
		if strings.HasPrefix(layout.Layout.Text, layoutPrefix) {
			layouts = append(layouts, layout)
		}

	}
	return layouts
}
