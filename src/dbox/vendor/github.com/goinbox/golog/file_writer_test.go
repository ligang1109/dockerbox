package golog

import (
	//     "fmt"
	"testing"
)

func TestFileWriter(t *testing.T) {
	path := "/tmp/test.log"

	writer, _ := NewFileWriter(path)

	writer.Write([]byte("test file writer\n"))
}

func TestFileWriterWithSplit(t *testing.T) {
	path := "/tmp/test.log"

	writer, _ := NewFileWriterWithSplit(path, SPLIT_BY_DAY)
	writer.Write([]byte("test file writer with split by day\n"))

	writer, _ = NewFileWriterWithSplit(path, SPLIT_BY_HOUR)
	writer.Write([]byte("test file writer with split by hour\n"))
}
