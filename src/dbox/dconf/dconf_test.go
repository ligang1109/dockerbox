package dconf

import (
	"os"
	"testing"
)

func TestRconf(t *testing.T) {
	prjHome := os.Getenv("GOPATH")
	dconfPath := prjHome + "/dconf.json.demo"

	err := Init(dconfPath)
	if err != nil {
		t.Error(err)
	}

	for name, item := range Dconf {
		t.Log(name, item)
	}
}
