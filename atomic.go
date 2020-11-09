package atomic

import (
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	sync.Mutex
	count uint64
}

func (a *AtomicCounter) Incr() {
	a.Lock()
	defer a.Unlock()
	atomic.AddUint64(&a.count, 1)
}

func (a *AtomicCounter) Count() uint64 {
	a.Lock()
	defer a.Unlock()
	return a.count
}

type AtomicInt struct {
	sync.Mutex
	val int64
}

func (a *AtomicInt) Incr() {
	a.Lock()
	defer a.Unlock()
	atomic.AddInt64(&a.val, 1)
}

func (a *AtomicInt) Decr() {
	a.Lock()
	defer a.Unlock()
	atomic.AddInt64(&a.val, -1)
}

func (a *AtomicInt) Add(delta int64) {
	a.Lock()
	defer a.Unlock()
	atomic.AddInt64(&a.val, delta)
}

func (a *AtomicInt) Value() int64 {
	a.Lock()
	defer a.Unlock()
	return a.val
}

type AtomicBool struct {
	sync.Mutex
	val bool
}

func (a *AtomicBool) Value() bool {
	a.Lock()
	defer a.Unlock()
	return a.val
}

func (a *AtomicBool) Set(v bool) {
	a.Lock()
	defer a.Unlock()
	a.val = v
}

func (a *AtomicBool) Flip() {
	a.Lock()
	defer a.Unlock()
	a.val = !a.val
}
