package expansion

import (
	"github.com/dominikbraun/graph"
	expansionunversioned "github.com/open-policy-agent/gatekeeper/v3/apis/expansion/unversioned"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type templateDB interface {
	upsert(*expansionunversioned.ExpansionTemplate) error
	remove(*expansionunversioned.ExpansionTemplate)
	templatesForGVK(gvk schema.GroupVersionKind) []*expansionunversioned.ExpansionTemplate
	getConflicts() IDSet
}

var _ templateDB = &db{}

type adjList map[schema.GroupVersionKind]IDSet

// hashID is passed to the graphing package.
var hashID = func(id TemplateID) string {
	return string(id)
}

type templateState struct {
	template     *expansionunversioned.ExpansionTemplate
	hasConflicts bool
}

type edge struct {
	x TemplateID
	y TemplateID
}

type db struct {
	store map[TemplateID]*templateState

	// graph stores a graph of ExpansionTemplate. A directed edge from template A
	// to B means that the template A's `generatedGVK` matches template B's `applyTo`.
	graph graph.Graph[string, TemplateID]

	// `matchers` and `generators` creates the necessary mappings to be able to
	// determine the inbound and outbound edges of a given template in O(1).
	// matchers is a mapping of GVKs to templates that match (applyTo) that GVK.
	matchers adjList
	// generators is a mapping of GVKs to templates that generate that GVK.
	generators adjList
}

func newDB() *db { _ = "STUB: not implemented"; return nil }

func (d *db) getConflicts() IDSet { _ = "STUB: not implemented"; return *new(IDSet) }

// handleAdd adds the template to the DB, returning true if a cycle was created.
// The template is added even if a cycle was found.
func (d *db) handleAdd(template *expansionunversioned.ExpansionTemplate) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// We should always remove the old template before handleAdd. If we
		// didn't, that's a bug.
		nil
}

// Update storage

// Update generators

// Update matchers

// Add vertex if DNE

// Add edges

func (d *db) edgesForTemplate(template *expansionunversioned.ExpansionTemplate) []edge {
	_ = "STUB: not implemented"
	return nil
}

// Add out-bound edges (from this template's generated GVK to other
// templates' matched GVKs)

// Add in-bound edges (from other templates' generated GVK to this template's
// matched GVK)

func (d *db) handleRemove(id TemplateID) {
	_ = "STUB: not implemented"
	// The template must exist. Existence checks should be done upstream.
	return
}

// Update storage

// Update generators

// Update matchers

// Remove edges

func (d *db) updateCycles() {
	_ = "STUB: not implemented"
	// First reset all conflicts
	return
}

// All strongly connect components containing more than 1 vertex are a cycle

func (d *db) upsert(template *expansionunversioned.ExpansionTemplate) error {
	_ = "STUB: not implemented"
	return nil
}

// If the new/updated template caused a cycle, or the previous template belonged
// to a cycle, then we need to re-check the graph for cycles

func (d *db) remove(template *expansionunversioned.ExpansionTemplate) {
	_ = "STUB: not implemented"
	return
}

// If the removed template was part of a cycle, we need to recheck the graph
// in case that cycle was resolved

func (d *db) templatesForGVK(gvk schema.GroupVersionKind) []*expansionunversioned.ExpansionTemplate {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check. In theory, this should never happen, but if it does, we
// can't recover.

func applyToGVKs(template *expansionunversioned.ExpansionTemplate) []schema.GroupVersionKind {
	_ = "STUB: not implemented"
	return nil
}
