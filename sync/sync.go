package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	// Current execution locks Counter
	// preventing other goroutines/calls from interrupting
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
