package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"

	"testing"
	"os"
)

func TestRconf(t *testing.T) {
	prjHome := os.Getenv("GOPATH")
	dconfPath := prjHome + "/dconf.json.demo"

	dconf.Init(dconfPath)
	logger, _ := golog.NewSimpleLogger(golog.NewStdoutWriter(), golog.LEVEL_DEBUG, golog.NewConsoleFormater())

	cmd := new(ExecCommand)
	cmd.Run(dconf.Dconf["mysql"], []string{"mysql", "-uroot", "-p123"}, logger)
}
