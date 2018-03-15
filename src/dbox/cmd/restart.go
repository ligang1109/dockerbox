package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"
)

type RestartCommand struct {
}

func (s *RestartCommand) Run(args []string, logger golog.ILogger) {
	containerKey, err := getContainerKeyFromArgs(args)
	if err != nil {
		logger.Error([]byte("get containerKey error: " + err.Error()))
		return
	}

	if containerKey == SPECIAL_CONTAINER_KEY_ALL {
		for name, _ := range dconf.Dconf {
			s.restart(name, []string{name}, logger)
		}
	} else {
		_, err := getDconfItemFromArgs(args)
		if err != nil {
			logger.Error([]byte("get dconfItem error: " + err.Error()))
			return
		}

		s.restart(containerKey, args, logger)
	}
}

func (s *RestartCommand) restart(containerKey string, args []string, logger golog.ILogger) {
	logger.Notice([]byte("restart container: " + containerKey))

	stopCmd := new(StopCommand)
	stopCmd.Run(args, logger)

	startCmd := new(StartCommand)
	startCmd.Run(args, logger)
}
