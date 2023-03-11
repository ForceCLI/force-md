package permissionGranter

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/octoberswimmer/force-md/permissionset"
	"github.com/octoberswimmer/force-md/profile"
	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

type PermissionGranter interface {
	GetVisibleApplications() []string
	GetEnabledClasses() []string
	GetEnabledCustomPermissions() []string
	GetGrantedFieldPermissions() []permissionset.FieldPermissions
	GetGrantedObjectPermissions() []permissionset.ObjectPermissions
	GetEnabledPageAccesses() []string
	GetVisibleRecordTypes() []string
	GetEnabledUserPermissions() []string
	/*
		* TODO: Decide how to map tab visiblity between profiles and permission sets
			GetVisibleTabs() []string
	*/
}

func Open(path string) (PermissionGranter, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = true

	type granter struct {
		XMLName xml.Name
	}

	var g granter
	if err := dec.Decode(&g); err != nil {
		return nil, errors.Wrap(err, "parsing xml in "+path)
	}
	switch g.XMLName.Local {
	case "Profile":
		return profile.Open(path)
	case "PermissionSet":
		return permissionset.Open(path)
	default:
		return nil, fmt.Errorf("Invalid type: " + g.XMLName.Local)
	}
}
