package test

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "test",
	Short: "Test that the provided SyncSet(s) and/or Config contain the GVKs required by the input templates.",
	Run:   run,
}

var (
	flagFilenames       []string
	flagImages          []string
	flagOmitGVKManifest bool
	flagTempDir         string
)

const (
	flagNameFilename = "filename"
	flagNameImage    = "image"
	flagNameForce    = "force-omit-gvk-manifest"
	flagNameTempDir  = "tempdir"
)

func init() {
	Cmd.Flags().StringArrayVarP(&flagFilenames, flagNameFilename, "n", []string{}, "a file or directory containing Kubernetes resources. Can be specified multiple times.")
	Cmd.Flags().StringArrayVarP(&flagImages, flagNameImage, "i", []string{}, "a URL to an OCI image containing policies. Can be specified multiple times.")
	Cmd.Flags().BoolVarP(&flagOmitGVKManifest, flagNameForce, "f", false, "Do not require a GVK manifest; if one is not provided, assume all GVKs listed in the requirements "+
		"and configs are supported by the cluster under test. If this assumption isn't true, the given config may cause errors or templates may not be enforced correctly even after passing this test.")
	Cmd.Flags().StringVarP(&flagTempDir, flagNameTempDir, "d", "", fmt.Sprintf("Specifies the temporary directory to download and unpack images to, if using the --%s flag. Optional.", flagNameImage))
}

func run(_ *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

func resultsToString[T any](results map[string]T) string { _ = "STUB: not implemented"; return "" }
