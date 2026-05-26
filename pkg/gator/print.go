package gator

import (
	"bytes"
	"io"

	v1 "github.com/open-policy-agent/opa/v1/topdown/print"
)

const (
	// DefaultPrintBufferLimit caps how many bytes of print output gator keeps in memory.
	DefaultPrintBufferLimit int64 = 1 << 20 // 1 MiB

	printOutputTruncatedMsg = "\n... additional print output truncated ...\n"
)

// PrintBuffer is an in-memory writer with a fixed size limit.
type PrintBuffer struct {
	buffer    bytes.Buffer
	remaining int64
	truncated bool
}

// NewPrintBuffer creates a buffer that stores at most limit bytes.
func NewPrintBuffer(limit int64) *PrintBuffer { _ = "STUB: not implemented"; return nil }

// Write writes up to the configured limit and silently discards the rest.
func (b *PrintBuffer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *PrintBuffer) Len() int { _ = "STUB: not implemented"; return 0 }

func (b *PrintBuffer) String() string { _ = "STUB: not implemented"; return "" }

// PrintHook implements the OPA print hook interface to capture print statement output from Rego policies.
type PrintHook struct {
	writer io.Writer
}

// NewPrintHook creates and returns a new instance of PrintHook and writes to writer.
func NewPrintHook(writer io.Writer) PrintHook {
	_ = "STUB: not implemented"
	return *

	// Print writes message to writer passed to PrintHook when it was created.
	new(PrintHook)
}

func (h PrintHook) Print(ctx v1.Context, message string) error {
	_ = "STUB: not implemented"
	return nil
}
