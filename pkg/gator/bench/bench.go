package bench

import (
	"context"
	"time"

	"github.com/open-policy-agent/frameworks/constraint/pkg/apis"
	constraintclient "github.com/open-policy-agent/frameworks/constraint/pkg/client"
	"github.com/open-policy-agent/frameworks/constraint/pkg/client/drivers/rego"
	"github.com/open-policy-agent/frameworks/constraint/pkg/instrumentation"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/drivers/k8scel"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	k8sruntime "k8s.io/apimachinery/pkg/runtime"
)

const (
	// MinIterationsForP99 is the minimum number of iterations recommended for
	// statistically meaningful P99 metrics.
	MinIterationsForP99 = 1000
)

var scheme *k8sruntime.Scheme

func init() {
	scheme = k8sruntime.NewScheme()
	if err := apis.AddToScheme(scheme); err != nil {
		panic(err)
	}
}

// Run executes the benchmark with the given options and returns results
// for each engine tested.
func Run(opts *Opts) ([]Results, error) {
	_ = "STUB: not implemented"
	// Warn if iterations are too low for meaningful P99 statistics
	return nil, nil
}

// Default concurrency to 1 (sequential)

// Read all resources from files/images

// Categorize objects

// Everything else is a potential review object

// Determine which engines to benchmark

// For "all" engine mode, record warning and continue with other engines

// Check if we have any results

// Add warnings about skipped engines to the first result for visibility

// runBenchmark runs the benchmark for a single engine.
func runBenchmark(
	engine Engine,
	templates []*unstructured.Unstructured,
	constraints []*unstructured.Unstructured,
	reviewObjs []*unstructured.Unstructured,
	opts *Opts,
) (*Results, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the client for this engine

// Add templates (with skip support for incompatible templates)

// Check if this is an engine compatibility issue

// Track the constraint kind this template creates

// Check if all templates were skipped

// Add constraints (skip those whose template was skipped)

// Check if all constraints were skipped

// Add all objects as data (for referential constraints)
// Note: CEL driver doesn't support referential constraints, so skip data loading for CEL

// Note: We don't populate skippedDataObjects for CEL engine because it's expected
// behavior (CEL doesn't support referential data), not an error. The
// ReferentialDataSupported field indicates this engine limitation.

// Warmup phase

// Measurement phase

// Memory profiling: capture memory stats before and after

// Run GC to get clean baseline

// Concurrent or sequential execution based on concurrency setting

// Capture memory stats after measurement

//nolint:gosec // overflow is acceptable for benchmark counts

// Calculate metrics

// makeClient creates a constraint client configured for the specified engine.
func makeClient(engine Engine, gatherStats bool) (*constraintclient.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeRegoDriver(gatherStats bool) (*rego.Driver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeCELDriver(gatherStats bool) (*k8scel.Driver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// runSequentialBenchmark runs the benchmark sequentially (single-threaded).
func runSequentialBenchmark(
	ctx context.Context,
	client *constraintclient.Client,
	reviewObjs []*unstructured.Unstructured,
	opts *Opts,
) ([]time.Duration, int64, []*instrumentation.StatsEntry, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil, nil
}

// Count violations

// Collect stats only from first iteration to avoid excessive data

// reviewResult holds the result of a single review for concurrent execution.
type reviewResult struct {
	duration     time.Duration
	violations   int
	statsEntries []*instrumentation.StatsEntry
	err          error
}

// runConcurrentBenchmark runs the benchmark with multiple goroutines.
func runConcurrentBenchmark(
	ctx context.Context,
	client *constraintclient.Client,
	reviewObjs []*unstructured.Unstructured,
	opts *Opts,
) ([]time.Duration, int64, []*instrumentation.StatsEntry, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil, nil
}

// Create a cancellable context for error propagation

// Create work items

// Result collection

// Launch worker goroutines

// Check if we should stop due to context cancellation

// Signal other goroutines to stop

// Collect stats only from first iteration to avoid excessive data
