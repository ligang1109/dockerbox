package cmd

import "testing"

func TestRmi(t *testing.T) {
	cmd := new(RmiCommand)
	cmd.Run([]string{"redis"}, tlogger)
}
