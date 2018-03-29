package cmd

import (
	"testing"
)

func TestStop(t *testing.T) {
	cmd := new(StopCommand)
	cmd.Run([]string{"redis"}, tlogger)
}
