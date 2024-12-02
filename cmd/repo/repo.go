package repo

import (
	"github.com/ForceCLI/force-md/repo"
)

var Metadata *repo.Repo

func init() {
	Metadata = repo.NewRepo()
}
