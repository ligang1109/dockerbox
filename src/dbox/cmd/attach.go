package cmd

import (
	"github.com/goinbox/golog"
	"github.com/goinbox/shell"
)

type AttachCommand struct {
}

func (a *AttachCommand) Run(args []string, logger golog.ILogger) {
	dconfItem, err := getDconfItemFromArgs(args)
	if err != nil {
		logger.Error([]byte("get dconfItem error: " + err.Error()))
		return
	}

	cmd := "sudo nsenter --target `docker inspect --format {{.State.Pid}} "
	cmd += dconfItem.ContainerName
	cmd += "` --mount --uts --ipc --net --pid"

	logger.Debug([]byte("cmd: " + cmd))

	shell.RunCmdBindTerminal(cmd)
}
