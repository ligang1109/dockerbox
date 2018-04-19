/**
* @file logger.go
* @author ligang
* @date 2016-02-04
 */

package golog

import (
	"errors"
	"sync"
)

type simpleLogger struct {
	globalLevel  int
	w            IWriter
	levelWriters map[int]IWriter
	formater     IFormater

	lock *sync.Mutex
}

func NewSimpleLogger(writer IWriter, globalLevel int, formater IFormater) (*simpleLogger, error) {
	_, ok := logLevels[globalLevel]
	if !ok {
		return nil, errors.New("Global level not exists")
	}

	s := &simpleLogger{
		globalLevel:  globalLevel,
		w:            writer,
		levelWriters: make(map[int]IWriter),

		lock: new(sync.Mutex),
	}

	noopWriter := new(NoopWriter)
	for level, _ := range logLevels {
		if level < globalLevel {
			s.levelWriters[level] = noopWriter
		} else {
			s.levelWriters[level] = s.w
		}
	}

	if formater == nil {
		formater = new(NoopFormater)
	}
	s.formater = formater

	return s, nil
}

func (s *simpleLogger) Debug(msg []byte) {
	s.Log(LEVEL_DEBUG, msg)
}

func (s *simpleLogger) Info(msg []byte) {
	s.Log(LEVEL_INFO, msg)
}

func (s *simpleLogger) Notice(msg []byte) {
	s.Log(LEVEL_NOTICE, msg)
}

func (s *simpleLogger) Warning(msg []byte) {
	s.Log(LEVEL_WARNING, msg)
}

func (s *simpleLogger) Error(msg []byte) {
	s.Log(LEVEL_ERROR, msg)
}

func (s *simpleLogger) Critical(msg []byte) {
	s.Log(LEVEL_CRITICAL, msg)
}

func (s *simpleLogger) Alert(msg []byte) {
	s.Log(LEVEL_ALERT, msg)
}

func (s *simpleLogger) Emergency(msg []byte) {
	s.Log(LEVEL_EMERGENCY, msg)
}

func (s *simpleLogger) Log(level int, msg []byte) error {
	writer, ok := s.levelWriters[level]
	if !ok {
		return errors.New("Level not exists")
	}

	msg = s.formater.Format(level, msg)

	s.lock.Lock()
	writer.Write(msg)
	s.lock.Unlock()

	return nil
}

func (s *simpleLogger) Flush() error {
	return s.w.Flush()
}

func (s *simpleLogger) Free() {
	s.w.Free()
}
