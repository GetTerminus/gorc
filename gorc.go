package gorc

import (
	"sync"
	"time"
)

type Gorc struct {
	count      int
	waitMillis time.Duration
	sync.Mutex
}

// Inc increases the counter by one.
func (g *Gorc) Inc() {
	g.Lock()
	defer g.Unlock()

	g.count++
}

// IncBy increases the counter by b.
func (g *Gorc) IncBy(b int) {
	g.Lock()
	defer g.Unlock()

	g.count += b
}

// Dec decreases the counter by one.
func (g *Gorc) Dec() {
	g.Lock()
	defer g.Unlock()

	g.count--
}

// DecBy decreases the counter by b.
func (g *Gorc) DecBy(b int) {
	g.Lock()
	defer g.Unlock()
	
	g.count -= b
}

// GetCount returns an integer holding the count.
func (g *Gorc) Get() int {
	g.Lock()
	defer g.Unlock()
	
	return int(g.count)
}

// SetWaitMillis sets the time in milliseconds the Wait function
// waits between checking the count against the given integer.
func (g *Gorc) SetWaitMillis(w int) {
	g.Lock()
	defer g.Unlock()

	g.waitMillis = time.Duration(w) * time.Millisecond
}

// Init initializes a new Gorc instance
func (g *Gorc) Init() {
	g.Lock()
	defer g.Unlock()

	g.count = 0
	g.waitMillis = 100 * time.Millisecond
}

// WaitLow will return as soon as the Gorc counter falls below w.
// e.g. wait until all but w goroutines are stopped.
func (g *Gorc) WaitLow(w int) {
	for g.count >= w {
		time.Sleep(g.waitMillis)
	}
	return
}

// WaitHigh will return as soon as the Gorc counter goes above w.
// e.g. wait until at least w goroutines are started.
func (g *Gorc) WaitHigh(w int) {
	for g.count <= w {
		time.Sleep(g.waitMillis)
	}
	return
}
