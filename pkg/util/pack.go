package util

import (
	"errors"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ErrInvalidPackedName indicates that the packed name of the request to be
// unpacked was invalid.
var ErrInvalidPackedName = errors.New("invalid packed name, want request.Name to match 'gvk:[Kind].[Version].[Group]:[Name]'")

// UnpackRequest unpacks the GVK from a reconcile.Request and returns the separated components.
// GVK is encoded as "Kind.Version.Group".
// Requests are expected to be in the format: {Name: "gvk:EncodedGVK:Name", Namespace: Namespace}.
func UnpackRequest(r reconcile.Request) (schema.GroupVersionKind, reconcile.Request, error) {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionKind), *new(reconcile.Request), nil
}

// EventPackerMapFunc maps an event into a reconcile.Request with embedded GVK information. Must
// be unpacked with UnpackRequest() before use.
func EventPackerMapFunc() handler.MapFunc { _ = "STUB: not implemented"; return *new(handler.MapFunc) }
