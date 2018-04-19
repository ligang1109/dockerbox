/**
* @file async.go
* @brief write msg in independency goroutine
* @author ligang
* @date 2016-07-19
 */

package golog

const (
	ASYNC_MSG_KIND_LOG          = 1
	ASYNC_MSG_KIND_FLUSH        = 2
	ASYNC_MSG_KIND_FREE_LOGGER  = 3
	ASYNC_MSG_KIND_FREE_ROUTINE = 4
)

type asyncMsg struct {
	kind   int
	logger ILogger

	level int
	msg   []byte
}

var alr *asyncLogRoutine

// must be called first
func InitAsyncLogRoutine(msgQueueLen int) {
	alr = &asyncLogRoutine{
		msgCh:  make(chan *asyncMsg, msgQueueLen),
		freeCh: make(chan int),
	}

	go alr.run()
}

func FreeAsyncLogRoutine() {
	alr.msgCh <- &asyncMsg{
		kind: ASYNC_MSG_KIND_FREE_ROUTINE,
	}
	<-alr.freeCh
}

/**
* @name async log routine
* @{ */

type asyncLogRoutine struct {
	msgCh  chan *asyncMsg
	freeCh chan int
}

func (a *asyncLogRoutine) run() {
	defer func() {
		a.freeCh <- 1
	}()

	for {
		am := <-a.msgCh
		free := a.processAsyncMsg(am)
		if free {
			return
		}
	}
}

func (a *asyncLogRoutine) processAsyncMsg(am *asyncMsg) bool {
	switch am.kind {
	case ASYNC_MSG_KIND_LOG:
		am.logger.Log(am.level, am.msg)
	case ASYNC_MSG_KIND_FLUSH:
		am.logger.Flush()
	case ASYNC_MSG_KIND_FREE_LOGGER:
		am.logger.Free()
	case ASYNC_MSG_KIND_FREE_ROUTINE:
		return true
	}

	return false
}

/**  @} */

/**
* @name async logger
* @{ */

type asyncLogger struct {
	logger ILogger
}

func NewAsyncLogger(logger ILogger) *asyncLogger {
	a := &asyncLogger{
		logger: logger,
	}

	return a
}

func (a *asyncLogger) Debug(msg []byte) {
	a.Log(LEVEL_DEBUG, msg)
}

func (a *asyncLogger) Info(msg []byte) {
	a.Log(LEVEL_INFO, msg)
}

func (a *asyncLogger) Notice(msg []byte) {
	a.Log(LEVEL_NOTICE, msg)
}

func (a *asyncLogger) Warning(msg []byte) {
	a.Log(LEVEL_WARNING, msg)
}

func (a *asyncLogger) Error(msg []byte) {
	a.Log(LEVEL_ERROR, msg)
}

func (a *asyncLogger) Critical(msg []byte) {
	a.Log(LEVEL_CRITICAL, msg)
}

func (a *asyncLogger) Alert(msg []byte) {
	a.Log(LEVEL_ALERT, msg)
}

func (a *asyncLogger) Emergency(msg []byte) {
	a.Log(LEVEL_EMERGENCY, msg)
}

func (a *asyncLogger) Log(level int, msg []byte) error {
	alr.msgCh <- &asyncMsg{
		kind:   ASYNC_MSG_KIND_LOG,
		logger: a.logger,

		msg:   msg,
		level: level,
	}

	return nil
}

func (a *asyncLogger) Flush() error {
	alr.msgCh <- &asyncMsg{
		kind:   ASYNC_MSG_KIND_FLUSH,
		logger: a.logger,
	}

	return nil
}

func (a *asyncLogger) Free() {
	alr.msgCh <- &asyncMsg{
		kind:   ASYNC_MSG_KIND_FREE_LOGGER,
		logger: a.logger,
	}
}

/**  @} */
