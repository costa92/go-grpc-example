package main

import (
	"context"
	proto "github.com/costa92/go-grpc-example/simple-tcp/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type SearchServices struct{
}

func( s *SearchServices) Search(ctx context.Context,r *proto.SearchRequest) (* proto.SearchResponse ,error){
	return &proto.SearchResponse{Response: r.GetSearchName() + "Server" },nil
}
const PORT = ":9001"
func main() {
	pathDir := "/Users/costalong/Documents/Code/Go/src/github.com/costa92/go-grpc-example/tls/config/"
	c, err := credentials.NewServerTLSFromFile(pathDir+"server.pem", pathDir+"server.key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	server := grpc.NewServer(grpc.Creds(c))
	proto.RegisterSearchServiceServer(server,&SearchServices{})

	listener , err := net.Listen("tcp",PORT)

	if err != nil{
		log.Fatalf("net.Listener err:%v",err)
	}

	server.Serve(listener)
}
