package util

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type KindVersionName struct {
	Group     string
	Kind      string
	Version   string
	Namespace string
	Name      string
}

func GetUniqueKey(obj unstructured.Unstructured) KindVersionName {
	_ = "STUB: not implemented"
	return *new(KindVersionName)
}
