package repo

import (
	"runtime"
	"sync"

	"github.com/ForceCLI/force-md/metadata"
)

// OpenParallel parses the given files concurrently and registers each
// successfully parsed item in the repo. The expensive part of opening a
// metadata file is the XML decode in MetadataFromPath, which has no shared
// mutable state, so it is run across a pool of workers. Registration into the
// repo's maps is serialized after parsing because those maps are not safe for
// concurrent writes.
//
// It returns a slice of errors parallel to paths: a nil entry means the file
// was parsed and registered successfully. The repo registration order is
// unspecified, matching the existing map-backed storage.
func (o *Repo) OpenParallel(paths []string) []error {
	errs := make([]error, len(paths))
	if len(paths) == 0 {
		return errs
	}

	results := make([]metadata.RegisterableMetadata, len(paths))

	workers := min(max(runtime.GOMAXPROCS(0), 1), len(paths))

	indexes := make(chan int)
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range indexes {
				m, err := MetadataFromPath(paths[i])
				results[i] = m
				errs[i] = err
			}
		}()
	}
	for i := range paths {
		indexes <- i
	}
	close(indexes)
	wg.Wait()

	for i, m := range results {
		if errs[i] == nil && m != nil {
			o.AddItem(m)
		}
	}
	return errs
}
