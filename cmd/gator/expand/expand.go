package expand

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	examples = `# expand resources in a manifest
gator expand --filename="manifest.yaml"

# expand a directory
gator expand --filename="config-and-policies/"

# Use multiple inputs
gator expand --filename="manifest.yaml" --filename="templates-and-constraints/"

# Output JSON to file
gator expand --filename="manifest.yaml" --format=json --outputfile=results.yaml`
)

var Cmd = &cobra.Command{
	Use:     "expand",
	Short:   "expand allow for testing generator resource expansion configs by expanding a generator resource and outputting the resultant resource(s)",
	Example: examples,
	Run:     run,
	Args:    cobra.ExactArgs(0),
}

var (
	flagFilenames []string
	flagFormat    string
	flagOutput    string
	flagImages    []string
	flagTempDir   string
)

const (
	flagNameFilename = "filename"
	flagNameFormat   = "format"
	flagNameOutput   = "outputfile"
	flagNameImage    = "image"
	flagNameTempDir  = "tempdir"

	stringJSON = "json"
	stringYAML = "yaml"

	delimeter = "---"
)

func init() {
	Cmd.Flags().StringArrayVarP(&flagFilenames, flagNameFilename, "n", []string{}, "a file or directory containing Kubernetes resources.  Can be specified multiple times.")
	Cmd.Flags().StringVarP(&flagFormat, flagNameFormat, "f", "", fmt.Sprintf("Output format.  One of: %s|%s.", stringJSON, stringYAML))
	Cmd.Flags().StringVarP(&flagOutput, flagNameOutput, "o", "", "Output file path. If the file already exists, it will be overwritten.")
	Cmd.Flags().StringArrayVarP(&flagImages, flagNameImage, "i", []string{}, "a URL to an OCI image containing policies. Can be specified multiple times.")
	Cmd.Flags().StringVarP(&flagTempDir, flagNameTempDir, "d", "", fmt.Sprintf("Specifies the temporary directory to download and unpack images to, if using the --%s flag. Optional.", flagNameImage))
}

func run(_ *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

// Sort resultants for deterministic output

func resourcetoYAMLString(resource *unstructured.Unstructured) string {
	_ = "STUB: not implemented"
	return ""
}

func resourceToJSONString(resource *unstructured.Unstructured) string {
	_ = "STUB: not implemented"
	return ""
}

func resourcesToString(resources []*unstructured.Unstructured, format string) string {
	_ = "STUB: not implemented"
	return ""
}

func sortUnstructs(objs []*unstructured.Unstructured) { _ = "STUB: not implemented"; return }
