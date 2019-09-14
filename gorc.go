package gorc

import (
	"sync/atomic"
	"time"
)

// Gorc is used to keep track of goroutines running.
type Gorc interface {
	IncBy(b int)
	Inc()

	DecBy(b int)
	Dec()

	GetCount() int

	SetWaitMillis(w int)
	WaitLow(w int)
	WaitHigh(w int)
}

// NewGorc makes you a new Gorc
func NewGorc() Gorc {
	g := &gorc{}
	var v int32 = 0
	g.count = &v
	g.waitMillis = 100 * time.Millisecond
	return g
}

type gorc struct {
	count      *int32
	waitMillis time.Duration
}

// Inc increases the counter by one.
func (g *gorc) Inc() {
	g.IncBy(1)
}

// IncBy increases the counter by b.
func (g *gorc) IncBy(b int) {
	atomic.AddInt32(g.count, int32(b))
}

// Dec decreases the counter by one.
func (g *gorc) Dec() {
	g.DecBy(1)
}

// DecBy decreases the counter by b.
func (g *gorc) DecBy(b int) {
	atomic.AddInt32(g.count, int32(b)*-1)
}

// GetCount returns an integer holding the count.
func (g gorc) GetCount() int {
	return int(atomic.LoadInt32(g.count))
}

// SetWaitMillis sets the time in milliseconds the Wait function
// waits between checking the count against the given integer.
func (g *gorc) SetWaitMillis(w int) {
	g.waitMillis = time.Duration(w) * time.Millisecond
}

// WaitLow will return as soon as the Gorc counter falls below w.
// e.g. wait until all but w goroutines are stopped.
func (g gorc) WaitLow(w int) {
	l := int32(w)
	for atomic.LoadInt32(g.count) >= l {
		time.Sleep(g.waitMillis)
	}
	return
}

// WaitHigh will return as soon as the Gorc counter goes above w.
// e.g. wait until at least w goroutines are started.
func (g gorc) WaitHigh(w int) {
	h := int32(w)
	for atomic.LoadInt32(g.count) <= h {
		time.Sleep(g.waitMillis)
	}
	return
}
