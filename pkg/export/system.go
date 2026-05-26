package export

import (
	"context"
	"sync"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/export/dapr"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/export/disk"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/export/driver"
)

var supportedDrivers = map[string]driver.Driver{
	dapr.Name: dapr.Connections,
	disk.Name: disk.Connections,
}

type Exporter interface {
	Publish(ctx context.Context, connectionName string, subject string, msg interface{}) error
	UpsertConnection(ctx context.Context, config interface{}, connectionName string, newDriver string) error
	CloseConnection(connectionName string) error
}

type System struct {
	mux                sync.RWMutex
	connectionToDriver map[string]string
}

func NewSystem() *System { _ = "STUB: not implemented"; return nil }

func (s *System) Publish(ctx context.Context, connectionName string, subject string, msg interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *System) UpsertConnection(ctx context.Context, config interface{}, connectionName string, newDriver string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if the connection already exists.

// If the provider is the same, update the existing connection.

// Check if the provider is supported.

// Close the existing connection after successfully creating the new one.

// Add the new connection and provider to the maps.

func (s *System) CloseConnection(connectionName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *System) closeConnection(connectionName string) error {
	_ = "STUB: not implemented"
	return nil
}

// connection should be deleted from the map before closing it to make sure old connection is not accessible if close fails
// also avoids not respecting the latest connection with the same name if close fails
