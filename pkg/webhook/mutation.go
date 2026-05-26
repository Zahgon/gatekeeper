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

package webhook

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/open-policy-agent/cert-controller/pkg/rotator"
	"github.com/open-policy-agent/gatekeeper/v3/apis"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func init() {
	AddToManagerFuncs = append(AddToManagerFuncs, AddMutatingWebhook)

	if err := apis.AddToScheme(runtimeScheme); err != nil {
		log.Error(err, "unable to add to scheme")
		panic(err)
	}
}

// +kubebuilder:webhook:verbs=create;update,path=/v1/mutate,mutating=true,failurePolicy=ignore,groups=*,resources=*,versions=*,name=mutation.gatekeeper.sh,sideEffects=None,admissionReviewVersions=v1;v1beta1,matchPolicy=Exact
// +kubebuilder:rbac:resourceNames=gatekeeper-mutating-webhook-configuration,groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;update;patch

// AddMutatingWebhook registers the mutating webhook server with the manager.
func AddMutatingWebhook(mgr manager.Manager, deps Dependencies) error {
	_ = "STUB: not implemented"
	return nil
}

var _ admission.Handler = &mutationHandler{}

type mutationHandler struct {
	webhookHandler
	mutationSystem *mutation.System
	deserializer   runtime.Decoder
	log            logr.Logger
}

// Handle the mutation request
// nolint: gocritic // Must accept admission.Request to satisfy interface.
func (h *mutationHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	_ = "STUB: not implemented"
	return *new(admission.Response)
}

// namespace is excluded from webhook using config

func (h *mutationHandler) mutateRequest(ctx context.Context, req *admission.Request) admission.Response {
	_ = "STUB: not implemented"
	return *

	// if the object being mutated is a namespace itself, we use it as namespace
	new(admission.Response)
}

// bypass cached client and ask api-server directly

// It is possible for the namespace to not be populated on an object.
// Assign the namespace from the request object (which will have the appropriate
// value), then restore the original value at the end to avoid sending a namespace patch.

func AppendMutationWebhookIfEnabled(webhooks []rotator.WebhookInfo) []rotator.WebhookInfo {
	_ = "STUB: not implemented"
	return nil
}
