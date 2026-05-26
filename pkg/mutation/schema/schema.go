package schema

import (
	"sync"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/parser"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	"k8s.io/apimachinery/pkg/runtime/schema"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// MutatorWithSchema is a mutator exposing the implied
// schema of the target object.
type MutatorWithSchema interface {
	types.Mutator

	// SchemaBindings returns the set of GVKs this Mutator applies to.
	SchemaBindings() []schema.GroupVersionKind

	// TerminalType specifies the inferred type of the last node in a path
	TerminalType() parser.NodeType
}

var log = logf.Log.WithName("mutation_schema")

// New returns a new schema database.
func New() *DB { _ = "STUB: not implemented"; return nil }

// DB is a database that caches all the implied schemas.
// Returns an error when upserting a mutator which conflicts with the existing ones.
//
// Mutators implicitly define a part of the schema of the object they intend
// to mutate. For example, modifying `spec.containers[name: foo].image` implies
// that:
// - spec is an object
// - containers is a list
// - image exists, but no type information
//
// If another mutator on the same GVK declares that it modifies
// `spec.containers.image`, then we know we have a contradiction as that path
// implies containers is an object.
//
// Conflicting schemas are stored within DB. HasConflicts returns true for
// the IDs of Mutators with conflicting schemas until Remove() is called on
// the Mutators which conflict with the ID.
type DB struct {
	mutex sync.RWMutex

	// cachedMutators is a cache of all seen Mutators.
	cachedMutators map[types.ID]MutatorWithSchema

	// schemas are the per-GVK implicit schemas.
	schemas map[schema.GroupVersionKind]*node

	conflicts IDSet
}

// Upsert inserts or updates the given mutator.
// Returns an error if the implicit schema in mutator conflicts with any
// mutators previously added.
//
// Schema conflicts are only detected using mutator.Path() - DB does not check
// that assigned types are compatible. For example, one Mutator might assign
// a string to a path and another might assign a list.
func (db *DB) Upsert(mutator MutatorWithSchema) error { _ = "STUB: not implemented"; return nil }

func (db *DB) upsert(mutator MutatorWithSchema) error { _ = "STUB: not implemented"; return nil }

// We've already added a Mutator which has the same path and bindings, so
// there's nothing to do.

// Adding this mutator had schema conflicts with another, so return an error.

// Remove removes the mutator with the given id from the db.
func (db *DB) Remove(id types.ID) { _ = "STUB: not implemented"; return }

func (db *DB) remove(id types.ID) { _ = "STUB: not implemented"; return }

// This means there's a bug in the schema code. This means a mutator
// is bound to this gvk with a previous call to upsert, but for some
// reason there is no corresponding schema.

// The root node is no longer referenced by any mutators.

// Remove the mutator from the cache.

// This ID's conflicts are resolved since the ID no longer exists.

// Check existing conflicts.
// TODO: Determine if there's a way of narrowing the list of potential conflicts.

// Check all current conflicts to see if they have been resolved.
// This optimizes for calls to HasConflicts()

// Only remove the conflict if all types now report there is no conflict
// at the path.

// HasConflicts returns true if the Mutator of the passed ID has been upserted
// in DB and has conflicts with another Mutator. Returns false if the Mutator
// does not exist.
func (db *DB) HasConflicts(id types.ID) bool { _ = "STUB: not implemented"; return false }

func (db *DB) GetConflicts(id types.ID) IDSet { _ = "STUB: not implemented"; return *new(IDSet) }
