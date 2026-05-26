package aggregator

import (
	gosync "sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Key defines a type, identifier tuple to store
// in the GVKAggregator.
type Key struct {
	// Source specifies the type of the source object.
	Source string
	// ID specifies the name of the instance of the source object.
	ID string
}

func NewGVKAggregator() *GVKAgreggator { _ = "STUB: not implemented"; return nil }

// GVKAgreggator is an implementation of a bi directional map
// that stores associations between Key K and GVKs and reverse associations
// between GVK g and Keys.
type GVKAgreggator struct {
	mu gosync.RWMutex

	// store keeps track of associations between a Key type and a set of GVKs.
	store map[Key]map[schema.GroupVersionKind]struct{}
	// reverseStore keeps track of associations between a GVK and the set of Key types
	// that references the GVK in the store map above. It is useful to have reverseStore
	// in order for IsPresent() and ListGVKs() to run in optimal time.
	reverseStore map[schema.GroupVersionKind]map[Key]struct{}
}

// IsPresent returns true if the given gvk is present in the GVKAggregator.
func (b *GVKAgreggator) IsPresent(gvk schema.GroupVersionKind) bool {
	_ = "STUB: not implemented"
	return false
}

// Remove deletes any associations that Key k has in the GVKAggregator.
// For any GVK in the association k --> [GVKs], we also delete any associations
// between the GVK and the Key k stored in the reverse map.
func (b *GVKAgreggator) Remove(k Key) { _ = "STUB: not implemented"; return }

// Upsert stores an association between Key k and the list of GVKs
// and also the reverse association between each GVK passed in and Key k.
// Any old associations are dropped, unless they are included in the new list of
// GVKs.
func (b *GVKAgreggator) Upsert(k Key, gvks []schema.GroupVersionKind) {
	_ = "STUB: not implemented"
	return
}

// gvksToRemove contains old GKVs that are not included in the new gvks list

// protect against empty inputs

// add reverse links

// List returnes the gvk set for a given Key.
func (b *GVKAgreggator) List(k Key) []schema.GroupVersionKind {
	_ = "STUB: not implemented"
	return nil
}

// GVKs returns a list of all of the schema.GroupVersionKind that are aggregated.
func (b *GVKAgreggator) GVKs() []schema.GroupVersionKind { _ = "STUB: not implemented"; return nil }

func (b *GVKAgreggator) pruneReverseStore(gvks map[schema.GroupVersionKind]struct{}, k Key) {
	_ = "STUB: not implemented"
	return
}

// by definition, nothing to prune

// remove GVK from reverseStore if it's not referenced by any Key anymore.

func makeSet(gvks []schema.GroupVersionKind) map[schema.GroupVersionKind]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func unreferencedOldGVKsToPrune(newGVKs []schema.GroupVersionKind, oldGVKs map[schema.GroupVersionKind]struct{}) map[schema.GroupVersionKind]struct{} {
	_ = "STUB: not implemented"
	// deep copy oldGVKs
	return nil
}

// intersection: exclude the oldGVKs that are present in the new GVKs as well.

// don't prune what is being already added
