package cmd

import (
	"testing"
)

func TestRestart(t *testing.T) {
	cmd := new(RestartCommand)
	//cmd.Run([]string{"mysql"}, tlogger)
	cmd.Run([]string{SPECIAL_CONTAINER_NAME_ALL}, tlogger)
}
