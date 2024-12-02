package profile

import (
	"errors"
	"strings"
)

type LayoutFilter func(LayoutAssignment) bool

func (p *Profile) GetLayouts(filters ...LayoutFilter) LayoutAssignmentList {
	var layouts LayoutAssignmentList
LAYOUTS:
	for _, layout := range p.LayoutAssignments {
		for _, filter := range filters {
			if !filter(layout) {
				continue LAYOUTS
			}
		}
		layouts = append(layouts, layout)

	}
	return layouts
}

func (p *Profile) SetObjectLayout(objectName, layoutName string) {
	layoutPrefix := objectName + "-"
	for i, f := range p.LayoutAssignments {
		if f.RecordType != nil {
			continue
		}
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

func (p *Profile) SetObjectLayoutForRecordType(objectName, layoutName, recordType string) {
	layoutPrefix := objectName + "-"
	recordType = strings.TrimPrefix(recordType, objectName+".")
	fullRecordTypeName := objectName + "." + recordType
	for i, f := range p.LayoutAssignments {
		if strings.HasPrefix(f.Layout, layoutPrefix) && f.RecordType != nil && f.RecordType.Text == fullRecordTypeName {
			p.LayoutAssignments[i].Layout = layoutPrefix + layoutName
			return
		}
	}
	p.LayoutAssignments = append(p.LayoutAssignments, LayoutAssignment{
		Layout: layoutPrefix + layoutName,
		RecordType: &RecordType{
			Text: fullRecordTypeName,
		},
	})
	p.LayoutAssignments.Tidy()
}

func (p *Profile) DeleteObjectLayout(objectName string, filters ...LayoutFilter) error {
	found := false
	newLayouts := p.LayoutAssignments[:0]
	layoutPrefix := strings.ToLower(objectName + "-")
LAYOUTS:
	for _, l := range p.LayoutAssignments {
		if !strings.HasPrefix(strings.ToLower(l.Layout), layoutPrefix) {
			newLayouts = append(newLayouts, l)
			continue
		}
		for _, filter := range filters {
			if !filter(l) {
				newLayouts = append(newLayouts, l)
				continue LAYOUTS
			}
		}
		found = true
	}
	if !found {
		return errors.New("layout assignment not found")
	}
	p.LayoutAssignments = newLayouts
	p.LayoutAssignments.Tidy()
	return nil
}
