module github.com/octoberswimmer/force-md

go 1.15

require (
	github.com/ForceCLI/force v0.33.0 // indirect
	github.com/imdario/mergo v0.3.11
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.6.1
	github.com/thediveo/enumflag v0.10.1
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2
)

replace github.com/imdario/mergo => github.com/cwarden/mergo v0.3.12-0.20210528180603-9b708ca2c584
