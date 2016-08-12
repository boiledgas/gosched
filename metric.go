package gosched

import (
	"sync/atomic"
)

type Metric struct {
	// Counters
	LeafCount      int32
	TreeDepth      int32
	IterationCount int32
	CacheReads     int32
	CacheWrites    int32
	CacheDelete    int32
	// Memory
	CacheSize int32
	TreeSize  int32
}

func (m *Metric) Node() {

}

func (m *Metric) Leaf() {
	atomic.AddInt32(&m.LeafCount, 1)
}

func (m *Metric) Iteration() {
	atomic.AddInt32(&m.IterationCount, 1)
}

func (m *Metric) CacheRead() {
	atomic.AddInt32(&m.CacheReads, 1)
}

func (m *Metric) CacheWrite() {
	atomic.AddInt32(&m.CacheWrites, 1)
}

func (m *Metric) CacheDelete() {
	atomic.AddInt32(&m.CacheDelete, 1)
}
