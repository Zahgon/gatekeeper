package target

import (
	"sync"

	"github.com/open-policy-agent/frameworks/constraint/pkg/handler"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type nsCache struct {
	lock  sync.RWMutex
	cache map[string]*corev1.Namespace
}

var _ handler.Cache = &nsCache{}

func (c *nsCache) Add(key []string, object interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *nsCache) AddNamespace(key string, ns *corev1.Namespace) { _ = "STUB: not implemented"; return }

func (c *nsCache) GetNamespace(name string) *corev1.Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (c *nsCache) Remove(key []string) { _ = "STUB: not implemented"; return }

func (c *nsCache) RemoveNamespace(key string) { _ = "STUB: not implemented"; return }

func toKey(parts []string) string { _ = "STUB: not implemented"; return "" }

func toNamespace(u *unstructured.Unstructured) (*corev1.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
