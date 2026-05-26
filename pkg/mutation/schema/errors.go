package schema

import (
	"github.com/open-policy-agent/gatekeeper/v3/pkg/util"
)

// ErrNilMutator reports that a method which expected an actual Mutator was
// a nil pointer.
const ErrNilMutator = util.Error("attempted to add nil mutator")

func NewErrConflictingSchema(ids IDSet) error { _ = "STUB: not implemented"; return nil }

const ErrConflictingSchemaType = "ErrConflictingSchema"

// ErrConflictingSchema reports that adding a Mutator to the DB resulted in
// conflicting implicit schemas.
type ErrConflictingSchema struct {
	Conflicts IDSet
}

func (e ErrConflictingSchema) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrConflictingSchema) Is(other error) bool { _ = "STUB: not implemented"; return false }
