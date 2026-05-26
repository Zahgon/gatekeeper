package transform

import (
	"sync"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var vapMux sync.RWMutex

var VapAPIEnabled *bool

var GroupVersion *schema.GroupVersion

// SetVapAPIEnabled sets the VapAPIEnabled flag in a thread-safe manner.
// Use this instead of directly assigning transform.VapAPIEnabled when the
// value may be read concurrently (e.g., by a running controller).
func SetVapAPIEnabled(enabled *bool) { _ = "STUB: not implemented"; return }

// SetGroupVersion sets the GroupVersion in a thread-safe manner.
// Use this instead of directly assigning transform.GroupVersion when the
// value may be read concurrently (e.g., by a running controller).
func SetGroupVersion(gv *schema.GroupVersion) { _ = "STUB: not implemented"; return }

func IsVapAPIEnabled(log *logr.Logger) (bool, *schema.GroupVersion) {
	_ = "STUB: not implemented"
	return false, nil
}

// Do not cache failure — allow retry on next reconcile

// Do not cache failure — allow retry on next reconcile

// Discovery failed — do not cache, allow retry on next reconcile
