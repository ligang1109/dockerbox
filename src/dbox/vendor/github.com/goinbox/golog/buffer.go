/**
* @file buffer.go
* @brief writer with buffer
* @author ligang
* @date 2016-02-04
 */

package golog

import (
	"bufio"
	"sync"
	"time"
)

var bfr *bufFlushRoutine
var bfrInitMutex sync.Mutex

// must be called first
func InitBufferAutoFlushRoutine(maxBufNum int, timeInterval time.Duration) {
	bfrInitMutex.Lock()

	if bfr == nil {
		bfr = &bufFlushRoutine{
			buffers: make(map[uint64]*buffer),

			bufAddCh: make(chan *buffer, maxBufNum),
			bufDelCh: make(chan *buffer, maxBufNum),
			freeCh:   make(chan int),
		}

		go bfr.run(timeInterval)
	}

	bfrInitMutex.Unlock()
}

func FreeBuffers() {
	bfr.freeCh <- 1
	<-bfr.freeCh
	bfr = nil
}

/**
* @name auto flush routine
* @{ */

type bufFlushRoutine struct {
	curIndex uint64
	buffers  map[uint64]*buffer

	bufAddCh chan *buffer
	bufDelCh chan *buffer
	freeCh   chan int
}

func (b *bufFlushRoutine) addBuffer(buf *buffer) {
	b.bufAddCh <- buf
}

func (b *bufFlushRoutine) delBuffer(buf *buffer) {
	b.bufDelCh <- buf
}

func (b *bufFlushRoutine) flushAll() {
	for index, buf := range b.buffers {
		if buf == nil || buf.buf == nil {
			delete(b.buffers, index)
		} else {
			buf.Flush()
		}
	}
}

func (b *bufFlushRoutine) run(timeInterval time.Duration) {
	ticker := time.NewTicker(timeInterval)

	for {
		select {
		case buf, _ := <-b.bufAddCh:
			buf.index = b.curIndex
			b.buffers[b.curIndex] = buf
			b.curIndex++
		case buf, _ := <-b.bufDelCh:
			delete(b.buffers, buf.index)
			buf.buf = nil
		case <-ticker.C:
			b.flushAll()
		case <-b.freeCh:
			b.flushAll()
			b.freeCh <- 1
			return
		}
	}
}

/**  @} */

/**
* @name buffer
* @{ */

type buffer struct {
	w   IWriter
	buf *bufio.Writer

	lock  *sync.Mutex
	index uint64
}

func NewBuffer(w IWriter, bufsize int) *buffer {
	b := &buffer{
		w:   w,
		buf: bufio.NewWriterSize(w, bufsize),

		lock: new(sync.Mutex),
	}

	bfr.addBuffer(b)

	return b
}

func (b *buffer) Write(p []byte) (int, error) {
	b.lock.Lock()
	n, err := b.buf.Write(p)
	b.lock.Unlock()

	return n, err
}

func (b *buffer) Flush() error {
	var err error

	b.lock.Lock()
	if b.buf != nil {
		err = b.buf.Flush()
	}
	b.lock.Unlock()

	return err
}

func (b *buffer) Free() {
	b.Flush()
	b.w.Free()

	bfr.delBuffer(b)
}

/**  @} */
