package schema

import (
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/parser"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/util"
)

// node is an element of an implicit schema.
// Allows for the definition of overlapping schemas. See Add.
type node struct {
	// ReferencedBy tracks the Mutations which reference this part of the schema tree.
	ReferencedBy IDSet

	// Children is the set of child Nodes at this location in the schema.
	// Each node defines a distinct child definition. If multiple Nodes are defined
	// for the same child, then there is a schema conflict.
	Children map[string]map[parser.NodeType]node

	// MustTerminate tracks which mutators require their path to terminate, i.e.
	// the path must not be extended by any other mutation.
	MustTerminate IDSet
}

// Add inserts the provided path, linked to the given ID.
//
// Returns the set of conflicts detected while adding the path. Conflicts occur
// when elements with the same path have different types, for example:
//
// spec.containers[name: foo].image
// spec.containers.image
//
// If mustTerminate is true, any future paths that extend beyond this path will
// result in a conflict, for example:
//
// spec.foo <- must terminate at foo
// spec.foo.bar <- conflict since foo is a terminal node
//
// If the returned references is non-nil it contains at least two elements, one
// of which is the passed id.
func (n *node) Add(id types.ID, path []parser.Node, terminalType parser.NodeType, mustTerminate bool) IDSet {
	_ = "STUB: not implemented"
	return *new(IDSet)
}

// This node is referenced by the passed ID.

// Base case; there is no more path to validate.

// Initialize child within n.children.

// Add the remaining path to the appropriate child, collecting any conflicts
// found when adding it.

// Detect conflicts at this node.
// We know there is a conflict if there is a child with the same Key but a
// different type.

const ErrNotFound = util.Error("path not found")

// Remove removes the id and path from the tree.
// Panics if the ID is not defined or was Add()ed with a different path.
func (n *node) Remove(id types.ID, path []parser.Node, terminalType parser.NodeType, mustTerminate bool) {
	_ = "STUB: not implemented"
	// This ID no longer references this node.
	return
}

// No more path to remove.

// The child does not exist.

// No child of the key and type exists.
// This is how we detect that the path for id is incomplete. If the path
// were complete, the type of the child was known when Add()ed but not when
// Remove()d and is files as unknown.

// Delete the type from the child if it is no longer referenced.

// No references to this child of this type exist.

// Delete the child if it is no longer referenced.

func (n *node) conflicts(childKey string) IDSet {
	_ = "STUB: not implemented"
	return *

	// Count the number of distinct types with this key.
	new(IDSet)
}

// Nodes whose types we are unable to determine do not count against this
// check.

// If we don't know the type of a node, we assume it conflicts with nothing.

// There are conflicts if either:
// 1) there are more than one non-unknown types for the Child, or
// 2) the Child is a List and defines multiple keys.

// we can return here if one of the following is true:
// 1) we don't have to terminate at this node
// 2) we have to terminate at this node but we don't have grandchildren for this child key

// since we have to terminate at this node,
// the terminal node and its grandchildren are considered conflicts

// If more than 1 non-unknown types are declared, this node is part of a
// schema conflict.

// GetConflicts returns all conflicts along the passed path.
func (n *node) GetConflicts(path []parser.Node, terminalType parser.NodeType) []types.ID {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) getConflicts(path []parser.Node, terminalType parser.NodeType) IDSet {
	_ = "STUB: not implemented"
	return *new(IDSet)
}

// Path has not been added, so there can be no conflicts.

// Path has not been added, so there can be no conflicts.

// Count the number of distinct types with this key.

// merge inserts elements from `from` into `into`. Returns `into`, or a
// reference to a new map if `into` is nil.
func merge(into, from IDSet) IDSet { _ = "STUB: not implemented"; return *new(IDSet) }

// headType returns the type of the second Node, if it exists.
// This is essential for determining whether the current location in a schema
// path is a list.
func headType(path []parser.Node, terminalType parser.NodeType) parser.NodeType {
	_ = "STUB: not implemented"

	// Default to the terminal type, as we are at the last path node.
	return *new(parser.NodeType)
}

func (n *node) DeepCopy() *node { _ = "STUB: not implemented"; return nil }

// key extracts the unique identifier of the next element in the path from the
// given Node for use in the node tree.
func key(n parser.Node) string { _ = "STUB: not implemented"; return "" }
