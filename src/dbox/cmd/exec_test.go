package cmd

import (
	"testing"
)

func TestExec(t *testing.T) {
	cmd := new(ExecCommand)
	cmd.Run([]string{"mysql", "mysql", "-uroot", "-p123"}, tlogger)
}
