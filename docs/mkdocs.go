package main

import (
	cmd "github.com/octoberswimmer/force-md/cmd"

	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCmd, "./docs")
	if err != nil {
		panic(err.Error())
	}
}
