package reader

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const warningMsg = "WARNING - Resource named %q (from %s) is already defined in %s"

type source struct {
	filename string
	image    string
	stdin    bool
	objs     []*unstructured.Unstructured
}

type gknn struct {
	schema.GroupKind
	name      string
	namespace string
}

type conflict struct {
	id gknn
	a  *source
	b  *source
}

func detectConflicts(sources []*source) []conflict { _ = "STUB: not implemented"; return nil }

func logConflict(c *conflict) { _ = "STUB: not implemented"; return }

// sourceDebugInfo returns a string identifying the source.
// For sources pulled from stdin: "stdin".
// For sources pulled from a file: "file: <filename>".
// For sources pulled from an image: "file: <filename>, image: <imgURL>".
func sourceDebugInfo(s *source) string { _ = "STUB: not implemented"; return "" }
