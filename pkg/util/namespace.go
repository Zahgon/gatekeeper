package util

import (
	"github.com/go-logr/logr"
	"github.com/open-policy-agent/frameworks/constraint/pkg/client/reviews"
	corev1 "k8s.io/api/core/v1"
)

// NamespaceReviewOpt converts a corev1.Namespace to a reviews.ReviewOpt for passing
// namespace context to the constraint client. Returns nil if the namespace is nil
// or if conversion fails (error is logged).
func NamespaceReviewOpt(ns *corev1.Namespace, log logr.Logger) reviews.ReviewOpt {
	_ = "STUB: not implemented"
	return *new(reviews.ReviewOpt)
}

// NamespaceToMap converts a corev1.Namespace to map[string]interface{} for passing
// to the constraint client. This enables CEL expressions to use namespaceObject
// and Rego policies to access input.review.namespaceObject.
func NamespaceToMap(ns *corev1.Namespace) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
