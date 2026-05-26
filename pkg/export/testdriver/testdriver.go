package testdriver

import (
	"context"
)

const (
	Name    = "testdriver"
	ErrName = "testdriver-error"
)

var FakeConn = &Connection{
	openConnections: make(map[string]FakeConnection),
}

// Connection represents driver to use testdriver.
type Connection struct {
	openConnections map[string]FakeConnection
}

type FakeConnection struct {
	name string
}

func (r *Connection) Publish(_ context.Context, _ string, _ interface{}, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Connection) CloseConnection(connectionName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Connection) UpdateConnection(_ context.Context, connectionName string, config interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Connection) CreateConnection(_ context.Context, connectionName string, config interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

var FakeErrConn = &ErrConnection{
	openErrConnections: make(map[string]FakeErrConnection),
}

// ErrConnection represents driver to use testdriver.
type ErrConnection struct {
	openErrConnections map[string]FakeErrConnection
}

type FakeErrConnection struct {
	name string
}

func (r *ErrConnection) Publish(_ context.Context, _ string, _ interface{}, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ErrConnection) CloseConnection(connectionName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ErrConnection) UpdateConnection(_ context.Context, connectionName string, config interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ErrConnection) CreateConnection(_ context.Context, connectionName string, config interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
