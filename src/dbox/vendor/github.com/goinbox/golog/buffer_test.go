package golog

import (
	//     "fmt"
	"testing"
	"time"
)

func TestBufferFileWriter(t *testing.T) {
	InitBufferAutoFlushRoutine(1024, time.Second*3)

	path := "/tmp/test_buffer.log"
	bufsize := 1024

	fw, _ := NewFileWriter(path)
	bw := NewBuffer(fw, bufsize)

	bw.Write([]byte("test file writer with buffer and time interval\n"))

	time.Sleep(time.Second * 5)
	bw.Free()

	FreeBuffers()
}
