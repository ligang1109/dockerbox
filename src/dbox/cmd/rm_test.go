package cmd

import "testing"

func TestRm(t *testing.T) {
	cmd := new(RmCommand)
	cmd.Run([]string{"redis"}, tlogger)
}
