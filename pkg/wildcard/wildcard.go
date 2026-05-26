package wildcard

// +kubebuilder:validation:Pattern=`^\*?[-:a-z0-9]*\*?$`

// A string that supports globbing at its front and end. Ex: "kube-*" will match "kube-system" or
// "kube-public", "*-system" will match "kube-system" or "gatekeeper-system", "*system*" will
// match "system-kube" or "kube-system".  The asterisk is required for wildcard matching.
//
//nolint:revive
type Wildcard string

// Matches returns true if the candidate parameter is either an exact match of the Wildcard,
// or if the Wildcard is a valid glob-match for the candidate.  The Wildcard must start or end
// in a "*" to be considered a glob.
func (w Wildcard) Matches(candidate string) bool { _ = "STUB: not implemented"; return false }

func (w Wildcard) MatchesGenerateName(candidate string) bool {
	_ = "STUB: not implemented"
	return false
}
