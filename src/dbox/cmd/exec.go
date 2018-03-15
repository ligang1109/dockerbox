package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"
	"github.com/goinbox/shell"

	"strings"
)

type ExecCommand struct {
}

func (e *ExecCommand) Run(dconfItem *dconf.DconfItem, args []string, logger golog.ILogger) {
	cmd := "docker exec -it " + dconfItem.ContainerName + " "
	cmd += dconfItem.Exec.ShellCmd + " '"
	cmd += dconfItem.Exec.PreCmd + ";"
	cmd += strings.Join(args, " ") + "'"

	logger.Debug([]byte("cmd: " + cmd))
	shell.RunCmdBindTerminal(cmd)
}
