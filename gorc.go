package gorc

import (
	"sync/atomic"
	"time"
)

// Gorc is used to keep track of goroutines running.
type Gorc struct {
	count      *int32
	waitMillis time.Duration
}

// Inc increases the counter by one.
func (g *Gorc) Inc() {
	atomic.AddInt32(g.count, 1)
}

// IncBy increases the counter by b.
func (g *Gorc) IncBy(b int) {
	atomic.AddInt32(g.count, int32(b))
}

// Dec decreases the counter by one.
func (g *Gorc) Dec() {
	atomic.AddInt32(g.count, -1)
}

// DecBy decreases the counter by b.
func (g *Gorc) DecBy(b int) {
	atomic.AddInt32(g.count, int32(b)*-1)
}

// GetCount returns an integer holding the count.
func (g *Gorc) GetCount() int {
	return int(atomic.LoadInt32(g.count))
}

// SetWaitMillis sets the time in milliseconds the Wait function
// waits between checking the count against the given integer.
func (g *Gorc) SetWaitMillis(w int) {
	g.waitMillis = time.Duration(w) * time.Millisecond
}

// Init initializes a new Gorc instance
func (g *Gorc) Init() {
	atomic.StoreInt32(g.count, 0)
	g.waitMillis = 100 * time.Millisecond
}

// WaitLow will return as soon as the Gorc counter falls below w.
// e.g. wait until all but w goroutines are stopped.
func (g *Gorc) WaitLow(w int) {
	l := int32(w)
	for atomic.LoadInt32(g.count) >= l {
		time.Sleep(g.waitMillis)
	}
	return
}

// WaitHigh will return as soon as the Gorc counter goes above w.
// e.g. wait until at least w goroutines are started.
func (g *Gorc) WaitHigh(w int) {
	h := int32(w)
	for atomic.LoadInt32(g.count) <= h {
		time.Sleep(g.waitMillis)
	}
	return
}
