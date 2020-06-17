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

type newCommandFunc func() ICommand

var commandTable = make(map[string]newCommandFunc)

func register(name string, ncf newCommandFunc) {
	commandTable[name] = ncf
}

func NewCommandByName(name string) ICommand {
	ncf, ok := commandTable[name]
	if !ok {
		return nil
	}

	return ncf()
}

func containerKeyFromArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("do not has containerKey arg")
	}

	return strings.TrimSpace(args[0]), nil
}

func dconfItemFromContainerKey(containerKey string) (*dconf.DconfItem, error) {
	dconfItem, ok := dconf.Dconf[containerKey]
	if !ok {
		return nil, errors.New("containerKey: " + containerKey + " not in dconf")
	}

	return dconfItem, nil
}

func dconfItemFromArgs(args []string) (*dconf.DconfItem, error) {
	containerKey, err := containerKeyFromArgs(args)
	if err != nil {
		return nil, err
	}

	return dconfItemFromContainerKey(containerKey)
}
