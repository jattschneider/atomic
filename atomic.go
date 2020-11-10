package atomic

import (
	"sync"
	"sync/atomic"
)

// NewCounter creates a new Counter (zero value)
func NewCounter() *Counter {
	return &Counter{}
}

// Counter A uint64 value that may be updated atomically
type Counter struct {
	sync.Mutex
	count uint64
}

// Incr Increments by one the current value
func (c *Counter) Incr() {
	c.Lock()
	defer c.Unlock()
	atomic.AddUint64(&c.count, 1)
}

// Count Returns the value as a uint64
func (c *Counter) Count() uint64 {
	c.Lock()
	defer c.Unlock()
	return c.count
}

// Reset resets to zero value
func (c *Counter) Reset() {
	c.Lock()
	defer c.Unlock()
	c.count = 0
}

// NewInt creates and initializes a new Int
func NewInt(val int64) *Int {
	return &Int{val: val}
}

// Int An int64 value that may be updated atomically
type Int struct {
	sync.Mutex
	val int64
}

// Incr Increments by one the current value
func (i *Int) Incr() {
	i.Lock()
	defer i.Unlock()
	atomic.AddInt64(&i.val, 1)
}

// Decr Decrements by one the current value
func (i *Int) Decr() {
	i.Lock()
	defer i.Unlock()
	atomic.AddInt64(&i.val, -1)
}

// Add adds the given detla value
func (i *Int) Add(delta int64) {
	i.Lock()
	defer i.Unlock()
	atomic.AddInt64(&i.val, delta)
}

// Value Returns the value as an int64
func (i *Int) Value() int64 {
	i.Lock()
	defer i.Unlock()
	return i.val
}

// Reset resets to zero value
func (i *Int) Reset() {
	i.Lock()
	defer i.Unlock()
	i.val = 0
}

// NewBool creates and initializes a new Bool
func NewBool(val bool) *Bool {
	return &Bool{val: val}
}

// Bool A boolean that may be updated atomically
type Bool struct {
	sync.Mutex
	val bool
}

// Value Returns the value as a boolean
func (b *Bool) Value() bool {
	b.Lock()
	defer b.Unlock()
	return b.val
}

// Set sets to the given value
func (b *Bool) Set(v bool) {
	b.Lock()
	defer b.Unlock()
	b.val = v
}

// Flip flips the current boolean value (value = NOT value)
func (b *Bool) Flip() {
	b.Lock()
	defer b.Unlock()
	b.val = !b.val
}

// Reset resets to zero value
func (b *Bool) Reset() {
	b.Lock()
	defer b.Unlock()
	b.val = false
}
