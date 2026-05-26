package syncutil

import "sync"

type ConcurrentErrorSlice struct {
	s  []error
	mu *sync.RWMutex
}

func NewConcurrentErrorSlice() ConcurrentErrorSlice {
	_ = "STUB: not implemented"
	return *new(ConcurrentErrorSlice)
}

func (c ConcurrentErrorSlice) Append(e error) ConcurrentErrorSlice {
	_ = "STUB: not implemented"
	return *new(ConcurrentErrorSlice)
}

func (c ConcurrentErrorSlice) Last() error { _ = "STUB: not implemented"; return nil }

func (c ConcurrentErrorSlice) GetSlice() []error { _ = "STUB: not implemented"; return nil }
