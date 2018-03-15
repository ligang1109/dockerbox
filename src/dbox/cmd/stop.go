package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"
	"github.com/goinbox/shell"
)

type StopCommand struct {
}

func (s *StopCommand) Run(args []string, logger golog.ILogger) {
	containerName, err := getContainerNameFromArgs(args)
	if err != nil {
		logger.Error([]byte("get containerName error: " + err.Error()))
		return
	}

	if containerName == SPECIAL_CONTAINER_NAME_ALL {
		for _, item := range dconf.Dconf {
			s.stop(item, logger)
		}
	} else {
		item, err := getDconfItemFromArgs(args)
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
