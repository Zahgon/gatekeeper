package client

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

// ConstraintTemplateGVR is the GroupVersionResource for ConstraintTemplates.
var ConstraintTemplateGVR = schema.GroupVersionResource{
	Group:    "templates.gatekeeper.sh",
	Version:  "v1",
	Resource: "constrainttemplates",
}

// InstalledPolicy represents an installed policy in the cluster.
type InstalledPolicy struct {
	Name        string
	Version     string
	Bundle      string
	InstalledAt string
	ManagedBy   string
}

// Client provides operations for managing policies in a Kubernetes cluster.
type Client interface {
	// GatekeeperInstalled checks if Gatekeeper CRDs are installed.
	GatekeeperInstalled(ctx context.Context) (bool, error)

	// ListManagedTemplates lists all ConstraintTemplates managed by gator.
	ListManagedTemplates(ctx context.Context) ([]InstalledPolicy, error)

	// GetTemplate gets a ConstraintTemplate by name.
	GetTemplate(ctx context.Context, name string) (*unstructured.Unstructured, error)

	// InstallTemplate installs or updates a ConstraintTemplate.
	InstallTemplate(ctx context.Context, template *unstructured.Unstructured) error

	// InstallConstraint installs or updates a Constraint.
	InstallConstraint(ctx context.Context, constraint *unstructured.Unstructured) error

	// GetConstraint gets a Constraint by GVR and name.
	GetConstraint(ctx context.Context, gvr schema.GroupVersionResource, name string) (*unstructured.Unstructured, error)

	// DeleteTemplate deletes a ConstraintTemplate.
	DeleteTemplate(ctx context.Context, name string) error

	// DeleteConstraint deletes a Constraint.
	DeleteConstraint(ctx context.Context, gvr schema.GroupVersionResource, name string) error

	// WaitForTemplateReady waits for a ConstraintTemplate to have status.created = true.
	WaitForTemplateReady(ctx context.Context, templateName string, timeout time.Duration) error

	// WaitForConstraintCRD waits for the Constraint CRD to be available.
	WaitForConstraintCRD(ctx context.Context, kind string, timeout time.Duration) error
}

// K8sClient implements Client using the Kubernetes API.
type K8sClient struct {
	dynamicClient dynamic.Interface
}

// NewK8sClient creates a new K8sClient using the default kubeconfig.
func NewK8sClient() (*K8sClient, error) { _ = "STUB: not implemented"; return nil, nil }

// NewK8sClientWithConfig creates a new K8sClient with the given config.
func NewK8sClientWithConfig(config *rest.Config) (*K8sClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getKubeConfig() (*rest.Config, error) {
	_ = "STUB: not implemented"
	// Try in-cluster config first
	return nil, nil
}

// Fall back to kubeconfig

// GatekeeperInstalled checks if Gatekeeper CRDs are installed.
func (c *K8sClient) GatekeeperInstalled(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ListManagedTemplates lists all ConstraintTemplates managed by gator.
func (c *K8sClient) ListManagedTemplates(ctx context.Context) ([]InstalledPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTemplate gets a ConstraintTemplate by name.
func (c *K8sClient) GetTemplate(ctx context.Context, name string) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InstallTemplate installs or updates a ConstraintTemplate.
func (c *K8sClient) InstallTemplate(ctx context.Context, template *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

// Create new

// Update existing

// InstallConstraint installs or updates a Constraint.
func (c *K8sClient) InstallConstraint(ctx context.Context, constraint *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

// Create new

// Update existing

// GetConstraint gets a Constraint by GVR and name.
func (c *K8sClient) GetConstraint(ctx context.Context, gvr schema.GroupVersionResource, name string) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteTemplate deletes a ConstraintTemplate.
func (c *K8sClient) DeleteTemplate(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteConstraint deletes a Constraint.
func (c *K8sClient) DeleteConstraint(ctx context.Context, gvr schema.GroupVersionResource, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForTemplateReady waits for a ConstraintTemplate to have status.created = true.
func (c *K8sClient) WaitForTemplateReady(ctx context.Context, templateName string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if status.created is true

// Continue polling

// WaitForConstraintCRD waits for the Constraint CRD to be available.
func (c *K8sClient) WaitForConstraintCRD(ctx context.Context, kind string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Continue polling

func getConstraintResource(kind string) string {
	_ = "STUB: not implemented"
	// Gatekeeper constraint CRDs use lowercase kind as the resource name
	// e.g., Kind=K8sPSPAppArmor -> resource=k8spspapparmor
	return ""
}

// isCRDNotRegisteredError checks if the error indicates the CRD/resource type is not registered.
// This is different from IsNotFound which indicates the resource instance doesn't exist.
func isCRDNotRegisteredError(err error) bool { _ = "STUB: not implemented"; return false }

// Check for "no matches" error string which occurs when CRD is not yet registered
