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

package token

import (
	"errors"
)

const eof = rune(-1)

// Base errors for scanning strings.
var (
	ErrUnterminatedString = errors.New("unterminated string")
	ErrInvalidCharacter   = errors.New("invalid character")
)

type Scanner struct {
	input   string
	pos     int // Current position
	readPos int // Next position to read
	ch      rune
	err     error // Last error if any
}

func NewScanner(input string) *Scanner { _ = "STUB: not implemented"; return nil }

func (s *Scanner) Next() Token { _ = "STUB: not implemented"; return *new(Token) }

// A match on these first set of cases leaves s.ch positioned at the next character to process.

// Any of these cases require a subsequent call to s.read() (below) to position the next character.

// default: current character is invalid at this location

// Make progress

// read consumes the next rune and advances.
func (s *Scanner) read() rune { _ = "STUB: not implemented"; return 0 }

// Mark last read position
// Advance for next read

// readString consumes a string token.
func (s *Scanner) readString() (string, error) {
	_ = "STUB: not implemented"
	// Will be ' or "
	return "", nil
}

// String terminated

// Escaped character

// Unterminated string

func (s *Scanner) readIdent() (string, error) { _ = "STUB: not implemented"; return "", nil }

// readInt scans a (positive) integer. Signs are not supported.
func (s *Scanner) readInt() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *Scanner) setError(err error) { _ = "STUB: not implemented"; return }

// isSpace returns true if the passed rune is a supported whitespace character.
func isSpace(r rune) bool { _ = "STUB: not implemented"; return false }

func isAlphaNum(r rune) bool { _ = "STUB: not implemented"; return false }

func isDigit(r rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) skipWhitespace() { _ = "STUB: not implemented"; return }

type ScanError struct {
	Inner    error
	Position int
}

func (e ScanError) Error() string { _ = "STUB: not implemented"; return "" }

// Unwrap allows errors.Is() to inspect the underlying error.
func (e ScanError) Unwrap() error { _ = "STUB: not implemented"; return nil }
