package client

import (
	"context"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/gator/policy/catalog"
)

// UpgradeOptions contains options for upgrading policies.
type UpgradeOptions struct {
	// Policies is the list of policy names to upgrade (empty with All=true upgrades all).
	Policies []string
	// All if true, upgrades all installed policies.
	All bool
	// EnforcementAction overrides the enforcement action for constraints.
	EnforcementAction string
	// DryRun if true, only prints what would be done.
	DryRun bool
}

// UpgradeResult contains the result of an upgrade operation.
type UpgradeResult struct {
	// Upgraded is the list of successfully upgraded policies with their version changes.
	Upgraded []VersionChange
	// AlreadyCurrent is the list of policies already at the latest version.
	AlreadyCurrent []string
	// NotFound is the list of policies not found in the catalog.
	NotFound []string
	// NotInstalled is the list of policies not installed in the cluster.
	NotInstalled []string
	// Failed is the list of policies that failed to upgrade.
	Failed []string
	// Errors contains error messages for failed policies.
	Errors map[string]string
}

// VersionChange represents a version change for a policy.
type VersionChange struct {
	Name        string
	FromVersion string
	ToVersion   string
}

// Upgrade upgrades installed policies to their latest versions.
func Upgrade(ctx context.Context, k8sClient Client, fetcher catalog.Fetcher, cat *catalog.PolicyCatalog, opts UpgradeOptions) (*UpgradeResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate Gatekeeper is installed (skip if dry-run)

// Get list of installed policies

// Build map of installed policies

// Determine which policies to upgrade

// Upgrade each policy

// Check if already at latest version

// Upgrade the policy

// Fail fast

func upgradePolicy(ctx context.Context, k8sClient Client, fetcher catalog.Fetcher, policy *catalog.Policy, bundleName string, opts UpgradeOptions) error {
	_ = "STUB: not implemented"
	// Use install with the existing bundle name to preserve constraint installation behavior
	return nil
}

// Pass bundle context so constraints are upgraded too

// Create a minimal catalog for the install

// If this was a bundle-installed policy, we need to ensure constraints are updated
// by temporarily including this policy in the bundle's policy list

// GetUpgradableCount returns the count of policies that have updates available.
func GetUpgradableCount(installed []InstalledPolicy, cat *catalog.PolicyCatalog) int {
	_ = "STUB: not implemented"
	return 0
}

// GetUpgradablePolicies returns a list of policies that have updates available.
func GetUpgradablePolicies(installed []InstalledPolicy, cat *catalog.PolicyCatalog) []VersionChange {
	_ = "STUB: not implemented"
	return nil
}
