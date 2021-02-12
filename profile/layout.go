package profile

import (
	"errors"
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

func (p *Profile) SetObjectLayout(objectName, layoutName string) error {
	layoutPrefix := objectName + "-"
	for i, f := range p.LayoutAssignments {
		if strings.HasPrefix(f.Layout.Text, layoutPrefix) {
			p.LayoutAssignments[i].Layout.Text = layoutPrefix + layoutName
			return nil
		}
	}
	return errors.New("object layout not found")
}
