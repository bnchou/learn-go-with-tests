package sync

import (
	"sync"
	"testing"
)

func NewCounter() *Counter {
	return &Counter{}
}

func assertCounter(t *testing.T, counter *Counter, i int) {
	t.Helper()
	if counter.Value() != i {
		t.Errorf("got %d, want %d", counter.Value(), i)
	}
}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
