package cmd

import (
	"testing"
)

func TestExec(t *testing.T) {
	cmd := new(ExecCommand)
	cmd.Run([]string{"nginx", "nginx", "-v"}, tlogger)
}
