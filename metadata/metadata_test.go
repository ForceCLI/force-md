//go:build unix

package metadata

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/nbio/xml"
)

type fakeMetadata struct {
	XMLName xml.Name `xml:"Account"`
	MetadataInfo
}

func (f *fakeMetadata) SetMetadata(m MetadataInfo) { f.MetadataInfo = m }

func writeFixture(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "Account.object")
	if err := os.WriteFile(path, []byte(`<?xml version="1.0" encoding="UTF-8"?><Account/>`), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

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

// Calling ParseMetadataXml more times than the open-file limit must succeed —
// each call has to close the file it opens. Pre-fix this looped to EMFILE.
func TestParseMetadataXmlDoesNotLeakFDs(t *testing.T) {
	path := writeFixture(t)
	withLowFDLimit(t, 64, func() {
		for i := range 256 {
			var m fakeMetadata
			if err := ParseMetadataXml(&m, path); err != nil {
				t.Fatalf("iteration %d: %v", i, err)
			}
		}
	})
}

func TestParseMetadataXmlIfPossibleDoesNotLeakFDs(t *testing.T) {
	path := writeFixture(t)
	withLowFDLimit(t, 64, func() {
		for i := range 256 {
			var m fakeMetadata
			if _, err := ParseMetadataXmlIfPossible(&m, path); err != nil {
				t.Fatalf("iteration %d: %v", i, err)
			}
		}
	})
}
