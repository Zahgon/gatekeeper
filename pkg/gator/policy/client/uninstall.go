package client

import (
	"context"
)

// UninstallOptions contains options for uninstalling policies.
type UninstallOptions struct {
	// Policies is the list of policy names to uninstall.
	Policies []string
	// DryRun if true, only prints what would be done.
	DryRun bool
}

// UninstallResult contains the result of an uninstall operation.
type UninstallResult struct {
	// Uninstalled is the list of successfully uninstalled policies.
	Uninstalled []string
	// NotFound is the list of policies that were not found.
	NotFound []string
	// NotManaged is the list of policies that exist but are not managed by gator.
	NotManaged []string
	// Failed is the list of policies that failed to uninstall.
	Failed []string
	// Errors contains error messages for failed policies.
	Errors map[string]string
}

// Uninstall removes policies from the cluster.
func Uninstall(ctx context.Context, k8sClient Client, opts UninstallOptions) (*UninstallResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate Gatekeeper is installed (skip if dry-run)

// Uninstall each policy - continue on managed-by conflicts, fail fast on
// unexpected errors (auth, network, etc.) to match install's behavior.

// Conflict errors (not managed by gator) are non-fatal — the policy
// is already tracked in result.NotManaged by uninstallPolicy, so we
// only need to record the message and keep going.

// Unexpected errors are fatal — stop processing.

func uninstallPolicy(ctx context.Context, k8sClient Client, policyName string, dryRun bool, result *UninstallResult) error {
	_ = "STUB: not implemented"
	// Get existing template
	return nil
}

// Distinguish between "not found" and other errors (auth, network, etc.)

// Not found is not an error for uninstall

// Real error - propagate it

// Check if managed by gator

// Delete template
// Note: When the ConstraintTemplate is deleted, Gatekeeper removes the associated
// Constraint CRD. Kubernetes garbage-collects any Constraint CRs when the CRD is deleted.
