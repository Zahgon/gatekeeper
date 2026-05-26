package verify

import (
	"regexp"
)

type Filter interface {
	// MatchesTest returns true if the Test should be run.
	MatchesTest(*Test) bool
	// MatchesCase returns true if Case caseName in Test testName should be run.
	MatchesCase(testName, caseName string) bool
}

// NewFilter parses run into a Filter for selecting constraint tests and
// individual cases to run.
//
// Empty string results in a Filter which matches all tests and their cases.
//
// Examples:
// 1) NewFiler("require-foo-label//missing-label")
// Matches tests containing the string "require-foo-label"
// and cases containing the string "missing-label". So this would match all
// of the following:
// - Test: "require-foo-label", Case: "missing-label"
// - Test: "not-require-foo-label, Case: "not-missing-label"
// - Test: "require-foo-label", Case: "missing-label-and-annotation"
//
// 2) NewFilter("missing-label")
// Matches cases which either have a name containing "missing-label" or which
// are in a test named "missing-label". Matches the following:
// - Test: "forbid-missing-label", Case: "with-foo-label"
// - Test: "required-labels", Case: "missing-label"
//
// 3) NewFilter("^require-foo-label$//")
// Matches tests which exactly match "require-foo-label". Matches the following:
// - Test: "require-foo-label", Case: "with-foo-label"
// - Test: "require-foo-label", Case: "no-labels"
//
// 4) NewFilter("//empty-object")
// Matches tests whose names contain the string "empty-object". Matches the
// following:
// - Test: "forbid-foo-label", Case: "empty-object"
// - Test: "forbid-foo-label", Case: "another-empty-object"
// - Test: "require-bar-annotation", Case: "empty-object".
func NewFilter(filter string) (Filter, error) { _ = "STUB: not implemented"; return *new(Filter), nil }

// nilFilter matches all tests.
type nilFilter struct{}

var _ Filter = &nilFilter{}

func (f *nilFilter) MatchesTest(*Test) bool { _ = "STUB: not implemented"; return false }

func (f *nilFilter) MatchesCase(string, string) bool {
	_ = "STUB: not implemented"

	// orFilter matches:
	// 1) Tests which are matched by regex.
	// 2) Tests which contain a Case matched by regex.
	// 3) Cases which are matched by regex.
	return false
}

type orFilter struct {
	regex *regexp.Regexp
}

func newOrFilter(filter string) (Filter, error) {
	_ = "STUB: not implemented"
	return *new(Filter), nil
}

func (f *orFilter) MatchesTest(t *Test) bool { _ = "STUB: not implemented"; return false }

func (f *orFilter) MatchesCase(testName, caseName string) bool {
	_ = "STUB: not implemented"
	return false
}

// andFilter matches Cases which match caseRegex which are in Tests which match
// testRegex.
type andFilter struct {
	testRegex *regexp.Regexp
	caseRegex *regexp.Regexp
}

var _ Filter = &andFilter{}

func newAndFilter(testFilter, caseFilter string) (Filter, error) {
	_ = "STUB: not implemented"
	return *new(Filter), nil
}

func (f *andFilter) MatchesTest(t *Test) bool { _ = "STUB: not implemented"; return false }

func (f *andFilter) MatchesCase(testName, caseName string) bool {
	_ = "STUB: not implemented"
	return false
}
