package dconf

import (
	"github.com/goinbox/gomisc"
)

type ExecItem struct {
	ShellCmd string `json:"shell_cmd"`
	Cwd      bool   `json:"cwd"`
	PreCmd   string `json:"pre_cmd"`
}

type StartItem struct {
	PreStart []string `json:"pre_start"`
}

type DconfItem struct {
	ContainerName string     `json:"container_name"`
	Exec          *ExecItem  `json:"exec"`
	Start         *StartItem `json:"start"`
}

var Dconf map[string]*DconfItem = make(map[string]*DconfItem)

func Init(dconfPath string) error {
	err := gomisc.ParseJsonFile(dconfPath, &Dconf)
	if err != nil {
		return err
	}

	return nil
}
