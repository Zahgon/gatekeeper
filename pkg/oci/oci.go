package oci

const tempFilePrefix = "gator-bundle-"

// PullImage pulls an OCI image at `imgURL` into a temporary directory with a
// random name, created under the path `tempDir`. If `tempDir` is empty, a
// default path from os.TempDir() is used. This func returns the directory path
// that the image was pulled to, a handler function to clean up the directory
// after it has been read, and an error (if any).
func PullImage(imgURL string, tempDir string) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Preserve the oras v1 behavior for local test registries on loopback hosts.

// shouldUsePlainHTTP returns true when the registry host is a loopback
// address (localhost, 127.x.x.x, ::1), which typically serves plain HTTP.
func shouldUsePlainHTTP(registryHost string) bool { _ = "STUB: not implemented"; return false }

// Strip IPv6 brackets that remain when no port is present (e.g. "[::1]").
