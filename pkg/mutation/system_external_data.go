package mutation

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/open-policy-agent/frameworks/constraint/pkg/externaldata"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/sets"
)

// resolvePlaceholders resolves all external data placeholders in the given object.
func (s *System) resolvePlaceholders(ctx context.Context, obj *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

// recurse object to find all existing external data placeholders

// gather and de-duplicate all keys for this
// provider so we can resolve them in batch

const defaultExternalDataRequestTimeout = 5 * time.Second

// sendRequests sends requests to all providers in parallel.
func (s *System) sendRequests(ctx context.Context, providerKeys map[string]sets.Set[string], clientCert *tls.Certificate) (map[string]map[string]*externaldata.Item, map[string]error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the provider name is the first key and the outbound data is the second key

// errors that might have occurred per provider

// mutateWithExternalData recursively traverses the given object and replaces
// all external data placeholders with the corresponding external data items.
func (s *System) mutateWithExternalData(object *unstructured.Unstructured, externalData map[string]map[string]*externaldata.Item, errors map[string]error) error {
	_ = "STUB: not implemented"
	return nil
}

// not a placeholder, let's continue recursing

// our base case - we found a placeholder and we should resolve it

// we expect the response to contain the key we're looking for

// getTLSCertificate returns the gatekeeper's TLS certificate.
func (s *System) getTLSCertificate() (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateExternalDataResponse validates the given external data response.
func validateExternalDataResponse(r *externaldata.ProviderResponse) error {
	_ = "STUB: not implemented"
	return nil
}
