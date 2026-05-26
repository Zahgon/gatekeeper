package target

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/constraints"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/match"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var _ constraints.Matcher = &Matcher{}

// Matcher implements constraint.Matcher.
type Matcher struct {
	match *match.Match
	cache *nsCache
}

func (m *Matcher) Match(review interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// No-op if Match unspecified.
		nil
}

func matchAny(m *Matcher, ns *corev1.Namespace, source types.SourceType, objs ...*unstructured.Unstructured) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func gkReviewToObject(req *gkReview) (*unstructured.Unstructured, *unstructured.Unstructured, *corev1.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}
