package bench

import (
	"fmt"
	"time"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/gator/bench"
	"github.com/spf13/cobra"
)

const (
	examples = `# Benchmark policies with default settings (1000 iterations, cel engine)
gator bench --filename="policies/"

# Benchmark with both Rego and CEL engines
gator bench --filename="policies/" --engine=all

# Benchmark with custom iterations and warmup
gator bench --filename="policies/" --iterations=500 --warmup=50

# Benchmark with concurrent load (simulates real webhook traffic)
gator bench --filename="policies/" --concurrency=10

# Output results as JSON
gator bench --filename="policies/" --output=json

# Benchmark policies from multiple sources
gator bench --filename="templates/" --filename="constraints/" --filename="resources/"

# Benchmark from OCI image
gator bench --image="ghcr.io/example/policies:latest"

# Benchmark with memory profiling
gator bench --filename="policies/" --memory

# Save benchmark results as baseline
gator bench --filename="policies/" --save=baseline.json

# Compare against baseline (fail only if BOTH >10% regression AND >1ms absolute increase)
# This prevents false positives for fast policies where small absolute changes appear as large percentages
gator bench --filename="policies/" --compare=baseline.json --threshold=10 --min-threshold=1ms`
)

// Cmd is the cobra command for the bench subcommand.
var Cmd = &cobra.Command{
	Use:   "bench",
	Short: "Benchmark policy evaluation performance",
	Long: `Benchmark evaluates the performance of Gatekeeper policies by running
constraint evaluation against test resources and measuring latency metrics.

This command loads ConstraintTemplates, Constraints, and Kubernetes resources
from the specified files or directories, then repeatedly evaluates the resources
against the constraints to gather performance statistics.

Supports both Rego and CEL policy engines for comparison.`,
	Example: examples,
	Run:     run,
	Args:    cobra.NoArgs,
}

var (
	flagFilenames    []string
	flagImages       []string
	flagTempDir      string
	flagEngine       string
	flagIterations   int
	flagWarmup       int
	flagConcurrency  int
	flagOutput       string
	flagStats        bool
	flagMemory       bool
	flagSave         string
	flagCompare      string
	flagThreshold    float64
	flagMinThreshold time.Duration
)

const (
	flagNameFilename     = "filename"
	flagNameImage        = "image"
	flagNameTempDir      = "tempdir"
	flagNameEngine       = "engine"
	flagNameIterations   = "iterations"
	flagNameWarmup       = "warmup"
	flagNameConcurrency  = "concurrency"
	flagNameOutput       = "output"
	flagNameStats        = "stats"
	flagNameMemory       = "memory"
	flagNameSave         = "save"
	flagNameCompare      = "compare"
	flagNameThreshold    = "threshold"
	flagNameMinThreshold = "min-threshold"
)

func init() {
	Cmd.Flags().StringArrayVarP(&flagFilenames, flagNameFilename, "f", []string{},
		"a file or directory containing ConstraintTemplates, Constraints, and resources to benchmark. Can be specified multiple times.")
	Cmd.Flags().StringArrayVarP(&flagImages, flagNameImage, "i", []string{},
		"a URL to an OCI image containing policies. Can be specified multiple times.")
	Cmd.Flags().StringVarP(&flagTempDir, flagNameTempDir, "d", "",
		"temporary directory to download and unpack images to.")
	Cmd.Flags().StringVarP(&flagEngine, flagNameEngine, "e", string(bench.EngineCEL),
		fmt.Sprintf("policy engine to benchmark. One of: %s|%s|%s", bench.EngineRego, bench.EngineCEL, bench.EngineAll))
	Cmd.Flags().IntVarP(&flagIterations, flagNameIterations, "n", 1000,
		"number of benchmark iterations to run. Use at least 1000 for meaningful P99 metrics.")
	Cmd.Flags().IntVar(&flagWarmup, flagNameWarmup, 10,
		"number of warmup iterations before measurement.")
	Cmd.Flags().IntVarP(&flagConcurrency, flagNameConcurrency, "c", 1,
		"number of concurrent goroutines for reviews. Higher values simulate realistic webhook load.")
	Cmd.Flags().StringVarP(&flagOutput, flagNameOutput, "o", "table",
		"output format. One of: table|json|yaml")
	Cmd.Flags().BoolVar(&flagStats, flagNameStats, false,
		"gather detailed statistics from the constraint framework.")
	Cmd.Flags().BoolVar(&flagMemory, flagNameMemory, false,
		"enable memory profiling to track allocations per review.")
	Cmd.Flags().StringVar(&flagSave, flagNameSave, "",
		"save benchmark results to this file for future comparison (supports .json and .yaml).")
	Cmd.Flags().StringVar(&flagCompare, flagNameCompare, "",
		"compare results against a baseline file (supports .json and .yaml).")
	Cmd.Flags().Float64Var(&flagThreshold, flagNameThreshold, 10.0,
		"regression threshold percentage for comparison. Exit code 1 if exceeded.")
	Cmd.Flags().DurationVar(&flagMinThreshold, flagNameMinThreshold, 0,
		"minimum absolute latency difference to consider a regression (e.g., 1ms). Prevents false positives on fast policies where small absolute changes appear as large percentages.")
}

func run(_ *cobra.Command, _ []string) {
	_ = "STUB: not implemented"
	// Validate engine flag
	return
}

// Validate output format

// Validate inputs

// Warn if warmup exceeds iterations (likely user error)

// Validate baseline file exists before running expensive benchmark

// Run benchmark

// Format and print results

// Save results if requested

// Compare against baseline if requested

// Check if any comparison failed

func parseEngine(s string) (bench.Engine, error) {
	_ = "STUB: not implemented"
	return *new(bench.Engine), nil
}
