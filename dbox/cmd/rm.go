package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"
	"github.com/goinbox/shell"
)

const (
	CMD_NAME_RM = "rm"
)

func init() {
	register(CMD_NAME_RM, newRmCommand)
}

func newRmCommand() ICommand {
	return new(RmCommand)
}

type RmCommand struct {
}

func (r *RmCommand) Run(args []string, logger golog.ILogger) {
	containerKey, err := containerKeyFromArgs(args)
	if err != nil {
		logger.Error([]byte("get containerKey error: " + err.Error()))
		return
	}

	if containerKey == SPECIAL_CONTAINER_KEY_ALL {
		for _, item := range dconf.Dconf {
			r.rm(item, logger)
		}
	} else {
		item, err := dconfItemFromArgs(args)
		if err != nil {
			logger.Error([]byte("get dconfItem error: " + err.Error()))
			return
		}

		r.rm(item, logger)
	}
}

func (r *RmCommand) rm(dconfItem *dconf.DconfItem, logger golog.ILogger) {
	cmd := "docker rm -f " + dconfItem.ContainerName

	logger.Warning([]byte("rm container: " + dconfItem.ContainerName))
	logger.Debug([]byte("cmd: " + cmd))

	shell.RunCmdBindTerminal(cmd)
}
