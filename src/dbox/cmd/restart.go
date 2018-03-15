package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"
)

type RestartCommand struct {
}

func (s *RestartCommand) Run(args []string, logger golog.ILogger) {
	containerName, err := getContainerNameFromArgs(args)
	if err != nil {
		logger.Error([]byte("get dconfItem error: " + err.Error()))
		return
	}

	if containerName == SPECIAL_CONTAINER_NAME_ALL {
		for name, _ := range dconf.Dconf {
			s.restart(name, []string{name}, logger)
		}
	} else {
		_, err := getDconfItemFromArgs(args)
		if err != nil {
			logger.Error([]byte("get dconfItem error: " + err.Error()))
			return
		}

		s.restart(containerName, args, logger)
	}
}

func (s *RestartCommand) restart(containerName string, args []string, logger golog.ILogger) {
	logger.Notice([]byte("restart container: " + containerName))

	stopCmd := new(StopCommand)
	stopCmd.Run(args, logger)

	startCmd := new(StartCommand)
	startCmd.Run(args, logger)
}
