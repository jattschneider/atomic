package atomic

import (
	"sync"
	"testing"
)

func TestAtomicCounter(t *testing.T) {
	c := &AtomicCounter{}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			c.Incr()
			wg.Done()
		}()
	}
	wg.Wait()

	c.Incr()

	if c.Count() != 101 {
		t.Fail()
	}
}

func TestAtomicInt(t *testing.T) {
	ai := &AtomicInt{}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			ai.Incr()
			ai.Decr()
			wg.Done()
		}()
	}
	wg.Wait()

	ai.Add(101)
	ai.Add(-100)

	if ai.Value() != 1 {
		t.Fail()
	}

	ai.Decr()

	if ai.Value() != 0 {
		t.Fail()
	}
}

func TestAtomicBool(t *testing.T) {
	b := &AtomicBool{}
	b.Set(true)

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			b.Flip()
			wg.Done()
		}()
	}
	wg.Wait()

	if !b.Value() {
		t.Fail()
	}

	b.Flip()

	if b.Value() {
		t.Fail()
	}
}
