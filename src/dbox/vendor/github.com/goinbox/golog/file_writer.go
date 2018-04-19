/**
* @file file.go
* @brief writer msg to file
* @author ligang
* @date 2016-02-03
 */

package golog

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/goinbox/gomisc"
)

/**
* @name file writer
* @{ */

type FileWriter struct {
	path        string
	lock        *sync.Mutex
	closeOnFree bool

	*os.File
}

func NewFileWriter(path string) (*FileWriter, error) {
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}

	return &FileWriter{
		path:        path,
		lock:        new(sync.Mutex),
		closeOnFree: false,

		File: file,
	}, nil
}

func (f *FileWriter) CloseOnFree(closeOneFree bool) *FileWriter {
	f.closeOnFree = closeOneFree

	return f
}

func (f *FileWriter) Write(msg []byte) (int, error) {
	// file may be deleted when doing logrotate
	if !gomisc.FileExist(f.path) {
		f.Close()
		f.File, _ = openFile(f.path)
	}

	f.lock.Lock()
	n, err := f.File.Write(msg)
	f.lock.Unlock()

	return n, err
}

func (f *FileWriter) Flush() error {
	return nil
}

func (f *FileWriter) Free() {
	if f.closeOnFree {
		f.File.Close()
	}
}

func (f *FileWriter) ForceFree() {
	f.File.Close()
}

func openFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}

/**  @} */

/**
* @name file writer with split
* @{ */

const (
	SPLIT_BY_DAY  = 1
	SPLIT_BY_HOUR = 2
)

type FileWithSplitWriter struct {
	path   string
	split  int
	suffix string

	*FileWriter
}

func NewFileWriterWithSplit(path string, split int) (*FileWithSplitWriter, error) {
	suffix := makeFileSuffix(split)
	if suffix == "" {
		return nil, errors.New("Split not support")
	}

	fw, err := NewFileWriter(path + "." + suffix)
	if err != nil {
		return nil, err
	}

	f := &FileWithSplitWriter{
		path:   path,
		split:  split,
		suffix: suffix,

		FileWriter: fw,
	}

	return f, nil
}

func (f *FileWithSplitWriter) Write(msg []byte) (int, error) {
	suffix := makeFileSuffix(f.split)

	//need split
	if suffix != f.suffix {
		f.Free()
		f.FileWriter, _ = NewFileWriter(f.path + "." + suffix)
		f.suffix = suffix
	}

	return f.File.Write(msg)
}

func makeFileSuffix(split int) string {
	switch split {
	case SPLIT_BY_DAY:
		return time.Now().Format(gomisc.TIME_FMT_STR_YEAR + gomisc.TIME_FMT_STR_MONTH + gomisc.TIME_FMT_STR_DAY)
	case SPLIT_BY_HOUR:
		return time.Now().Format(gomisc.TIME_FMT_STR_YEAR + gomisc.TIME_FMT_STR_MONTH + gomisc.TIME_FMT_STR_DAY + gomisc.TIME_FMT_STR_HOUR)
	default:
		return ""
	}
}

/**  @} */
