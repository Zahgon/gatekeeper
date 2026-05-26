package verify

import (
	"errors"
	"io/fs"
)

var (
	// ErrNoFileSystem means a method which expects a filesystem got nil
	// instead. This is likely a bug in the code.
	ErrNoFileSystem = errors.New("no filesystem")
	// ErrNoTarget indicates that the user did not specify a target directory or
	// file.
	ErrNoTarget = errors.New("target not specified")
	// ErrUnsupportedExtension indicates that a user attempted to run tests in
	// a file type which is not supported.
	ErrUnsupportedExtension = errors.New("unsupported extension")
	// ErrNotADirectory indicates that a user is mistakenly attempting to
	// perform a directory-only action on a file (for example, recursively
	// traversing it).
	ErrNotADirectory = errors.New("not a directory")
)

const (
	// Group is the API Group for Test YAML objects.
	Group = "test.gatekeeper.sh"
	// Kind is the Kind for Suite YAML objects.
	Kind = "Suite"
)

// ReadSuites returns the set of test Suites selected by path.
//
//  1. If path is a path to a Suite, parses and returns the Suite.
//  2. If the path is a directory and recursive is false, returns only the Suites
//     defined in that directory.
//  3. If the path is a directory and recursive is true returns all Suites in that
//     directory and its subdirectories.
//
// Returns an error if:
// - path is a file that does not define a Suite
// - any matched files containing Suites are not parseable.
func ReadSuites(f fs.FS, target, originalPath string, recursive bool) ([]*Suite, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// target is a file.

// target is a file, but the user specified it should be traversed recursively.

// target is a directory, but user did not specify to test subdirectories.

// target is a directory, and the user specified it should be traversed recursively.

// readSuites reads the passed set of files into Suites on the given filesystem.
// originalPath argument is used to construct paths relative to the original input
// path from the traversed file system walks.
func readSuites(f fs.FS, files []string, originalPath string) ([]*Suite, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// trim any prefixes like "/", "./" or "../" in order for the
// .Cut call below to actually work with the absolute path
// contained in the file var.

// Ensure Suites are returned in a deterministic order.

// fileList is a convenience type for breaking apart and deduplicating code
// related to collecting the set of files which may contain Suites.
type fileList []string

func (l *fileList) addFile(target string) error {
	_ = "STUB: not implemented"
	// target is a file.
	return nil
}

func (l *fileList) addDirectory(f fs.FS, target string) error {
	_ = "STUB: not implemented"
	// target is a directory, but user did not specify to test subdirectories.
	return nil
}

func (l *fileList) walkEntry(path string, d fs.DirEntry, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func isYAMLFile(d fs.DirEntry) bool { _ = "STUB: not implemented"; return false }

func readSuite(f fs.FS, path string) (*Suite, error) { _ = "STUB: not implemented"; return nil, nil }

// Not a test file; we can safely ignore this.
