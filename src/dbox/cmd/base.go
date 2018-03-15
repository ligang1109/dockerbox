package cmd

import (
	"dbox/dconf"

	"github.com/goinbox/golog"
)

type ICommand interface {
	Run(dconfItem *dconf.DconfItem, args []string, logger golog.ILogger)
}
