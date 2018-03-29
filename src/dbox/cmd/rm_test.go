package cmd

import "testing"

func TestRmCommand(t *testing.T) {
	cmd := new(RmCommand)
	cmd.Run([]string{"redis"}, tlogger)
}
