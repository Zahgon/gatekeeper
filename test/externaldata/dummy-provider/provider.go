package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/open-policy-agent/frameworks/constraint/pkg/externaldata"
)

const (
	timeout    = 1 * time.Second
	apiVersion = "externaldata.gatekeeper.sh/v1alpha1"
)

func main() {
	fmt.Println("starting server...")

	// load Gatekeeper's CA certificate
	caCert, err := os.ReadFile("/tmp/gatekeeper/ca.crt")
	if err != nil {
		panic(err)
	}

	clientCAs := x509.NewCertPool()
	clientCAs.AppendCertsFromPEM(caCert)

	mux := http.NewServeMux()
	mux.HandleFunc("/validate", processTimeout(validate, timeout))

	server := &http.Server{
		Addr:              ":8090",
		Handler:           mux,
		ReadHeaderTimeout: timeout,
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  clientCAs,
			MinVersion: tls.VersionTLS13,
		},
	}

	if err := server.ListenAndServeTLS("/etc/ssl/certs/server.crt", "/etc/ssl/certs/server.key"); err != nil {
		panic(err)
	}
}

func validate(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	// only accept POST requests
	return
}

// read request body

// parse request body

// iterate over all keys

// Providers should add a caching mechanism to avoid extra calls to external data sources.

// following checks are for testing purposes only
// check if key contains "_systemError" to trigger a system error

// check if key contains "error_" to trigger an error

// valid key will have "_valid" appended as return value

// sendResponse sends back the response to Gatekeeper.
func sendResponse(results *[]externaldata.Item, systemErr string, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func processTimeout(h http.HandlerFunc, duration time.Duration) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
