package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"

	"os"
)

var tlogger golog.ILogger

func init() {
	prjHome := os.Getenv("GOPATH")
	dconfPath := prjHome + "/dconf.json.demo"

	dconf.Init(dconfPath)
	tlogger, _ = golog.NewSimpleLogger(golog.NewStdoutWriter(), golog.LEVEL_DEBUG, golog.NewConsoleFormater())
}
