package dapr

import (
	"context"

	daprClient "github.com/dapr/go-sdk/client"
)

type Connection struct {
	// Name of the component object to use in Dapr
	component string

	client daprClient.Client
}

// Dapr represents driver to use Dapr.
type Dapr struct {
	openConnections map[string]Connection
}

const (
	Name = "dapr"
)

var Connections = &Dapr{
	openConnections: make(map[string]Connection),
}

func (r *Dapr) Publish(_ context.Context, connectionName string, data interface{}, topic string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Dapr) CloseConnection(connectionName string) error { _ = "STUB: not implemented"; return nil }

func (r *Dapr) UpdateConnection(_ context.Context, connectionName string, config interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Dapr) CreateConnection(_ context.Context, connectionName string, config interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
