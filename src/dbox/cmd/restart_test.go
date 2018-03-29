package cmd

import (
	"testing"
)

func TestRestart(t *testing.T) {
	cmd := new(RestartCommand)
	cmd.Run([]string{"redis"}, tlogger)
}
