package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestSemaphoreBasic(t *testing.T) {
	const count = 5
	s := make(Semaphore, count)

	var doneCounter int32

	for i := 0; i < count; i++ {
		go func() {
			time.Sleep(10 * time.Millisecond)
			atomic.AddInt32(&doneCounter, 1)
			s.Increment(1)
		}()
	}

	s.Decrement(count)

	if atomic.LoadInt32(&doneCounter) != count {
		t.Errorf("ожидалось %d завершённых горутин, получено %d",
			count, doneCounter)
	}
}

func TestSemaphoreZero(t *testing.T) {
	s := make(Semaphore, 10)

	s.Decrement(0)
}

func TestSemaphoreParallel(t *testing.T) {
	s := make(Semaphore, 100)

	const n = 50
	var done int32

	for i := 0; i < n; i++ {
		go func() {
			time.Sleep(time.Millisecond)
			atomic.AddInt32(&done, 1)
			s.Increment(1)
		}()
	}

	s.Decrement(n)

	if atomic.LoadInt32(&done) != n {
		t.Errorf("expected %d done, got %d", n, done)
	}
}
