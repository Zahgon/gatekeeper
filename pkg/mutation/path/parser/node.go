/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parser

type NodeType string

const (
	// ListNode is an array element of a path.
	ListNode NodeType = "List"
	// ObjectNode is the final Node in a path, what is being referenced.
	ObjectNode NodeType = "Object"
)

type Node interface {
	Type() NodeType
	DeepCopyNode() Node
	// String converts the Node into an equivalent String representation.
	// Calling Parse on the result yields an equivalent Node, but may differ in
	// structure if the Node is a Path containing Path Nodes.
	String() string
}

// Path represents an entire parsed path specification.
type Path struct {
	Nodes []Node
}

func (r Path) DeepCopy() Path { _ = "STUB: not implemented"; return *new(Path) }

func (r Path) String() string { _ = "STUB: not implemented"; return "" }

// No leading separator, and no separators before List Nodes.

type Object struct {
	Reference string
}

var _ Node = Object{}

func (o Object) Type() NodeType { _ = "STUB: not implemented"; return *new(NodeType) }

func (o Object) DeepCopyNode() Node { _ = "STUB: not implemented"; return *new(Node) }

func (o Object) DeepCopy() Object { _ = "STUB: not implemented"; return *new(Object) }

func (o Object) String() string { _ = "STUB: not implemented"; return "" }

type List struct {
	KeyField string
	KeyValue interface{}
	Glob     bool
}

var _ Node = List{}

func (l List) Type() NodeType { _ = "STUB: not implemented"; return *new(NodeType) }

func (l List) DeepCopyNode() Node { _ = "STUB: not implemented"; return *new(Node) }

func (l List) DeepCopy() List { _ = "STUB: not implemented"; return *new(List) }

// KeyValue (interface{}) will be one of: [string, int, nil]

func (l List) String() string { _ = "STUB: not implemented"; return "" }

// Represents an improperly specified List node.

// quote optionally adds double quotes around the passed string if needed.
// Quotes are needed for:
//   - Strings containing whitespace, quotes, or other "ambiguous" characters that will
//     be tokenized as non-strings and need escaping.
//   - Strings starting digits, that would otherwise be tokenized as an integer
//   - Empty strings
func quote(s string) string { _ = "STUB: not implemented"; return "" }

// Using fmt.Sprintf with %q converts whitespace to escape sequences, and we
// don't want that.
