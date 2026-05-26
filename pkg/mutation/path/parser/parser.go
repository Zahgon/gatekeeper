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

import (
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/token"
)

type parser struct {
	input     string
	scanner   *token.Scanner
	curToken  token.Token
	peekToken token.Token
	err       error
}

// Parse parses the provided input and returns an abstract representation if successful.
func Parse(input string) (Path, error) { _ = "STUB: not implemented"; return *new(Path), nil }

func newParser(input string) *parser { _ = "STUB: not implemented"; return nil }

// next advances to the next token in the stream.
func (p *parser) next() { _ = "STUB: not implemented"; return }

// expect returns whether the next token matches our expectation,
// and if so advances to that token.
// Otherwise returns false and doesn't advance.
func (p *parser) expect(t token.Type) bool { _ = "STUB: not implemented"; return false }

// expectPeek returns whether the next token matches our expectation.
// The current token is not advanced either way.
func (p *parser) expectPeek(t token.Type) bool { _ = "STUB: not implemented"; return false }

func (p *parser) Parse() (Path, error) { _ = "STUB: not implemented"; return *new(Path), nil }

// Check for optional listSpec operator

// Advance past separator if needed and ensure no unexpected tokens follow.
// NOTE: expect() advances the current position if the next token is a match.

// block trailing separators

// Skip past the separator

// Allowed. Loop will exit.

// parseList tries to parse the current position as List match node, e.g. [key: val]
// returns nil if it cannot be parsed as a List.
func (p *parser) parseList() Node {
	_ = "STUB: not implemented"

	// keyField is required
	return *new(Node)
}

func (p *parser) parseObject() Node { _ = "STUB: not implemented"; return *new(Node) }

func (p *parser) setError(err error) {
	_ = "STUB: not implemented"
	// Support only the first error for now
	return
}

// parseInt64 will return the int64 representation of the decimal encoded in the string s.
// This function was written because strconv.ParseInt() parses octal and hexadecimal representations
// which we are not supporting in our syntax.
func parseInt64(s string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
