package gorc

import (
	"sync/atomic"
	"time"
)

// Gorc is used to keep track of goroutines running.
type Gorc interface {
	// IncBy increases the counter by b.
	IncBy(b int)
	// Inc increases the counter by one.
	Inc()

	// DecBy decreases the counter by b.
	DecBy(b int)
	// Dec decreases the counter by one.
	Dec()

	// GetCount returns an integer holding the count.
	GetCount() int

	// WaitLow will return as soon as the Gorc counter falls below w.
	// e.g. wait until all but w goroutines are stopped.
	WaitLow(w int)

	// WaitHigh will return as soon as the Gorc counter goes above w.
	// e.g. wait until at least w-1 goroutines are started.
	WaitHigh(w int)
}

// NewGorc makes you a new Gorc
func NewGorc(waitDwellMilliseconds int) Gorc {
	g := &gorc{}
	var v int32 = 0
	g.count = &v
	g.waitMillis = time.Duration(waitDwellMilliseconds) * time.Millisecond
	return g
}

type gorc struct {
	count      *int32
	waitMillis time.Duration
}

func (g *gorc) Inc() {
	g.IncBy(1)
}

func (g *gorc) IncBy(b int) {
	atomic.AddInt32(g.count, int32(b))
}

func (g *gorc) Dec() {
	g.DecBy(1)
}

func (g *gorc) DecBy(b int) {
	atomic.AddInt32(g.count, int32(b)*-1)
}

func (g gorc) GetCount() int {
	return int(atomic.LoadInt32(g.count))
}

func (g gorc) WaitLow(w int) {
	l := int32(w)
	for atomic.LoadInt32(g.count) >= l {
		time.Sleep(g.waitMillis)
	}
	return
}

func (g gorc) WaitHigh(w int) {
	h := int32(w)
	for atomic.LoadInt32(g.count) <= h {
		time.Sleep(g.waitMillis)
	}
	return
}
