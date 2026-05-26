package test

import (
	"fmt"

	"github.com/open-policy-agent/frameworks/constraint/pkg/instrumentation"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/gator/test"
	"github.com/spf13/cobra"
)

const (
	examples = `# test a manifest containing Kubernetes objects, Constraint Templates, and Constraints
gator test --filename="manifest.yaml"

# test a directory
gator test --filename="config-and-policies/"

# Use multiple inputs
gator test --filename="manifest.yaml" --filename="templates-and-constraints/"

# Receive input from stdin
cat manifest.yaml | gator test

# Output structured violations data
gator test --filename="manifest.yaml" --output=json`
)

var Cmd = &cobra.Command{
	Use:     "test",
	Short:   "test evaluates resources against policies as defined by constraint templates and constraints.",
	Example: examples,
	Run:     run,
	Args:    cobra.NoArgs,
}

var (
	flagFilenames    []string
	flagOutput       string
	flagIncludeTrace bool
	flagGatherStats  bool
	flagImages       []string
	flagTempDir      string
	flagEnableK8sCel bool
	flagDenyOnly     bool
	flagVerbose      bool
)

const (
	flagNameFilename = "filename"
	flagNameOutput   = "output"
	flagNameImage    = "image"
	flagNameTempDir  = "tempdir"
	flagNameVerbose  = "verbose"

	stringJSON          = "json"
	stringYAML          = "yaml"
	stringHumanFriendly = "default"

	fourSpaceTab = "    "
)

func init() {
	Cmd.Flags().StringArrayVarP(&flagFilenames, flagNameFilename, "f", []string{}, "a file or directory containing Kubernetes resources.  Can be specified multiple times.")
	Cmd.Flags().StringVarP(&flagOutput, flagNameOutput, "o", "", fmt.Sprintf("Output format.  One of: %s|%s.", stringJSON, stringYAML))
	Cmd.Flags().BoolVarP(&flagIncludeTrace, "trace", "t", false, "include a trace for the underlying Constraint Framework evaluation.")
	Cmd.Flags().BoolVarP(&flagGatherStats, "stats", "", false, "include performance stats returned from the Constraint Framework.")
	Cmd.Flags().BoolVarP(&flagEnableK8sCel, "enable-k8s-native-validation", "", true, "enable the validating admission policy driver")
	Cmd.Flags().StringArrayVarP(&flagImages, flagNameImage, "i", []string{}, "a URL to an OCI image containing policies. Can be specified multiple times.")
	Cmd.Flags().StringVarP(&flagTempDir, flagNameTempDir, "d", "", fmt.Sprintf("Specifies the temporary directory to download and unpack images to, if using the --%s flag. Optional.", flagNameImage))
	Cmd.Flags().BoolVarP(&flagDenyOnly, "deny-only", "", false, "output only denied constraints")
	Cmd.Flags().BoolVarP(&flagVerbose, flagNameVerbose, "v", false, "print extended test output")
}

func run(_ *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

// Whether or not we return non-zero depends on whether we have a `deny`
// enforcementAction on one of the violated constraints

func formatOutput(flagOutput string, allResults []*test.GatorResult, stats []*instrumentation.StatsEntry) string {
	_ = "STUB: not implemented"
	return ""
}

func enforceableFailures(results []*test.GatorResult) bool { _ = "STUB: not implemented"; return false }

func enforceableFailure(result *test.GatorResult) bool { _ = "STUB: not implemented"; return false }
