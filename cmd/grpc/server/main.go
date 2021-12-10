package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"jim/global"
	"jim/internal/rpc/proto_build"
	"jim/internal/rpc/service"
	"net"
)

func newGrpcServer() *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)
	proto_build.RegisterUserServiceServer(server, service.NewUserService())
	proto_build.RegisterPingServiceServer(server, service.NewPingService())
	return server
}

func main() {
	defer func() {
		global.Close()
	}()
	rpc := global.Config.Rpc
	server := newGrpcServer()
	address := fmt.Sprintf("%s:%d", rpc.Host, rpc.Port)
	grpclog.Infoln("listen on :" + address)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		grpclog.Fatalln("failed to listen：", err)
	}
	_ = server.Serve(listen)
}
