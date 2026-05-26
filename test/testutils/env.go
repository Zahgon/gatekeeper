package testutils

import (
	"testing"
)

// Setenv sets os environment variable key to value.
// Registers the environment variable to be set to its original value at the end
// of the test.
//
// This prevents cross-talk between tests, as some may implicitly come to rely on other tests running before them in
// order to succeed.
func Setenv(t *testing.T, key, value string) { _ = "STUB: not implemented"; return }
