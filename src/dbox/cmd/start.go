package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"
	"github.com/goinbox/shell"
)

type StartCommand struct {
}

func (s *StartCommand) Run(args []string, logger golog.ILogger) {
	containerKey, err := getContainerKeyFromArgs(args)
	if err != nil {
		logger.Error([]byte("get containerKey error: " + err.Error()))
		return
	}

	if containerKey == SPECIAL_CONTAINER_KEY_ALL {
		for _, item := range dconf.Dconf {
			s.start(item, logger)
		}
	} else {
		item, err := getDconfItemFromArgs(args)
		if err != nil {
			logger.Error([]byte("get dconfItem error: " + err.Error()))
			return
		}

		s.start(item, logger)
	}
}

func (s *StartCommand) start(dconfItem *dconf.DconfItem, logger golog.ILogger) {
	cmd := "docker start " + dconfItem.ContainerName

	logger.Warning([]byte("start container: " + dconfItem.ContainerName))
	logger.Debug([]byte("cmd: " + cmd))

	shell.RunCmdBindTerminal(cmd)
}
