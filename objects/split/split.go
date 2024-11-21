package split

import (
	"path/filepath"
	"regexp"

	"github.com/ForceCLI/force-md/internal"

	log "github.com/sirupsen/logrus"
)

var pathRegex = regexp.MustCompile(`.*/objects/([^/]+)/[^/]+/([^/]+\.[^-]+)-meta\.xml$`)

func NameFromPath(path string) string {
	path = filepath.ToSlash(filepath.Clean(path))

	matches := pathRegex.FindStringSubmatch(path)

	if len(matches) != 3 {
		log.Warnf("Could not match regex: %s", path)
		return internal.NameFromPath(path)
	}

	objectName := matches[1]
	subComponentName := matches[2]

	return objectName + "." + subComponentName
}
