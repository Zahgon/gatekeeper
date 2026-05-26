package version

const (
	gatorState = "beta"
	unknown    = "unknown"
)

// Version is the gatekeeper version.
var Version string

// GetUserAgent returns Gatekeeper and Gator version information.
func GetUserAgent(name string) string { _ = "STUB: not implemented"; return "" }

// OPA and Frameworks version used by Gatekeeper and Gator

// if LDFLAGS are not set, use revision info
