package webhook

import (
	"context"
	"flag"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/util"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	exemptNamespace       = util.NewFlagSet()
	exemptNamespacePrefix = util.NewFlagSet()
	exemptNamespaceSuffix = util.NewFlagSet()
)

func init() {
	AddToManagerFuncs = append(AddToManagerFuncs, AddLabelWebhook)
	flag.Var(exemptNamespace, "exempt-namespace", "The specified namespace is allowed to set the admission.gatekeeper.sh/ignore label. To exempt multiple namespaces, this flag can be declared more than once.")
	flag.Var(exemptNamespacePrefix, "exempt-namespace-prefix", "A namespace with the specified prefix is allowed to set the admission.gatekeeper.sh/ignore label. To exempt multiple prefixes, this flag can be declared more than once.")
	flag.Var(exemptNamespaceSuffix, "exempt-namespace-suffix", "A namespace with the specified suffix is allowed to set the admission.gatekeeper.sh/ignore label. To exempt multiple suffixes, this flag can be declared more than once.")
}

const ignoreLabel = "admission.gatekeeper.sh/ignore"

// +kubebuilder:webhook:verbs=CREATE;UPDATE,path=/v1/admitlabel,mutating=false,failurePolicy=fail,groups="",resources=namespaces,versions=*,name=check-ignore-label.gatekeeper.sh,sideEffects=None,admissionReviewVersions=v1;v1beta1,matchPolicy=Exact

// AddLabelWebhook registers the label webhook server with the manager.
func AddLabelWebhook(mgr manager.Manager, _ Dependencies) error {
	_ = "STUB: not implemented"
	return nil
}

var _ admission.Handler = &namespaceLabelHandler{}

type namespaceLabelHandler struct{}

//nolint:gocritic // Must accept admission.Request as a struct to satisfy Handler interface.
func (h *namespaceLabelHandler) Handle(_ context.Context, req admission.Request) admission.Response {
	_ = "STUB: not implemented"
	return *new(admission.Response)
}

func matchesPrefix(s string) bool { _ = "STUB: not implemented"; return false }

func matchesSuffix(s string) bool { _ = "STUB: not implemented"; return false }

func GetAllExemptedNamespacesWithWildcard() []string { _ = "STUB: not implemented"; return nil }
