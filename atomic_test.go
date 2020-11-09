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
