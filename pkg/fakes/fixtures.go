package fakes

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func DenyAllRegoTemplate() *templates.ConstraintTemplate { _ = "STUB: not implemented"; return nil }

func DenyAllConstraint() *unstructured.Unstructured { _ = "STUB: not implemented"; return nil }

func ScopedConstraintFor(ep string) *unstructured.Unstructured {
	_ = "STUB: not implemented"
	return nil
}

func ConstraintFor(kind string) *unstructured.Unstructured { _ = "STUB: not implemented"; return nil }
