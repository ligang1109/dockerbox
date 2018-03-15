package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"

	"errors"
	"strings"
)

const (
	SPECIAL_CONTAINER_KEY_ALL = "all"
)

type ICommand interface {
	Run(args []string, logger golog.ILogger)
}

func getContainerKeyFromArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("do not has containerKey arg")
	}

	return strings.TrimSpace(args[0]), nil
}

func getDconfItemFromArgs(args []string) (*dconf.DconfItem, error) {
	containerKey, err := getContainerKeyFromArgs(args)
	if err != nil {
		return nil, err
	}

	dconfItem, ok := dconf.Dconf[containerKey]
	if !ok {
		return nil, errors.New("containerKey: " + containerKey + " not in dconf")
	}

	return dconfItem, nil
}
