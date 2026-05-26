package process

import (
	"sync"

	configv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/config/v1alpha1"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/wildcard"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Process indicates the Gatekeeper component from which the resource will be excluded.
type Process string

// The set of defined Gatekeeper processes.
const (
	Audit    = Process("audit")
	Sync     = Process("sync")
	Webhook  = Process("webhook")
	Mutation = Process("mutation-webhook")
	Star     = Process("*")
)

type Excluder struct {
	mux                sync.RWMutex
	excludedNamespaces map[Process]map[wildcard.Wildcard]bool
}

var allProcesses = []Process{
	Audit,
	Webhook,
	Mutation,
	Sync,
}

var processExcluder = &Excluder{
	excludedNamespaces: make(map[Process]map[wildcard.Wildcard]bool),
}

func Get() *Excluder { _ = "STUB: not implemented"; return nil }

func New() *Excluder { _ = "STUB: not implemented"; return nil }

func (s *Excluder) Add(entry []configv1alpha1.MatchEntry) { _ = "STUB: not implemented"; return }

// adding excluded namespace to all processes for "*"

func (s *Excluder) Replace(new *Excluder) {
	_ = "STUB: not implemented" // nolint:revive
	return
}

func (s *Excluder) Equals(new *Excluder) bool {
	_ = "STUB: not implemented" // nolint:revive
	return false
}

// EqualsForProcess checks if the excluded namespaces for a specific process are equal.
func (s *Excluder) EqualsForProcess(process Process, new *Excluder) bool {
	_ = "STUB: not implemented" // nolint:revive
	return false
}

func (s *Excluder) IsNamespaceExcluded(process Process, obj client.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetExcludedNamespaces returns a list of excluded namespace patterns for the given process.
func (s *Excluder) GetExcludedNamespaces(process Process) []string {
	_ = "STUB: not implemented"
	return nil
}

func exactOrWildcardMatch(boolMap map[wildcard.Wildcard]bool, ns string) bool {
	_ = "STUB: not implemented"
	return false
}
