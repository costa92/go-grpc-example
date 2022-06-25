package main

import (
	"context"
	helloword "github.com/costa92/go-grpc-example/helloword/services/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloServiceImpl struct {
}

func (p *HelloServiceImpl) mustEmbedUnimplementedHelloServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *helloword.String,
) (*helloword.String, error) {
	reply := &helloword.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	// 生成服务
	helloword.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
