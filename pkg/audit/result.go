package audit

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Result struct {
	*types.Result
	obj *unstructured.Unstructured
}

func ToResults(obj *unstructured.Unstructured, resp *types.Responses) []Result {
	_ = "STUB: not implemented"
	return nil
}
