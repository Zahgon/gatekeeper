package verify

import (
	"regexp"
	"sync"

	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// An Assertion is a declaration about the data returned by running an object
// against a Constraint.
type Assertion struct {
	// Violations, if set, indicates either whether there are violations, or how
	// many violations match this assertion.
	//
	// The value may be either an integer, of a string. If an integer, exactly
	// this number of violations must otherwise match this Assertion. If a string,
	// must be either "yes" or "no". If "yes" at least one violation must match
	// the Assertion to be satisfied. If "no", there must be zero violations
	// matching the Assertion to be satisfied.
	//
	// Defaults to "yes".
	Violations *intstr.IntOrString `json:"violations,omitempty"`

	// Message is a regular expression which matches the Msg field of individual
	// violations.
	//
	// If unset, has no effect and all violations match this Assertion.
	Message *string `json:"message,omitempty"`

	onceMsgRegex sync.Once
	msgRegex     *regexp.Regexp
}

func (a *Assertion) Run(results []*types.Result) error { _ = "STUB: not implemented"; return nil }

// Default to assuming the object fails validation.

func (a *Assertion) matchesCount(matching int32) error { _ = "STUB: not implemented"; return nil }

// Requires a bug in intstr unmarshalling code, or a misuse of the IntOrStr
// type in Go code.

func (a *Assertion) matchesCountInt(matching int32) error { _ = "STUB: not implemented"; return nil }

func (a *Assertion) matchesCountStr(matching int32) error { _ = "STUB: not implemented"; return nil }

func (a *Assertion) matches(result *types.Result) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (a *Assertion) getMsgRegex() (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
