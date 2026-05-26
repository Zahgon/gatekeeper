package catalog

import (
	"regexp"
)

// GeneratorOptions contains options for generating a catalog.
type GeneratorOptions struct {
	// LibraryPath is the root path of the gatekeeper-library repository.
	LibraryPath string
	// CatalogName is the name of the catalog.
	CatalogName string
	// CatalogVersion is the version of the catalog.
	CatalogVersion string
	// Repository is the repository URL.
	Repository string
	// BundlesFile is an optional path to a bundles definition file.
	BundlesFile string
	// BaseURL is the base URL for template/constraint paths.
	// If set, paths will be converted to full URLs (e.g., https://raw.githubusercontent.com/.../library/...).
	// If empty, relative paths will be used (e.g., library/...).
	BaseURL string
}

// Bundle annotation key for metadata.gatekeeper.sh/bundle.
const bundleAnnotationKey = "metadata.gatekeeper.sh/bundle"

// bundleDescriptions provides descriptions for well-known bundles.
var bundleDescriptions = map[string]string{
	"pod-security-baseline": `Enforces Pod Security Standards at Baseline level. Prevents known privilege escalations.
See https://kubernetes.io/docs/concepts/security/pod-security-standards/`,
	"pod-security-restricted": `Enforces Pod Security Standards at Restricted level. Includes all Baseline controls plus additional hardening.
See https://kubernetes.io/docs/concepts/security/pod-security-standards/`,
}

// GenerateCatalog generates a PolicyCatalog from a gatekeeper-library directory structure.
func GenerateCatalog(opts *GeneratorOptions) (*PolicyCatalog, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Walk the library directory to find policies

// Look for template.yaml files

// Sort policies by name for consistent output

// Load bundles from bundles file if provided

// Update policies with bundle membership

// Auto-generate Pod Security Standards bundles

// Convert paths to URLs if base URL is provided

// parsePolicyFromTemplate reads a template.yaml and extracts policy metadata.
func parsePolicyFromTemplate(templatePath, libraryRoot string) (*Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse template to extract metadata

// Not a constraint template

// Get relative path from library root

// Determine category from path (e.g., library/general/requiredlabels -> general)

// Extract version from annotations or use default
// The gatekeeper-library uses metadata.gatekeeper.sh/version annotation

// Normalize version to have v prefix

// Get description from annotations

// Look for constraint files in samples directory

// Get documentation URL

// Get bundle membership from annotation

// Parse comma-separated bundle names

// Build BundleConstraints by discovering per-bundle constraint files.
// The library convention is that sample directories with names containing
// a bundle keyword (e.g., "baseline", "restricted") provide bundle-specific
// constraint configurations. If no bundle-specific directory is found,
// the first constraint file discovered is used as a fallback.

// extractCategory extracts the category from the library path.
func extractCategory(relPath string) string {
	_ = "STUB: not implemented"
	// Path format: library/<category>/<policy>/template.yaml
	return ""
}

// Normalize category names

// Handle pod-security-policy -> pod-security

// findBundleConstraints discovers per-bundle constraint files in the samples directory.
// It maps each bundle to its constraint file by matching sample directory names
// to bundle keywords. For example, bundle "pod-security-baseline" matches a sample
// directory named "psp-capabilities-baseline" (contains "baseline").
// If no bundle-specific match is found, the first constraint file is used as fallback.
// Returns nil if no bundles or no constraint files are found.
func findBundleConstraints(templateDir, libraryRoot string, bundles []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Scan all sample subdirectories for constraint files, indexed by dir name.
// constraintsByDir maps sample directory name → relative constraint path.

// Try .yml extension

// Match bundles to sample directories.
// Extract the distinguishing keyword from bundle name (last segment after "pod-security-").
// e.g., "pod-security-baseline" → "baseline", "pod-security-restricted" → "restricted"

// Extract keyword: use the last hyphen-separated segment of the bundle name

// Look for a sample dir whose name contains the keyword

// isConstraintFile checks if the YAML content is a Constraint.
func isConstraintFile(data []byte) bool { _ = "STUB: not implemented"; return false }

// BundlesFile represents the structure of a bundles definition file.
type BundlesFile struct {
	Bundles []Bundle `yaml:"bundles" json:"bundles"`
}

// loadBundlesFile loads bundle definitions from a YAML file.
func loadBundlesFile(path string) ([]Bundle, error) { _ = "STUB: not implemented"; return nil, nil }

// updatePolicyBundles updates the Bundles field of each policy based on bundle membership.
func updatePolicyBundles(catalog *PolicyCatalog) { _ = "STUB: not implemented"; return }

// Build reverse mapping from policy to bundles

// Update policies

// WriteCatalog writes a PolicyCatalog to a YAML file.
func WriteCatalog(catalog *PolicyCatalog, outputPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// Add header comment

// ValidateCatalogSchema validates a catalog against the expected schema.
func ValidateCatalogSchema(catalog *PolicyCatalog) error { _ = "STUB: not implemented"; return nil }

// Validate policies

// Validate bundles reference existing policies

// semverPattern is a compiled regex for validating semantic version strings.
var semverPattern = regexp.MustCompile(`^v?\d+\.\d+\.\d+(-[\w.]+)?(\+[\w.]+)?$`)

// isValidVersion checks if a version string is valid semver format.
func isValidVersion(version string) bool { _ = "STUB: not implemented"; return false }

// generatePSSBundles auto-generates bundles from policy annotations.
// Bundles are created based on metadata.gatekeeper.sh/bundle annotations in templates.
func generatePSSBundles(catalog *PolicyCatalog) {
	_ = "STUB: not implemented"
	// Build mapping from bundle name to policies
	return
}

// Create bundles from the mapping
// Sort bundle names for consistent output

// Sort policy names for consistent output

// convertPathsToURLs converts all relative paths in the catalog to full URLs.
// The baseURL should be the raw content URL prefix (e.g., https://raw.githubusercontent.com/open-policy-agent/gatekeeper-library/master).
func convertPathsToURLs(catalog *PolicyCatalog, baseURL string) {
	_ = "STUB: not implemented"
	// Ensure baseURL doesn't have trailing slash
	return
}
