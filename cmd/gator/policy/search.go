package policy

import (
	"github.com/spf13/cobra"
)

var (
	searchCategory string
	searchBundle   string
	searchOutput   string
)

func newSearchCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runSearch(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Load catalog from cache

// If no cache, try to fetch

// Build bundle filter set if --bundle is specified

// Search policies

// Filter by bundle if specified

// Filter by category if specified

// Search in name and description

// Output results
