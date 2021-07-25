package main

import (
	"context"
	server "github.com/costa92/go-grpc-example/simple-tcp/server"
	"google.golang.org/grpc"
	"log"
	"net"
)


type SearchServices struct{
}

func( s *SearchServices) Search(ctx context.Context,r *server.SearchRequest) (* server.SearchResponse ,error){
	return &server.SearchResponse{Response: r.GetSearchName() + "Server" },nil
}

const (
	PORT = ":8081"
)
func main() {
	servers := grpc.NewServer()
	server.RegisterSearchServiceServer(servers,&SearchServices{})
	listener,err := net.Listen("tcp",PORT)
	if err != nil {
		log.Fatalf("net.Listern err:%v",err)
	}
	servers.Serve(listener)
}
