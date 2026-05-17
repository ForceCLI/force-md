//go:build unix

package permissionGranter

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

const minimalPermissionSet = `<?xml version="1.0" encoding="UTF-8"?>
<PermissionSet xmlns="http://soap.sforce.com/2006/04/metadata">
	<hasActivationRequired>false</hasActivationRequired>
	<label>Test</label>
</PermissionSet>
`

func withLowFDLimit(t *testing.T, limit uint64, fn func()) {
	t.Helper()
	var orig syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &orig); err != nil {
		t.Fatalf("getrlimit: %v", err)
	}
	new := syscall.Rlimit{Cur: limit, Max: orig.Max}
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &new); err != nil {
		t.Fatalf("setrlimit: %v", err)
	}
	t.Cleanup(func() {
		_ = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &orig)
	})
	fn()
}

// Open must close the file it opens; pre-fix this loop hit EMFILE because
// each call leaked one fd before delegating to permissionset.Open.
func TestOpenDoesNotLeakFDs(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Test.permissionset-meta.xml")
	if err := os.WriteFile(path, []byte(minimalPermissionSet), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	withLowFDLimit(t, 64, func() {
		for i := range 256 {
			if _, err := Open(path); err != nil {
				t.Fatalf("iteration %d: %v", i, err)
			}
		}
	})
}
