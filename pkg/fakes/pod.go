package fakes

import (
	corev1 "k8s.io/api/core/v1"
)

// Pod creates a Pod for use in testing or debugging logic.
func Pod(opts ...Opt) *corev1.Pod { _ = "STUB: not implemented"; return nil }
