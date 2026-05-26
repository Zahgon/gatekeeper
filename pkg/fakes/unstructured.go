package fakes

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func UnstructuredFor(gvk schema.GroupVersionKind, namespace, name string) *unstructured.Unstructured {
	_ = "STUB: not implemented"
	return nil
}
