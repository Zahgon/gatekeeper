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

package main

import (
	"context"
	"flag"
	"io"
	_ "net/http/pprof"
	"os"
	"time"

	api "github.com/open-policy-agent/gatekeeper/v3/apis"
	configv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/config/v1alpha1"
	connectionv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/connection/v1alpha1"
	expansionv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/expansion/v1alpha1"
	expansionv1beta1 "github.com/open-policy-agent/gatekeeper/v3/apis/expansion/v1beta1"
	mutationsv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/mutations/v1alpha1"
	mutationsv1beta1 "github.com/open-policy-agent/gatekeeper/v3/apis/mutations/v1beta1"
	statusv1beta1 "github.com/open-policy-agent/gatekeeper/v3/apis/status/v1beta1"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/util"
	"go.uber.org/zap/zapcore"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	secretName     = "gatekeeper-webhook-server-cert"
	caName         = "gatekeeper-ca"
	caOrganization = "gatekeeper"
	certName       = "tls.crt"
	keyName        = "tls.key"
)

var (
	scheme           = runtime.NewScheme()
	setupLog         = ctrl.Log.WithName("setup")
	logLevelEncoders = map[string]zapcore.LevelEncoder{
		"lower":        zapcore.LowercaseLevelEncoder,
		"capital":      zapcore.CapitalLevelEncoder,
		"color":        zapcore.LowercaseColorLevelEncoder,
		"capitalcolor": zapcore.CapitalColorLevelEncoder,
	}
)

var (
	logFile                              = flag.String("log-file", "", "Log to file, if specified. Default is to log to stderr.")
	logLevel                             = flag.String("log-level", "INFO", "Minimum log level. For example, DEBUG, INFO, WARNING, ERROR. Defaulted to INFO if unspecified.")
	logLevelKey                          = flag.String("log-level-key", "level", "JSON key for the log level field, defaults to `level`")
	logLevelEncoder                      = flag.String("log-level-encoder", "lower", "Encoder for the value of the log level field. Valid values: [`lower`, `capital`, `color`, `capitalcolor`], default: `lower`")
	healthAddr                           = flag.String("health-addr", ":9090", "The address to which the health endpoint binds.")
	metricsAddr                          = flag.String("metrics-addr", "0", "The address the metric endpoint binds to.")
	port                                 = flag.Int("port", 443, "port for the server. defaulted to 443 if unspecified ")
	host                                 = flag.String("host", "", "the host address the webhook server listens on. defaults to all addresses.")
	certDir                              = flag.String("cert-dir", "/certs", "The directory where certs are stored, defaults to /certs")
	disableCertRotation                  = flag.Bool("disable-cert-rotation", false, "disable automatic generation and rotation of webhook TLS certificates/keys")
	enableProfile                        = flag.Bool("enable-pprof", false, "enable pprof profiling")
	profilePort                          = flag.Int("pprof-port", 6060, "port for pprof profiling. defaulted to 6060 if unspecified")
	certServiceName                      = flag.String("cert-service-name", "gatekeeper-webhook-service", "The service name used to generate the TLS cert's hostname. Defaults to gatekeeper-webhook-service")
	enableTLSHealthcheck                 = flag.Bool("enable-tls-healthcheck", false, "enable probing webhook API with certificate stored in certDir")
	disabledBuiltins                     = util.NewFlagSet()
	enableK8sCel                         = flag.Bool("enable-k8s-native-validation", true, "enable the validating admission policy driver")
	externaldataProviderResponseCacheTTL = flag.Duration("external-data-provider-response-cache-ttl", 3*time.Minute, "TTL for the external data provider response cache. Specify the duration in 'h', 'm', or 's' for hours, minutes, or seconds respectively. Defaults to 3 minutes if unspecified. Setting the TTL to 0 disables the cache.")
	enableReferential                    = flag.Bool("enable-referential-rules", true, "Enable referential rules. This flag defaults to true. Set this value to false if you want to disallow referential constraints. Because referential constraints read objects other than the object-under-test, they may be subject to race conditions. Users concerned about this may want to disable referential rules")
	shutdownDelay                        = flag.Int("shutdown-delay", 10, "Time in seconds the controller runtime shutdown gets delayed after receiving a pod termination event. Prevents failing webhooks on pod shutdown. default: 10")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	_ = api.AddToScheme(scheme)

	_ = configv1alpha1.AddToScheme(scheme)
	_ = statusv1beta1.AddToScheme(scheme)
	_ = mutationsv1alpha1.AddToScheme(scheme)
	_ = mutationsv1beta1.AddToScheme(scheme)
	_ = expansionv1alpha1.AddToScheme(scheme)
	_ = expansionv1beta1.AddToScheme(scheme)
	_ = connectionv1alpha1.AddToScheme(scheme)

	// +kubebuilder:scaffold:scheme
	flag.Var(disabledBuiltins, "disable-opa-builtin", "disable opa built-in function, this flag can be declared more than once.")
}

func main() {
	os.Exit(innerMain())
}

func innerMain() int { _ = "STUB: not implemented"; return 0 }

// Disable high-cardinality REST client metrics (rest_client_request_latency).
// Must be called before ctrl.NewManager!

// Make sure certs are generated and valid if cert rotation is enabled.

// Setup tracker and register readiness probe.

// +kubebuilder:scaffold:builder

// only setup healthcheck when flag is set

// Setup controllers asynchronously, they will block for certificate generation if needed.

// Setup termination with grace period. Required to give K8s Services time to disconnect the Pod endpoint on termination.
// Derived from how the controller-runtime sets up a signal handler with ctrl.SetupSignalHandler()
// controller-runtime upstream issue: https://github.com/kubernetes-sigs/controller-runtime/issues/3113

// second signal. Exit directly.

// block until either setupControllers or mgr has an error, or mgr exits.
// end after two events (one per goroutine) to guard against deadlock.

// if manager has returned, we should exit the program

func setupControllers(ctx context.Context, mgr ctrl.Manager, tracker *readiness.Tracker, setupFinished chan struct{}) error {
	_ = "STUB: not implemented"
	// Block until the setup (certificate generation) finishes.
	return nil
}

// certWatcher is used to watch for changes to Gatekeeper's certificate and key files.

// register the client cert watcher to the driver

// register the client cert watcher to the mutation system

// processExcluder is used for namespace exclusion for specified processes in config

// Setup all Controllers

// Events ch will be used to receive events from dynamic watches registered
// via the registrar below.

func setLoggerForProduction(encoder zapcore.LevelEncoder, dest io.Writer) {
	_ = "STUB: not implemented"
	return
}
