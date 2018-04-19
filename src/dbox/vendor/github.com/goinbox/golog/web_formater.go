package golog

import (
	"time"

	"github.com/goinbox/gomisc"
)

type webFormater struct {
	logId []byte
	ip    []byte
}

func NewWebFormater(logId, ip []byte) *webFormater {
	return &webFormater{
		logId: logId[:],
		ip:    ip[:],
	}
}

func (w *webFormater) Format(level int, msg []byte) []byte {
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
		w.ip,
		[]byte("\t"),
		w.logId,
		[]byte("\t"),
		msg,
		[]byte("\n"),
	)
}
