package core

import (
	statusv1beta1 "github.com/open-policy-agent/gatekeeper/v3/apis/status/v1beta1"
	apiTypes "k8s.io/apimachinery/pkg/types"
)

type statusUpdate func(status *statusv1beta1.MutatorPodStatus)

func setID(id apiTypes.UID) statusUpdate { _ = "STUB: not implemented"; return *new(statusUpdate) }

func setGeneration(generation int64) statusUpdate {
	_ = "STUB: not implemented"
	return *new(statusUpdate)
}

func setErrors(err error) statusUpdate { _ = "STUB: not implemented"; return *new(statusUpdate) }

// Replaces any existing errors, if there was one.

func setEnforced(isEnforced bool) statusUpdate {
	_ = "STUB: not implemented"
	return *new(statusUpdate)
}
