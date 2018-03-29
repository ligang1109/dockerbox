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

type NewCommandFunc func() ICommand

var commandTable = make(map[string]NewCommandFunc)

func Register(name string, ncf NewCommandFunc) {
	commandTable[name] = ncf
}

func NewCommandByName(name string) ICommand {
	ncf, ok := commandTable[name]
	if !ok {
		return nil
	}

	return ncf()
}

func ContainerKeyFromArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("do not has containerKey arg")
	}

	return strings.TrimSpace(args[0]), nil
}

func DconfItemFromArgs(args []string) (*dconf.DconfItem, error) {
	containerKey, err := ContainerKeyFromArgs(args)
	if err != nil {
		return nil, err
	}

	dconfItem, ok := dconf.Dconf[containerKey]
	if !ok {
		return nil, errors.New("containerKey: " + containerKey + " not in dconf")
	}

	return dconfItem, nil
}
