/**
* @file base.go
* @brief format msg before send to writer
* @author ligang
* @date 2016-07-12
 */

package golog

import (
	"time"

	"github.com/goinbox/gomisc"
)

type simpleFormater struct {
}

func NewSimpleFormater() *simpleFormater {
	return new(simpleFormater)
}

func (s *simpleFormater) Format(level int, msg []byte) []byte {
	lm, ok := logLevels[level]
	if !ok {
		lm = []byte("-")
	}
	return gomisc.AppendBytes(
		[]byte("["),
		lm,
		[]byte("]\t"),
		[]byte("["),
		[]byte(time.Now().Format(gomisc.TimeGeneralLayout())),
		[]byte("]\t"),
		msg,
		[]byte("\n"),
	)
}
