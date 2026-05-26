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

package readiness

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// SetupTracker sets up a readiness tracker and registers it to run under control of the
// provided Manager object.
// NOTE: Must be called _before_ the manager is started.
func SetupTracker(mgr manager.Manager, mutationEnabled, externalDataEnabled, expansionEnabled bool) (*Tracker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetupTrackerNoReadyz sets up a readiness tracker and registers it to run under control of the
// provided Manager object without instantiating /readyz (used for testing).
func SetupTrackerNoReadyz(mgr manager.Manager, mutationEnabled, externalDataEnabled, expansionEnabled bool) (*Tracker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
