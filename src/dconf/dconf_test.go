package dconf

import (
	"os"
	"testing"
)

//export DCONF_PATH=/Users/gibsonli/devspace/personal/dockerbox/dconf.json.demo ;go test -v
func TestRconf(t *testing.T) {
	dconfPath := os.Getenv("DCONF_PATH")

	err := Init(dconfPath)
	if err != nil {
		t.Error(err)
	}

	for name, item := range Dconf {
		t.Log(name, item.ContainerName, item.Exec)
	}
}
