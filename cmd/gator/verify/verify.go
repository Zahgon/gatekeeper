package verify

import (
	"context"
	"io/fs"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/gator/verify"
	"github.com/spf13/cobra"
)

const (
	examples = `# Run all tests in label-tests.yaml
gator verify label-tests.yaml

# Run all tests whose names contain "forbid-labels".
gator verify tests/... --run forbid-labels//

# Run all cases whose names contain "nginx-deployment".
gator verify tests/... --run //nginx-deployment

# Run all cases whose names exactly match "nginx-deployment".
gator verify tests/... --run '//^nginx-deployment$'

# Run all cases that are either named "forbid-labels" or are
# in tests named "forbid-labels".
gator verify tests/... --run '^forbid-labels$'`
)

var (
	run              string
	verbose          bool
	includeTrace     bool
	flagEnableK8sCel bool
)

func init() {
	Cmd.Flags().StringVarP(&run, "run", "r", "",
		`regular expression which filters tests to run by name`)
	Cmd.Flags().BoolVarP(&verbose, "verbose", "v", false,
		`print extended test output`)
	Cmd.Flags().BoolVarP(&includeTrace, "trace", "t", false,
		`include a trace for the underlying constraint framework evaluation`)
	Cmd.Flags().BoolVarP(&flagEnableK8sCel, "enable-k8s-native-validation", "", true,
		`Beta: enable the validating admission policy driver`)
}

// Cmd is the gator verify subcommand.
var Cmd = &cobra.Command{
	Use:     "verify path [--run=name]",
	Short:   "verify suites of tests on Gatekeeper Constraints",
	Example: examples,
	Args:    cobra.ExactArgs(1),
	RunE:    runE,
}

func runE(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Convert path to be absolute. Allowing for relative and absolute paths
// everywhere in the code leads to unnecessary complexity, so the first
// thing we do on encountering a path is to convert it to an absolute path.

// Create the base file system. We use fs.FS rather than direct calls to
// os.ReadFile or filepath.WalkDir to make testing easier and keep logic
// os-independent.

// fs.FS does not allow the drive prefix for absolute Windows paths.

// filepath.Abs strips "/..." from the end of Windows paths, so check the original string.

func runSuites(ctx context.Context, fileSystem fs.FS, suites []*verify.Suite, filter verify.Filter) error {
	_ = "STUB: not implemented"
	return nil
}

// At least one test failed or there was a problem executing tests in at
// least one file.

func getFS(path string) fs.FS {
	_ = "STUB: not implemented"
	// TODO(#1397): Check that this produces the correct file system string on
	//
	//	Windows. We may need to add a trailing `/` for fs.filesystem to function properly.
	return *new(fs.FS)
}

// We are running on a unix-like filesystem without volume names, so the
// file system root is `/`.
