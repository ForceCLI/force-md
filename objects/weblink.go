package objects

import (
	"github.com/pkg/errors"

	"github.com/ForceCLI/force-md/objects/weblink"
)

func (o *CustomObject) GetWebLinks() []weblink.WebLink {
	return o.WebLinks
}

func (p *CustomObject) DeleteWebLink(webLinkName string) error {
	found := false
	newWebLinks := p.WebLinks[:0]
	for _, f := range p.WebLinks {
		if f.FullName == webLinkName {
			found = true
		} else {
			newWebLinks = append(newWebLinks, f)
		}
	}
	if !found {
		return errors.New("web link not found")
	}
	p.WebLinks = newWebLinks
	return nil
}
