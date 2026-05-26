// operations stores the operations assigned to the pod via the --operation flag
// It is meant to be read-only and only set once, when flags are parsed.

package operations

import (
	"flag"
	"sync"
)

type Operation string

// All defined Operations.
const (
	Audit              = Operation("audit")
	MutationController = Operation("mutation-controller")
	MutationStatus     = Operation("mutation-status")
	MutationWebhook    = Operation("mutation-webhook")
	Status             = Operation("status")
	Webhook            = Operation("webhook")
	Generate           = Operation("generate")
)

var (
	// allOperations is a list of all possible Operations that can be assigned to
	// a pod. It is NOT intended to be mutated.
	allOperations = []Operation{
		Audit,
		Generate,
		MutationController,
		MutationStatus,
		MutationWebhook,
		Status,
		Webhook,
	}

	operationsMtx sync.RWMutex
	operations    = newOperationSet()
)

type opSet struct {
	validOperations    map[Operation]bool
	assignedOperations map[Operation]bool
	assignedStringList []string // cached serialization of the opSet
	initialized        bool
}

var _ flag.Value = &opSet{}

func newOperationSet() *opSet { _ = "STUB: not implemented"; return nil }

// default to all operations enabled

func (l *opSet) String() string { _ = "STUB: not implemented"; return "" }

func (l *opSet) Set(s string) error {
	_ = "STUB: not implemented"

	// When the user sets an explicit value, start fresh (no default all-values)
	return nil
}

func init() {
	flag.Var(operations, "operation", "The operation to be performed by this instance. e.g. audit, webhook. This flag can be declared more than once. Omitting will default to supporting all operations.")
}

// IsAssigned returns true when the provided operation is assigned to the pod.
func IsAssigned(op Operation) bool { _ = "STUB: not implemented"; return false }

// AssignedStringList returns a list of all operations assigned to the pod
// as a sorted list of strings.
func AssignedStringList() []string {
	_ = "STUB: not implemented"
	// Use a read lock so we can exit early without potentially having multiple
	// threads try to write this simultaneously.
	return nil
}

// Verify the list hasn't been set since we last checked.

// HasValidationOperations returns `true` if there
// are any operations that would require a constraint or template controller
// or a sync controller.
func HasValidationOperations() bool { _ = "STUB: not implemented"; return false }
