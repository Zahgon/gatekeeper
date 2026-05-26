package common

import (
	"sync"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

var (
	opts            []metric.Option
	res             *resource.Resource
	mutex           sync.Mutex
	requiredReaders int
)

// SetRequiredReaders sets the number of required readers for the MeterProvider.
func SetRequiredReaders(num int) { _ = "STUB: not implemented"; return }

// AddReader adds a reader to the options and updates the MeterProvider if the required conditions are met.
func AddReader(opt metric.Option) { _ = "STUB: not implemented"; return }

// SetResource sets the resource to be used by the MeterProvider.
func SetResource(r *resource.Resource) { _ = "STUB: not implemented"; return }

// setMeterProvider sets the MeterProvider if the required conditions are met.
// The required conditions are:
// 1. The number of readers initiated is equal to the number of metrics backends that are not in error state.
// 2. There is at least one reader.
func setMeterProvider() {
	_ = "STUB: not implemented"
	// Check if we have the required number of readers and at least one reader.
	return
}

// Start with the existing options.

// Add views to the options.

// If a resource is available, add it to the options.
