package policy

import (
	"github.com/spf13/cobra"
)

const generateCatalogExamples = `# Generate catalog from gatekeeper-library
gator policy generate-catalog --library-path=/path/to/gatekeeper-library

# Generate with custom output path
gator policy generate-catalog --library-path=. --output=catalog.yaml

# Generate with bundles file
gator policy generate-catalog --library-path=. --bundles=bundles.yaml

# Generate with custom version
gator policy generate-catalog --library-path=. --version=v1.2.0

# Generate with URLs instead of local paths (for publishing)
gator policy generate-catalog --library-path=. --base-url=https://raw.githubusercontent.com/open-policy-agent/gatekeeper-library/master`

func newGenerateCatalogCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

type generateCatalogOptions struct {
	libraryPath    string
	outputPath     string
	catalogName    string
	catalogVersion string
	bundlesFile    string
	baseURL        string
	validate       bool
}

func runGenerateCatalog(opts *generateCatalogOptions) error {
	_ = "STUB: not implemented"
	// Resolve absolute path
	return nil
}

// Verify library path exists

// Generate catalog

// Validate if requested

// Write catalog
