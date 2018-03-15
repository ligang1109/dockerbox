package cmd

import (
	"testing"
)

func TestStop(t *testing.T) {
	cmd := new(StopCommand)
	cmd.Run([]string{"mysql"}, tlogger)
	cmd.Run([]string{SPECIAL_CONTAINER_KEY_ALL}, tlogger)
}
