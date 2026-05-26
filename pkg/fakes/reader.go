package fakes

import (
	"context"
	"sync"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SpyReader struct {
	client.Reader
	ListFunc func(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error
}

func (r SpyReader) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

// FailureInjector can be used in combination with the SpyReader to simulate transient
// failures for network calls.
type FailureInjector struct {
	mu       sync.Mutex
	failures map[string]int // registers GVK.Kind and how many times to fail
}

func (f *FailureInjector) SetFailures(kind string, failures int) { _ = "STUB: not implemented"; return }

// CheckFailures looks at the count of failures and returns true
// if there are still failures for the kind to consume, false otherwise.
func (f *FailureInjector) CheckFailures(kind string) bool { _ = "STUB: not implemented"; return false }

func NewFailureInjector() *FailureInjector { _ = "STUB: not implemented"; return nil }
