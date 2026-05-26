package policy

import (
	"github.com/spf13/cobra"
)

var (
	updateInsecure bool
	updateOutput   string
)

func newUpdateCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runUpdate(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

// Create printer

// Progress message to stderr so it doesn't pollute structured output

// Fetch catalog

// Check for insecure HTTP error and provide helpful message

// Parse to validate

// Save to cache

// Build update result

// Check for upgradable policies if cluster is accessible
