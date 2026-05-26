package expansion

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
)

const (
	childMsgPrefix = "[Implied by %s]"

	ChildStatLabel = "Implied by"
)

// AggregateResponses aggregates all responses from children into the parent.
// Child result messages will be prefixed with a string to indicate the msg
// is implied by a ExpansionTemplate.
func AggregateResponses(templateName string, parent *types.Responses, child *types.Responses) {
	_ = "STUB: not implemented"
	return
}

// AggregateStats aggregates all stats from the child Responses.StatsEntry
// into the parent Responses.StatsEntry. Child Stats will have a label to
// indicate that they come from an ExpansionTemplate usage.
func AggregateStats(templateName string, parent *types.Responses, child *types.Responses) {
	_ = "STUB: not implemented"
	return
}

func OverrideEnforcementAction(action string, resps *types.Responses) {
	_ = "STUB: not implemented"
	// If the enforcement action is empty, do not override
	return
}

func addPrefixToChildMsgs(templateName string, res *types.Response) {
	_ = "STUB: not implemented"
	return
}
