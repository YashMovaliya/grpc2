package main

import (
	"context"

	"github.com/YashMovaliya/grpc2"
	grpc2grpc "github.com/YashMovaliya/grpc2/grpc"
)

// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
	userService grpc2.Service
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService grpc2.Service) grpc2grpc.UserServiceServer {
	return &userServiceController{
		userService: userService,
	}
}

// GetUsers calls the core service's GetUsers method and maps the result to a grpc service response.
func (ctlr *userServiceController) GetUsers(ctx context.Context, req *grpc2grpc.GetUsersRequest) (resp *grpc2grpc.GetUsersResponse, err error) {
	resultMap, err := ctlr.userService.GetUsers(req.GetIds())
	if err != nil {
		return
	}
	resp := &grpc2grpc.GetUsersResponse{}
	for _, u := range resultMap {
		resp.Users = append(resp.Users, marshalUser(&u))
	}
	return
}

// marshalUser marshals a business object User into a gRPC layer User.
func marshalUser(u *grpc2.User) *grpc2grpc.User {
	return &grpc2grpc.User{Id: u.ID, Name: u.Name}
}
