package match

import (
	"errors"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ErrMatch = errors.New("failed to run Match criteria")

// Wildcard represents matching any Group, Version, or Kind.
// Only for use in Match, not ApplyTo.
const Wildcard = "*"

// Matchable represent an object to be matched along with its metadata.
// +kubebuilder:object:generate=false
type Matchable struct {
	Object    client.Object
	Namespace *corev1.Namespace
	Source    types.SourceType
}

// Matches verifies if the given object belonging to the given namespace
// matches Match. Only returns true if all parts of the Match succeed.
func Matches(match *Match, target *Matchable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Simply checking if obj == nil is insufficient here.
// obj can be an interface pointer to nil, such as client.Object(nil), which
// is not equal to just "nil".

// We fail the match if any of these returns false.

// One of the matchers didn't match, so we can exit early.

// matchFunc defines the matching logic of a Top Level Matcher.  A TLM receives the match criteria,
// an object, and the namespace of the object and decides if there is a reason why the object does
// not match.  If the TLM associated with the matching function is not defined by the user, the
// matchFunc should return true.
type matchFunc func(match *Match, target *Matchable) (bool, error)

func namespaceSelectorMatch(match *Match, target *Matchable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Match all non-Namespace cluster-scoped objects.

func labelSelectorMatch(match *Match, target *Matchable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func excludedNamespacesMatch(match *Match, target *Matchable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If we don't have a namespace, we can't disqualify the match

// Fall back to the Namespace the Object declares in case it isn't specified,
// such as for the gator CLI.

// obj is cluster-scoped and not a Namespace.

func namespacesMatch(match *Match, target *Matchable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If we don't have a namespace, we can't disqualify the match

// Fall back to the Namespace the Object declares in case it isn't specified,
// such as for the gator CLI.

func kindsMatch(match *Match, target *Matchable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func namesMatch(match *Match, target *Matchable) (bool, error) {
	_ = "STUB: not implemented"
	// A blank string could be undefined or an intentional blank string by the user.  Either way,
	// we will assume this means "any name".  This goes with the undefined == match everything
	// pattern that we've already got going in the Match.
	return false, nil
}

func scopeMatch(match *Match, target *Matchable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// This includes invalid scopes, such as typos like "cluster" or "Namespace".

func sourceMatch(match *Match, target *Matchable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// An empty 'source' field will default to 'All'

func IsNamespace(obj client.Object) bool { _ = "STUB: not implemented"; return false }

// contains returns true is element is in set.
func contains(set []string, element string) bool { _ = "STUB: not implemented"; return false }
