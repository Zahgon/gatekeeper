package assignimage

import (
	"regexp"
)

var (
	// We perform validation on the components of an image string to ensure that
	// the user cannot define a mutator which does not converge. This would
	// otherwise be possible by injecting tokens we use to split an image string,
	// [@:/], into components that would cause that component to be split the next
	// time the mutation is applied and "leak" to its neighbor. Some validation is
	// done as regex on individual components, and other validation which looks at
	// multiple components together is done in code. All validation for domain and
	// tag must be put in validateDomain and validateTag respectively.

	// domainRegexp defines a schema for a domain component.
	domainRegexp = regexp.MustCompile(`(^\w[\w\-_]*\.[\w\-_\.]*[\w](:\d+)?$)|(^localhost(:\d+)?$)`)

	// pathRegexp defines a schema for a location component. It follows the convention
	// specified in the docker distribution reference. The regex  restricts
	// location-components to start with an alphanumeric character, with following
	// parts able to be separated by a separator (one period, one or two
	// underscore and multiple dashes).
	pathRegexp = regexp.MustCompile(`^[a-z0-9]+(?:(?:(?:[._/]|__|[-]*)[a-z0-9]+)+)?`)

	// tagRegexp defines a schema for a tag component. It must start with `:` or `@`.
	tagRegexp = regexp.MustCompile(`(^:[\w][\w.-]{0,127}$)|(^@[A-Za-z][A-Za-z0-9]*([-_+.][A-Za-z][A-Za-z0-9]*)*[:][0-9A-Fa-f]{32,}$)`)
)

type image struct {
	domain string
	path   string
	tag    string
}

func mutateImage(domain, path, tag, mutableImgRef string) string {
	_ = "STUB: not implemented"
	return ""
}

func newImage(imageRef string) image { _ = "STUB: not implemented"; return *new(image) }

// splitTag separates the path and tag components from a string.
func splitTag(remainder string) (string, string) { _ = "STUB: not implemented"; return "", "" }

func (img image) newMutatedImage(domain, path, tag string) image {
	_ = "STUB: not implemented"
	return *new(image)
}

// ignoreUnset returns `new` if `new` is set, otherwise it returns `old`.
func ignoreUnset(old, new string) string {
	_ = "STUB: not implemented" // nolint:revive
	return ""
}

func (img image) fullRef() string { _ = "STUB: not implemented"; return "" }

func splitDomain(name string) (domain, remainder string) { _ = "STUB: not implemented"; return "", "" }

func validateDomain(domain string) error { _ = "STUB: not implemented"; return nil }

// The error below should theoretically be unreachable, as the regex
// validation should preclude this from happening. This check is included
// anyway to prevent code drift, and ensure that if a domain is validated
// it can also be recognized as a domain.

func validateTag(tag string) error { _ = "STUB: not implemented"; return nil }

// This error should never happen because the regex above prevents it, but the
// check is included to prevent drift. Splitting the tag should return itself,
// and splitting a valid tag should never return a path.

func validateImageParts(domain, path, tag string) error { _ = "STUB: not implemented"; return nil }

// match the whole string for path (anchoring with `$` is tricky here)

// Check if the path looks like a domain string, and the domain is not set.
// This prevents part of the path field from "leaking" to the domain, causing
// non convergent behavior.
// For example, suppose: domain="", path="gcr.io/repo", tag=""
// Suppose no value is currently set on the mutable, so the result is
// just "gcr.io/repo". When this value mutated again, "gcr.io" is parsed into
// the domain component, so the result would be "gcr.io/gcr.io/repo" and so on.
