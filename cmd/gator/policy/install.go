package policy

import (
	"context"
	"time"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/gator/policy/client"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	installBundles           []string
	installEnforcementAction string
	installDryRun            bool
	installOutput            string
)

func newInstallCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runInstall(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Validate arguments

// Note: --bundle with positional policies IS allowed per design.
// Bundle is processed first, then individual policies are added (template-only).

// Validate enforcement action

// Warn if enforcement action specified without bundle (template-only installs don't have constraints)

// Create printer

// Parse policy names

// Load catalog

// Create fetcher for templates/constraints with the cached catalog source URL as base

// Create Kubernetes client (unless dry-run)

// Build install options

// Perform installation

// Check for specific error types

// Build output result

// Print results

// Return appropriate error for non-success cases

// dryRunClient is a no-op client for dry-run mode.
type dryRunClient struct{}

func (c *dryRunClient) GatekeeperInstalled(_ context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *dryRunClient) ListManagedTemplates(_ context.Context) ([]client.InstalledPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *dryRunClient) GetTemplate(_ context.Context, _ string) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *dryRunClient) InstallTemplate(_ context.Context, _ *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) InstallConstraint(_ context.Context, _ *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) GetConstraint(_ context.Context, _ schema.GroupVersionResource, _ string) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *dryRunClient) DeleteTemplate(_ context.Context, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) DeleteConstraint(_ context.Context, _ schema.GroupVersionResource, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) WaitForTemplateReady(_ context.Context, _ string, _ time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dryRunClient) WaitForConstraintCRD(_ context.Context, _ string, _ time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
