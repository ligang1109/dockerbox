package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"
	"github.com/goinbox/shell"
)

const (
	CMD_NAME_STOP = "stop"
)

func init() {
	Register(CMD_NAME_STOP, newStopCommand)
}

func newStopCommand() ICommand {
	return new(StopCommand)
}

type StopCommand struct {
}

func (s *StopCommand) Run(args []string, logger golog.ILogger) {
	containerKey, err := ContainerKeyFromArgs(args)
	if err != nil {
		logger.Error([]byte("get containerKey error: " + err.Error()))
		return
	}

	if containerKey == SPECIAL_CONTAINER_KEY_ALL {
		for _, item := range dconf.Dconf {
			s.stop(item, logger)
		}
	} else {
		item, err := DconfItemFromArgs(args)
		if err != nil {
			logger.Error([]byte("get dconfItem error: " + err.Error()))
			return
		}

		s.stop(item, logger)
	}
}

func (s *StopCommand) stop(dconfItem *dconf.DconfItem, logger golog.ILogger) {
	cmd := "docker stop " + dconfItem.ContainerName

	logger.Warning([]byte("stop container: " + dconfItem.ContainerName))
	logger.Debug([]byte("cmd: " + cmd))

	shell.RunCmdBindTerminal(cmd)
}
