package cmd

import (
	"github.com/goinbox/golog"
	"github.com/goinbox/shell"

	"strings"
)

type ExecCommand struct {
}

func (e *ExecCommand) Run(args []string, logger golog.ILogger) {
	dconfItem, err := getDconfItemFromArgs(args)
	if err != nil {
		logger.Error([]byte("get dconfItem error: " + err.Error()))
		return
	}

	if dconfItem.Exec == nil {
		logger.Error([]byte("container: " + dconfItem.ContainerName + " not has exec conf"))
		return
	}

	cmd := "docker exec -it " + dconfItem.ContainerName + " "
	if dconfItem.Exec.ShellCmd != "" {
		cmd += dconfItem.Exec.ShellCmd + " '"
		if dconfItem.Exec.PreCmd != "" {
			cmd += dconfItem.Exec.PreCmd + ";"
		}
	}
	cmd += strings.Join(args[1:], " ")
	if dconfItem.Exec.ShellCmd != "" {
		cmd += "'"
	}

	logger.Debug([]byte("cmd: " + cmd))

	shell.RunCmdBindTerminal(cmd)
}
