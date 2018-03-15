package cmd

import (
	"testing"
)

func TestStart(t *testing.T) {
	cmd := new(StartCommand)
	cmd.Run([]string{"mysql"}, tlogger)
	cmd.Run([]string{SPECIAL_CONTAINER_KEY_ALL}, tlogger)
}
