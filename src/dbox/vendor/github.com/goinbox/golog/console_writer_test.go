package golog

import "testing"

func TestConsoleWriter(t *testing.T) {
	writer := NewStdoutWriter()
	writer.Write([]byte("test stdout console writer\n"))

	writer = NewStderrWriter()
	writer.Write([]byte("test stderr console writer\n"))
}
