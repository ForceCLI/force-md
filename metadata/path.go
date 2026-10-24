package metadata

import (
	"path/filepath"
	"strings"
)

// PathWithinDirectory returns the part of path below the last path segment
// named dir, such as "Folder/Report.report" for
// "src/reports/Folder/Report.report" and dir "reports".  The second return
// value is false if no segment of path is named dir.  Matching whole
// segments means that a folder whose name ends with dir, like
// "RegistrationsDBreports", is not mistaken for the directory itself.
func PathWithinDirectory(path string, dir string) (string, bool) {
	parts := strings.Split(filepath.ToSlash(path), "/")
	for i := len(parts) - 2; i >= 0; i-- {
		if parts[i] == dir {
			return strings.Join(parts[i+1:], "/"), true
		}
	}
	return "", false
}
