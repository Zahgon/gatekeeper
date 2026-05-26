package testutils

import (
	"io"
	"sync"
	"testing"

	"github.com/go-logr/logr"
)

// NewLogger creates a logger specifically for t which logs directly to the test.
// Use test-specific loggers so that when tests fail, only the log messages from the offending test are printed rather
// than log messages for every test in the package.
func NewLogger(t *testing.T) logr.Logger { _ = "STUB: not implemented"; return *new(logr.Logger) }

type Writer struct {
	t *testing.T

	// stopped tracks whether the test has ended. In this case, further attempts to write log messages will fail and
	// cause flaky panics in unrelated tests. We expect managers to send log messages even after shutting down as there
	// is no guarantee of the order of execution of cancellation logic and the manager shutting down. Still-executing
	// functions will return errors such as "context canceled", but these are not relevant to the test.
	stopped    bool
	stoppedMtx sync.RWMutex
}

var _ io.Writer = &Writer{}

func NewTestWriter(t *testing.T) *Writer { _ = "STUB: not implemented"; return nil }

func (w *Writer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// The test has completed and we shouldn't log anything else to the test runner.

// t.Log is threadsafe.
