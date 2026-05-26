package policy

// Exit codes for the gator policy command.
const (
	// ExitSuccess indicates successful execution.
	ExitSuccess = 0
	// ExitGeneralError indicates a general error (network, invalid args).
	ExitGeneralError = 1
	// ExitClusterError indicates a cluster error (not found, permission denied).
	ExitClusterError = 2
	// ExitConflictError indicates a conflict (resource exists, not managed by gator).
	ExitConflictError = 3
	// ExitPartialSuccess indicates partial success (some resources failed).
	ExitPartialSuccess = 4
)

// ExitError represents an error with an associated exit code.
type ExitError struct {
	Code    int
	Message string
}

// Error implements the error interface.
func (e *ExitError) Error() string {
	_ = "STUB: not implemented"

	// NewExitError creates a new ExitError with the given code and message.
	return ""
}

func NewExitError(code int, message string) *ExitError { _ = "STUB: not implemented"; return nil }

// NewGeneralError creates an ExitError for general errors.
func NewGeneralError(message string) *ExitError { _ = "STUB: not implemented"; return nil }

// NewClusterError creates an ExitError for cluster errors.
func NewClusterError(message string) *ExitError { _ = "STUB: not implemented"; return nil }

// NewConflictError creates an ExitError for conflict errors.
func NewConflictError(message string) *ExitError { _ = "STUB: not implemented"; return nil }

// NewPartialSuccessError creates an ExitError for partial success.
func NewPartialSuccessError(message string) *ExitError { _ = "STUB: not implemented"; return nil }
