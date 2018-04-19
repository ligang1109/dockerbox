package golog

import (
	"testing"
	"time"
)

func TestSimpleLogger(t *testing.T) {
	fw, _ := NewFileWriter("/tmp/test_simple_logger.log")
	logger, _ := NewSimpleLogger(fw, LEVEL_DEBUG, NewSimpleFormater())

	msg := []byte("test simple logger")

	testLogger(logger, msg)

	logger.Free()
}

func TestSimpleBufferLogger(t *testing.T) {
	InitBufferAutoFlushRoutine(1024, time.Second*7)

	fw, _ := NewFileWriter("/tmp/test_simple_buffer_logger.log")
	bw := NewBuffer(fw, 1024)
	logger, _ := NewSimpleLogger(bw, LEVEL_INFO, NewSimpleFormater())

	msg := []byte("test simple buffer logger")

	testLogger(logger, msg)

	logger.Free()
}
