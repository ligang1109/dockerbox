/**
* @file logger.go
* @author ligang
* @date 2016-02-04
 */

package golog

import "io"

const (
	LEVEL_DEBUG     = 1
	LEVEL_INFO      = 2
	LEVEL_NOTICE    = 3
	LEVEL_WARNING   = 4
	LEVEL_ERROR     = 5
	LEVEL_CRITICAL  = 6
	LEVEL_ALERT     = 7
	LEVEL_EMERGENCY = 8
)

var logLevels map[int][]byte = map[int][]byte{
	LEVEL_DEBUG:     []byte("debug"),
	LEVEL_INFO:      []byte("info"),
	LEVEL_NOTICE:    []byte("notice"),
	LEVEL_WARNING:   []byte("warning"),
	LEVEL_ERROR:     []byte("error"),
	LEVEL_CRITICAL:  []byte("critical"),
	LEVEL_ALERT:     []byte("alert"),
	LEVEL_EMERGENCY: []byte("emergency"),
}

type ILogger interface {
	Debug(msg []byte)
	Info(msg []byte)
	Notice(msg []byte)
	Warning(msg []byte)
	Error(msg []byte)
	Critical(msg []byte)
	Alert(msg []byte)
	Emergency(msg []byte)

	Log(level int, msg []byte) error

	Flush() error
	Free()
}

type IFormater interface {
	Format(level int, msg []byte) []byte
}

type IWriter interface {
	io.Writer

	Flush() error
	Free()
}
