module github.com/octoberswimmer/force-md

go 1.15

require (
	github.com/imdario/mergo v0.3.11
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v0.0.7
	github.com/thediveo/enumflag v0.10.1 // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
)

replace github.com/imdario/mergo => github.com/cwarden/mergo v0.3.12-0.20210528180603-9b708ca2c584
