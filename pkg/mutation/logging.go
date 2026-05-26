package mutation

import (
	"github.com/google/uuid"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func logAppliedMutations(message string, mutationUUID uuid.UUID, obj *unstructured.Unstructured, allAppliedMutations [][]types.Mutator, source types.SourceType) {
	_ = "STUB: not implemented"
	return
}

func getNameOrGenerateName(obj *unstructured.Unstructured) string {
	_ = "STUB: not implemented"
	return ""
}

// for generated resources on CREATE, like a pod from a deployment,
// the name has not been populated yet, so we use the GeneratedName instead.
