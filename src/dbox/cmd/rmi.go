package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"
	"github.com/goinbox/shell"

	"regexp"
)

const (
	CMD_NAME_RMI = "rmi"
)

func init() {
	Register(CMD_NAME_RMI, newRmiCommand)
}

func newRmiCommand() ICommand {
	return new(RmiCommand)
}

type RmiCommand struct {
}

func (r *RmiCommand) Run(args []string, logger golog.ILogger) {
	containerKey, err := ContainerKeyFromArgs(args)
	if err != nil {
		logger.Error([]byte("get containerKey error: " + err.Error()))
		return
	}

	if containerKey == SPECIAL_CONTAINER_KEY_ALL {
		for _, item := range dconf.Dconf {
			r.rmi(item, logger)
		}
	} else {
		item, err := DconfItemFromArgs(args)
		if err != nil {
			logger.Error([]byte("get dconfItem error: " + err.Error()))
			return
		}

		r.rmi(item, logger)
	}
}

var regex *regexp.Regexp = regexp.MustCompile(`"Image": "sha256:([a-z0-9]+)",`)

func (r *RmiCommand) rmi(dconfItem *dconf.DconfItem, logger golog.ILogger) {
	cmd := "docker inspect " + dconfItem.ContainerName

	logger.Debug([]byte("cmd: " + cmd))

	sr := shell.RunCmd(cmd)
	if !sr.Ok {
		logger.Error([]byte("inspect container: " + dconfItem.ContainerName + " error: " + string(sr.Output)))
		return
	}

	matchBytes := regex.FindSubmatch(sr.Output)
	if len(matchBytes) < 2 {
		logger.Error([]byte("inspect container: " + dconfItem.ContainerName + " error: " + string(sr.Output)))
		return
	}

	imageId := string(matchBytes[1])
	cmd = "docker rm -f " + dconfItem.ContainerName

	logger.Warning([]byte("rm container: " + dconfItem.ContainerName))
	logger.Debug([]byte("cmd: " + cmd))

	shell.RunCmdBindTerminal(cmd)

	cmd = "docker rmi -f " + imageId

	logger.Warning([]byte("rmi image: " + dconfItem.ContainerName))
	logger.Debug([]byte("cmd: " + cmd))

	shell.RunCmdBindTerminal(cmd)
}
