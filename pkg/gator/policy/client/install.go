package client

import (
	"context"
	"time"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/gator/policy/catalog"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// DefaultReconcileTimeout is the default timeout for waiting on Gatekeeper to reconcile resources.
	DefaultReconcileTimeout = 120 * time.Second
)

// InstallOptions contains options for installing policies.
type InstallOptions struct {
	// Policies is the list of policy names to install.
	Policies []string
	// Bundles is the list of bundle names to install.
	Bundles []string
	// EnforcementAction overrides the enforcement action for constraints.
	EnforcementAction string
	// DryRun if true, only prints what would be done.
	DryRun bool
}

// InstallResult contains the result of an install operation.
type InstallResult struct {
	// Installed is the list of successfully installed policies.
	Installed []string
	// Skipped is the list of skipped policies (already at same version).
	Skipped []string
	// Failed is the list of policies that failed to install.
	Failed []string
	// Errors contains error messages for failed policies.
	Errors map[string]string
	// ConflictErr is set if a conflict error occurred (resource not managed by gator).
	ConflictErr *ConflictError
	// ConstraintsInstalled is the number of constraints installed.
	ConstraintsInstalled int
	// TemplatesInstalled is the number of templates installed.
	TemplatesInstalled int
	// TotalRequested is the total number of policies requested for installation.
	TotalRequested int
}

// Install installs policies from the catalog.
func Install(ctx context.Context, k8sClient Client, fetcher catalog.Fetcher, cat *catalog.PolicyCatalog, opts *InstallOptions) (*InstallResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Determine which policies to install

// policyBundle tracks which bundle a policy was resolved from (first match wins).

// If bundles are specified, resolve bundle policies first

// Add any additional positional policies (deduplicated)
// These are installed as template-only (no constraints) even when bundle is set

// Track total policies requested

// Validate Gatekeeper is installed (skip if dry-run)

// Install each policy

// Fail fast per MVP design

// Determine if this policy should install constraints.
// Only bundle-resolved policies get constraints; positional policies get template-only.

// Preserve typed error for conflict detection

// Fail fast - stop on first error

func installPolicy(ctx context.Context, k8sClient Client, fetcher catalog.Fetcher, policy *catalog.Policy, bundleName string, opts *InstallOptions, result *InstallResult) (skipped bool, err error) {
	_ = "STUB: not implemented"
	// Fetch template YAML
	return false, nil
}

// Parse template

// Check for existing template

// Template exists - check if managed by gator

// Check if same version

// Add labels and annotations

// Install or update template if not already at same version

// Install constraint if bundle has a constraint path defined

// Return whether this policy was skipped (already at same version)

func installConstraint(ctx context.Context, k8sClient Client, fetcher catalog.Fetcher, policy *catalog.Policy, constraintPath string, bundleName string, opts *InstallOptions, result *InstallResult, template *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

// Override enforcement action if specified

// Add labels

// Install constraint

// Wait for the template status to show created=true

// Wait for the constraint CRD to be available

// Check if constraint already exists and is not managed by gator

// Preserve existing enforcement action if not explicitly overridden.
// This ensures upgrades don't silently revert a user's enforcement setting.

func setEnforcementAction(constraint *unstructured.Unstructured, action string) error {
	_ = "STUB: not implemented"
	return nil
}

// constraintGVR returns the GroupVersionResource for a constraint kind.
func constraintGVR(kind string) schema.GroupVersionResource {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionResource)
}

// GatekeeperNotInstalledError is returned when Gatekeeper CRDs are not found.
type GatekeeperNotInstalledError struct{}

func (e *GatekeeperNotInstalledError) Error() string { _ = "STUB: not implemented"; return "" }

// ConflictError is returned when a resource exists but is not managed by gator.
type ConflictError struct {
	ResourceKind string
	ResourceName string
}

func (e *ConflictError) Error() string { _ = "STUB: not implemented"; return "" }
