package reader

import (
	"github.com/open-policy-agent/gatekeeper/v3/pkg/gator"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var allowedExtensions = []string{gator.ExtYAML, gator.ExtYML, gator.ExtJSON}

func ReadSources(filenames []string, images []string, tempDir string) ([]*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil,

		// Read from --filename flag
		nil
}

// Read from --image flag

// Read stdin

func readImage(image string, tempDir string) ([]*source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readImages(images []string, tempDir string) ([]*source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readFile(filename string) ([]*source, error) { _ = "STUB: not implemented"; return nil, nil }

func readFiles(filenames []string) ([]*source, error) { _ = "STUB: not implemented"; return nil, nil }

// verifyFile checks that the filenames aren't themselves disallowed extensions.
// This yields a much better user experience when the user mis-uses the
// --filename flag.
func verifyFile(filename string) error {
	_ = "STUB: not implemented"
	// make sure it's a file, not a directory
	return nil
}

func readStdin() ([]*unstructured.Unstructured, error) { _ = "STUB: not implemented"; return nil, nil }

// check if data is being piped or redirected to stdin

func expandDirectories(filenames []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filesBelow walks the filetree from startPath and below, collecting a list of
// all the filepaths.  Directories are excluded.
func filesBelow(startPath string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// only add files to the normalized output

// make sure the file extension is valid

func allowedExtension(path string) bool { _ = "STUB: not implemented"; return false }

func sourcesToUnstruct(sources []*source) []*unstructured.Unstructured {
	_ = "STUB: not implemented"
	return nil
}
