package util

import "sync"

var (
	skipPodOwnerRef bool
	podOwnerRefMux  sync.RWMutex
)

// SetSkipPodOwnerRef configures whether status resources should skip setting
// Pod OwnerReference. This should be enabled when Gatekeeper is running in
// a mode where the pod does not exist in the target cluster (--enable-remote-cluster).
func SetSkipPodOwnerRef(skip bool) { _ = "STUB: not implemented"; return }

// ShouldSkipPodOwnerRef returns true if status resources should not set a
// Pod OwnerReference. This is used in remote cluster mode.
func ShouldSkipPodOwnerRef() bool { _ = "STUB: not implemented"; return false }
