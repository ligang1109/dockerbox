package cmd

import (
	"testing"
)

func TestAttach(t *testing.T) {
	cmd := new(AttachCommand)
	cmd.Run([]string{"mysql"}, tlogger)
}
