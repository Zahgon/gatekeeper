package catalog

import (
	"sync"
	"time"
)

// PolicyCatalog represents the root catalog structure for policy discovery.
type PolicyCatalog struct {
	APIVersion string          `json:"apiVersion" yaml:"apiVersion"`
	Kind       string          `json:"kind" yaml:"kind"`
	Metadata   CatalogMetadata `json:"metadata" yaml:"metadata"`
	Bundles    []Bundle        `json:"bundles" yaml:"bundles"`
	Policies   []Policy        `json:"policies" yaml:"policies"`

	// Cached indexes for O(1) lookups (built lazily, thread-safe)
	policyIndex     map[string]int `json:"-" yaml:"-"`
	bundleIndex     map[string]int `json:"-" yaml:"-"`
	policyIndexOnce sync.Once      `json:"-" yaml:"-"`
	bundleIndexOnce sync.Once      `json:"-" yaml:"-"`
}

// CatalogMetadata contains metadata about the catalog itself.
type CatalogMetadata struct {
	Name       string    `json:"name" yaml:"name"`
	Version    string    `json:"version" yaml:"version"`
	UpdatedAt  time.Time `json:"updatedAt" yaml:"updatedAt"`
	Repository string    `json:"repository" yaml:"repository"`
}

// Bundle represents a curated set of policies with pre-configured constraints.
type Bundle struct {
	// Name is the unique identifier for this bundle (e.g., "pod-security-baseline").
	Name string `json:"name" yaml:"name"`
	// Description provides a human-readable summary of the bundle's purpose.
	Description string `json:"description" yaml:"description"`
	// Inherits specifies a parent bundle whose policies are included in this bundle.
	Inherits string `json:"inherits,omitempty" yaml:"inherits,omitempty"`
	// Policies lists policy names (not full Policy objects) included in this bundle.
	Policies []string `json:"policies" yaml:"policies"`
}

// Policy represents a single policy available in the catalog.
type Policy struct {
	// Name is the unique identifier for this policy, typically matching the ConstraintTemplate name.
	Name string `json:"name" yaml:"name"`
	// Version is the semantic version of this policy (e.g., "v1.2.3").
	Version string `json:"version" yaml:"version"`
	// Description provides a human-readable summary of what this policy enforces.
	Description string `json:"description" yaml:"description"`
	// Category groups related policies (e.g., "general", "pod-security").
	Category string `json:"category" yaml:"category"`
	// TemplatePath is the URL or relative path to the ConstraintTemplate YAML.
	TemplatePath string `json:"templatePath" yaml:"templatePath"`
	// BundleConstraints maps bundle names to their constraint file paths.
	// Different bundles may require different constraint configurations for the same template
	// (e.g., baseline vs restricted PSS profiles need different constraint values).
	BundleConstraints map[string]string `json:"bundleConstraints,omitempty" yaml:"bundleConstraints,omitempty"`
	// DocumentationURL links to external documentation for this policy.
	DocumentationURL string `json:"documentationUrl,omitempty" yaml:"documentationUrl,omitempty"`
	// Bundles lists which bundles include this policy (reverse reference for discovery).
	Bundles []string `json:"bundles,omitempty" yaml:"bundles,omitempty"`
}

// GetPolicy returns the policy with the given name, or nil if not found.
// Uses O(1) indexed lookup after first call. Thread-safe.
func (c *PolicyCatalog) GetPolicy(name string) *Policy {
	_ = "STUB: not implemented"
	// Build index lazily on first lookup (thread-safe)
	return nil
}

// GetBundle returns the bundle with the given name, or nil if not found.
// Uses O(1) indexed lookup after first call. Thread-safe.
func (c *PolicyCatalog) GetBundle(name string) *Bundle {
	_ = "STUB: not implemented"
	// Build index lazily on first lookup (thread-safe)
	return nil
}

// MaxInheritanceDepth is the maximum depth of bundle inheritance to prevent runaway expansion.
const MaxInheritanceDepth = 10

// ResolveBundlePolicies returns all policy names for a bundle, including inherited policies.
// Inheritance is processed parent-first (deepest ancestor first), so child policies can override.
func (c *PolicyCatalog) ResolveBundlePolicies(bundleName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First, build the inheritance chain from child to ancestors

// Check for circular inheritance

// Check depth limit

// Now process in reverse order (parent-first, deepest ancestor first)
// This ensures parent policies are added first, and child policies can be deduplicated

// BundleNotFoundError is returned when a bundle cannot be found.
type BundleNotFoundError struct {
	Name string
}

func (e *BundleNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

// PolicyNotFoundError is returned when a policy cannot be found.
type PolicyNotFoundError struct {
	Name string
}

func (e *PolicyNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }
