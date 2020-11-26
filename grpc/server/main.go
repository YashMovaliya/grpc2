package main

import (
	"net"
	"os"

	grpc2core "github.com/YashMovaliya/grpc2/core"
	grpc2grpc "github.com/YashMovaliya/grpc2/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// configure our core service
	userService := grpc2core.NewService()
	// configure our gRPC service controller
	userServiceController := NewUserServiceController(userService)
	// start a gRPC server
	server := grpc.NewServer()
	grpc2grpc.RegisterUserServiceServer(server, userServiceController)
	reflection.Register(server)
	con, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		panic(err)
	}
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
