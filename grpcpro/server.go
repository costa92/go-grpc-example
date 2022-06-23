package main

import (
	"github.com/costa92/go-grpc-example/grpcpro/services"
	"google.golang.org/grpc"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, &services.ProdService{})
	list, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	rpcServer.Serve(list)
}
