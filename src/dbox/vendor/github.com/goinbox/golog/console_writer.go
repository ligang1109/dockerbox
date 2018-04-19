package golog

import (
	"os"
	"sync"
)

type ConsoleWriter struct {
	lock *sync.Mutex

	*os.File
}

func NewStdoutWriter() *ConsoleWriter {
	return &ConsoleWriter{
		lock: new(sync.Mutex),

		File: os.Stdout,
	}
}

func NewStderrWriter() *ConsoleWriter {
	return &ConsoleWriter{
		lock: new(sync.Mutex),

		File: os.Stderr,
	}
}

func (c *ConsoleWriter) Write(msg []byte) (int, error) {
	c.lock.Lock()
	n, err := c.File.Write(msg)
	c.lock.Unlock()

	return n, err
}

func (c *ConsoleWriter) Flush() error {
	return nil
}

func (c *ConsoleWriter) Free() {
}
