package dconf

import (
	"github.com/goinbox/gomisc"
)

type DconfItem struct {
	ContainerName string `json:"container_name"`
}

var Dconf map[string]*DconfItem = make(map[string]*DconfItem)

func Init(dconfPath string) error {
	err := gomisc.ParseJsonFile(dconfPath, &Dconf)
	if err != nil {
		return err
	}

	return nil
}
