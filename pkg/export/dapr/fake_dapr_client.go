package dapr

import (
	"context"
	"log"
	"os"
	"sync"

	daprClient "github.com/dapr/go-sdk/client"
	commonv1pb "github.com/dapr/go-sdk/dapr/proto/common/v1"
	pb "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/export/driver"
)

const (
	testBufSize = 1024 * 1024
	testSocket  = "/tmp/dapr.socket"
	valueSuffix = "_value"
)

var logger = log.New(os.Stdout, "", 0)

func getTestClient(_ context.Context) (client daprClient.Client, closer func()) {
	_ = "STUB: not implemented"
	return *new(daprClient.Client), nil
}

// Replace "" with `serviceAddress` to specify the service to connect to

type testDaprServer struct {
	pb.UnimplementedDaprServer
	state                             map[string][]byte
	configurationSubscriptionIDMapLoc sync.Mutex
	configurationSubscriptionID       map[string]chan struct{}
}

func (s *testDaprServer) TryLockAlpha1(_ context.Context, _ *pb.TryLockRequest) (*pb.TryLockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) UnlockAlpha1(_ context.Context, _ *pb.UnlockRequest) (*pb.UnlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) InvokeService(_ context.Context, req *pb.InvokeServiceRequest) (*commonv1pb.InvokeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) GetState(_ context.Context, req *pb.GetStateRequest) (*pb.GetStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) GetBulkState(_ context.Context, in *pb.GetBulkStateRequest) (*pb.GetBulkStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) SaveState(_ context.Context, req *pb.SaveStateRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) QueryStateAlpha1(_ context.Context, req *pb.QueryStateRequest) (*pb.QueryStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) DeleteState(_ context.Context, req *pb.DeleteStateRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) DeleteBulkState(_ context.Context, req *pb.DeleteBulkStateRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) ExecuteStateTransaction(_ context.Context, in *pb.ExecuteStateTransactionRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) PublishEvent(_ context.Context, _ *pb.PublishEventRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) InvokeBinding(_ context.Context, req *pb.InvokeBindingRequest) (*pb.InvokeBindingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) GetSecret(_ context.Context, _ *pb.GetSecretRequest) (*pb.GetSecretResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) GetBulkSecret(_ context.Context, _ *pb.GetBulkSecretRequest) (*pb.GetBulkSecretResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) RegisterActorReminder(_ context.Context, _ *pb.RegisterActorReminderRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) UnregisterActorReminder(_ context.Context, _ *pb.UnregisterActorReminderRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) RenameActorReminder(_ context.Context, _ *pb.RenameActorReminderRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) InvokeActor(context.Context, *pb.InvokeActorRequest) (*pb.InvokeActorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) RegisterActorTimer(context.Context, *pb.RegisterActorTimerRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) UnregisterActorTimer(context.Context, *pb.UnregisterActorTimerRequest) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) Shutdown(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) GetConfiguration(_ context.Context, in *pb.GetConfigurationRequest) (*pb.GetConfigurationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testDaprServer) SubscribeConfiguration(in *pb.SubscribeConfigurationRequest, server pb.Dapr_SubscribeConfigurationServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Send subscription ID in the first response.

func (s *testDaprServer) UnsubscribeConfiguration(_ context.Context, in *pb.UnsubscribeConfigurationRequest) (*pb.UnsubscribeConfigurationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BulkPublishEventAlpha1 mocks the BulkPublishEventAlpha1 API.
// It will fail to publish events that start with "fail".
// It will fail the entire request if an event starts with "failall".
func (s *testDaprServer) BulkPublishEventAlpha1(_ context.Context, req *pb.BulkPublishRequest) (*pb.BulkPublishResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fail the entire request

// fail this entry

func FakeConnection() (driver.Driver, func()) {
	_ = "STUB: not implemented"
	return *new(driver.Driver), nil
}

type FakeDaprConnection struct {
	component string

	client daprClient.Client
	// closing function
	f func()
}

type FakeDapr struct {
	openConnections map[string]FakeDaprConnection
}

func (r *FakeDapr) Publish(_ context.Context, _ string, _ interface{}, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *FakeDapr) CloseConnection(connectionName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *FakeDapr) UpdateConnection(_ context.Context, connectionName string, config interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *FakeDapr) CreateConnection(ctx context.Context, connectionName string, config interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

var FakeConn = &FakeDapr{
	openConnections: map[string]FakeDaprConnection{},
}
