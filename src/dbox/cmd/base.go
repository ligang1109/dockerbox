package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"

	"errors"
	"strings"
)

const (
	SPECIAL_CONTAINER_NAME_ALL = "all"
)

type ICommand interface {
	Run(args []string, logger golog.ILogger)
}

func getContainerNameFromArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("do not has containerName arg")
	}

	return strings.TrimSpace(args[0]), nil
}

func getDconfItemFromArgs(args []string) (*dconf.DconfItem, error) {
	containerName, err := getContainerNameFromArgs(args)
	if err != nil {
		return nil, err
	}

	dconfItem, ok := dconf.Dconf[containerName]
	if !ok {
		return nil, errors.New("containerName: " + containerName + " not in dconf")
	}

	return dconfItem, nil
}
