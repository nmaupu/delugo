package main

import (
	"github.com/nmaupu/delugo/cli"
)

const (
	appName = "delugo"
	appDesc = "Deals with a deluged torrents daemon"
)

var (
	appVersion string
)

func main() {
	if appVersion == "" {
		appVersion = "master"
	}
	cli.Process(appName, appDesc, appVersion)
}
