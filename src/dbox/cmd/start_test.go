package cmd

import (
	"testing"
)

func TestStart(t *testing.T) {
	cmd := new(StartCommand)
	cmd.Run([]string{"redis"}, tlogger)
}
