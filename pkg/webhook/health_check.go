package webhook

import (
	"context"
	"net/http"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var tlsCheckerLog = logf.Log.WithName("webhook-tls-checker")

func NewTLSChecker(certDir, host string, port int) func(*http.Request) error {
	_ = "STUB: not implemented"
	//nolint:forcetypeassert
	return nil
}

// disabling gosec linting here as the http client used in this checking is intended to skip CA verification
//
//nolint:gosec

// disable keep alives to ensure that http connection aren't reused, otherwise the check may
// fail if the cert was rotated in between

func probeTLSURL(ctx context.Context, insecureClient *http.Client, probeURL string, expectedCerts [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// explicitly discard the body to avoid any memory leak

// compare certificate in resp and the certificate in certDir

func tlsProbeHosts(host string) []string { _ = "STUB: not implemented"; return nil }
