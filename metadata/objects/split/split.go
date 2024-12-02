package split

import (
	"path/filepath"
	"regexp"

	"github.com/ForceCLI/force-md/metadata"
	log "github.com/sirupsen/logrus"
)

var pathRegex = regexp.MustCompile(`.*/objects/([^/]+)/[^/]+/([^/]+\.[^-]+)-meta\.xml$`)

func NameFromPath(path string) metadata.MetadataObjectName {
	path = filepath.ToSlash(filepath.Clean(path))

	matches := pathRegex.FindStringSubmatch(path)

	if len(matches) != 3 {
		log.Warnf("Could not match regex: %s", path)
		return metadata.NameFromPath(path)
	}

	objectName := matches[1]
	subComponentName := matches[2]

	return metadata.MetadataObjectName(objectName + "." + subComponentName)
}
