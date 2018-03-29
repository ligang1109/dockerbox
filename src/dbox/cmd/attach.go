package cmd

import (
	"github.com/goinbox/golog"
	"github.com/goinbox/shell"
)

const (
	CMD_NAME_ATTACH = "attach"
)

func init() {
	register(CMD_NAME_ATTACH, newAttachCommand)
}

func newAttachCommand() ICommand {
	return new(AttachCommand)
}

type AttachCommand struct {
}

func (a *AttachCommand) Run(args []string, logger golog.ILogger) {
	dconfItem, err := dconfItemFromArgs(args)
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
