package mutation

import (
	"github.com/google/uuid"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	annotationMutations  = "gatekeeper.sh/mutations"
	annotationMutationID = "gatekeeper.sh/mutation-id"
)

func mutationAnnotations(obj *unstructured.Unstructured, allAppliedMutations [][]types.Mutator, mutationUUID uuid.UUID) {
	_ = "STUB: not implemented"
	return
}

func toAnnotationMutationsValue(allAppliedMutations [][]types.Mutator) string {
	_ = "STUB: not implemented"
	return ""
}
