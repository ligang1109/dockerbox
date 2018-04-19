package golog

import (
	"time"

	"github.com/goinbox/color"
	"github.com/goinbox/gomisc"
)

type colorFunc func(msg []byte) []byte

type consoleFormater struct {
	levelColorFuncs map[int]colorFunc
}

func NewConsoleFormater() *consoleFormater {
	c := &consoleFormater{
		levelColorFuncs: make(map[int]colorFunc),
	}

	c.levelColorFuncs[LEVEL_DEBUG] = color.Yellow
	c.levelColorFuncs[LEVEL_INFO] = color.Blue
	c.levelColorFuncs[LEVEL_NOTICE] = color.Cyan
	c.levelColorFuncs[LEVEL_WARNING] = color.Maganta
	c.levelColorFuncs[LEVEL_ERROR] = color.Red
	c.levelColorFuncs[LEVEL_CRITICAL] = color.Black
	c.levelColorFuncs[LEVEL_ALERT] = color.White
	c.levelColorFuncs[LEVEL_EMERGENCY] = color.Green

	return c
}

func (c *consoleFormater) SetColor(level int, cf colorFunc) *consoleFormater {
	c.levelColorFuncs[level] = cf

	return c
}

func (c *consoleFormater) Format(level int, msg []byte) []byte {
	lm, ok := logLevels[level]
	if !ok {
		lm = []byte("-")
	}

	msg = gomisc.AppendBytes(
		[]byte("["),
		lm,
		[]byte("]\t"),
		[]byte("["),
		[]byte(time.Now().Format(gomisc.TimeGeneralLayout())),
		[]byte("]\t"),
		msg,
	)
	msg = c.levelColorFuncs[level](msg)

	return gomisc.AppendBytes(msg, []byte("\n"))
}
