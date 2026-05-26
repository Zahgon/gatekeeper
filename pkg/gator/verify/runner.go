package verify

import (
	"context"
	"io/fs"

	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/gator"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/gator/expand"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// Runner defines logic independent of how tests are run and the results are
// printed.
type Runner struct {
	// filesystem is the filesystem the Runner interacts with to read Suites and objects.
	filesystem fs.FS

	// newClient instantiates a Client for compiling Templates/Constraints, and
	// validating objects against them.
	newClient func(opts ...gator.Opt) (gator.Client, error)

	scheme *runtime.Scheme

	includeTrace bool
}

func NewRunner(filesystem fs.FS, newClient func(opts ...gator.Opt) (gator.Client, error), opts ...RunnerOptions) (*Runner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RunnerOptions func(*Runner)

func IncludeTrace(includeTrace bool) RunnerOptions {
	_ = "STUB: not implemented"
	return *new(RunnerOptions)
}

// Run executes all Tests in the Suite and returns the results.
func (r *Runner) Run(ctx context.Context, filter Filter, s *Suite) SuiteResult {
	_ = "STUB: not implemented"
	return *new(SuiteResult)
}

// runTests runs every Test in Suite.
func (r *Runner) runTests(ctx context.Context, filter Filter, suitePath string, tests []Test) ([]TestResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Runner) skipTest(t *Test) TestResult { _ = "STUB: not implemented"; return *new(TestResult) }

// runTest runs an individual Test.
func (r *Runner) runTest(ctx context.Context, suiteDir string, filter Filter, t *Test) TestResult {
	_ = "STUB: not implemented"
	return *new(TestResult)
}

func (r *Runner) tryAddConstraint(ctx context.Context, suiteDir string, t *Test) error {
	_ = "STUB: not implemented"
	return nil
}

// runCases executes every Case in the Test. Returns the results for every Case,
// or an error if there was a problem executing the Test.
func (r *Runner) runCases(ctx context.Context, suiteDir string, filter Filter, t *Test) ([]CaseResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Runner) skipCase(tc *Case) CaseResult { _ = "STUB: not implemented"; return *new(CaseResult) }

func (r *Runner) makeTestClient(ctx context.Context, newClient func() (gator.Client, error), suiteDir string, t *Test) (gator.Client, error) {
	_ = "STUB: not implemented"
	return *new(gator.Client), nil
}

func (r *Runner) makeTestExpander(suiteDir string, t *Test) (*expand.Expander, error) {
	_ = "STUB: not implemented"
	// Support Mutator logic? Then we need to add support for mutators as well or do we just ignore them?
	return nil, nil
}

func (r *Runner) addConstraint(ctx context.Context, suiteDir, constraintPath string, client gator.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Runner) addTemplate(suiteDir, templatePath string, client gator.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// RunCase executes a Case and returns the result of the run.
func (r *Runner) runCase(ctx context.Context, newClient func(opts ...gator.Opt) (gator.Client, error), newExpander func() (*expand.Expander, error), suiteDir string, tc *Case) CaseResult {
	_ = "STUB: not implemented"
	return *new(CaseResult)
}

func (r *Runner) checkCase(ctx context.Context, newClient func() (gator.Client, error), newExpander func() (*expand.Expander, error), suiteDir string, tc *Case) (trace *string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Test cases must define at least one assertion.

func (r *Runner) runReview(ctx context.Context, newClient func() (gator.Client, error), newExpander func() (*expand.Expander, error), suiteDir string, tc *Case) (*types.Responses, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check to see if obj is an AdmissionReview kind

// otherwise our object is some other k8s object

func (r *Runner) validateAndReviewAdmissionReviewRequest(ctx context.Context, c gator.Client, toReview *unstructured.Unstructured) (*types.Responses, error) {
	_ = "STUB: not implemented"
	// convert unstructured into AdmissionReview, don't allow unknown fields
	return nil, nil
}

// then this admission review did not actually pass in an AdmissionRequest

// validate the AdmissionReview to match k8s api server behavior

// make sure that kind is set for object, oldObject if present since we expect it on Decode
// https://github.com/kubernetes/apimachinery/blob/27a96d86e70e9bbd40639a0539a14423f55afa07/pkg/apis/meta/v1/unstructured/helpers.go#L341

func (r *Runner) addInventory(ctx context.Context, c gator.Client, suiteDir, inventoryPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// readCaseObjects reads objects at path in filesystem f. Returns:
// 1) The object to review for the test case
// 2) The objects to add to data.inventory
// 3) Any errors encountered parsing the objects
//
// The final object in path is the object to review.
func readObjects(f fs.FS, path string) ([]*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
