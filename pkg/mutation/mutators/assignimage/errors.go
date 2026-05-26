package assignimage

type baseError struct {
	s string
}

func (e baseError) Error() string {
	_ = "STUB: not implemented"

	// Component field (domain|path|tag) errors.
	return ""
}

type invalidDomainError struct{ baseError }

type (
	invalidPathError       struct{ baseError }
	invalidTagError        struct{ baseError }
	missingComponentsError struct{ baseError }
	domainLikePathError    struct{ baseError }
)

// Location field errors.
type (
	listTerminalError struct{ baseError }
	metadataRootError struct{ baseError }
)

func newInvalidDomainError(domain string) invalidDomainError {
	_ = "STUB: not implemented"
	return *new(invalidDomainError)
}

func newInvalidPathError(path string) invalidPathError {
	_ = "STUB: not implemented"
	return *new(invalidPathError)
}

func newInvalidTagError(tag string) invalidTagError {
	_ = "STUB: not implemented"
	return *new(invalidTagError)
}

func newMissingComponentsError() missingComponentsError {
	_ = "STUB: not implemented"
	return *new(missingComponentsError)
}

func newDomainLikePathError(path string) domainLikePathError {
	_ = "STUB: not implemented"
	return *new(domainLikePathError)
}

func newListTerminalError(name string) listTerminalError {
	_ = "STUB: not implemented"
	return *new(listTerminalError)
}

func newMetadataRootError(name string) metadataRootError {
	_ = "STUB: not implemented"
	return *new(metadataRootError)
}
