package client

import (
	"context"
	"time"

	"github.com/YashMovaliya/grpc2"
	grpc2grpc "github.com/YashMovaliya/grpc2/grpc"
	"google.golang.org/grpc"
)

var defaultRequestTimeout = time.Second * 10

type grpcService struct {
	grpcClient grpc2grpc.UserServiceClient
}

// NewGRPCService creates a new gRPC user service connection using the specified connection string.
func NewGRPCService(connString string) (grpc2.Service, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{grpcClient: grpc2grpc.NewUserServiceClient(conn)}, nil
}
func (s *grpcService) GetUsers(ids []int64) (result map[int64]grpc2.User, err error) {
	result = map[int64]grpc2.User{}
	req := &grpc2grpc.GetUsersRequest{
		Ids: ids,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUsers(ctx, req)
	if err != nil {
		return
	}
	for _, grpcUser := range resp.GetUsers() {
		u := unmarshalUser(grpcUser)
		result[u.ID] = u
	}
	return
}
func (s *grpcService) GetUser(id int64) (result grpc2.User, err error) {
	req := &grpc2grpc.GetUsersRequest{
		Ids: []int64{id},
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUsers(ctx, req)
	if err != nil {
		return
	}
	for _, grpcUser := range resp.GetUsers() {
		// sanity check: only the requested user should be present in results
		if grpcUser.GetId() == id {
			return unmarshalUser(grpcUser), nil
		}
	}
	return result, grpc2.ErrNotFound
}
func unmarshalUser(grpcUser *grpc2grpc.User) (result grpc2.User) {
	result.ID = grpcUser.Id
	result.Name = grpcUser.Name
	return
}
