package main

import (
	proto "github.com/costa92/go-grpc-example/stream/server"
	"google.golang.org/grpc"
	"log"
	"net"
)


const PORT = ":9001"

func main() {
	server := grpc.NewServer()
	proto.RegisterStreamServiceServer(server,&proto.StreamService{})
	listener , err := net.Listen("tcp",PORT)
	if err != nil {
		log.Fatalf("net.Listene err:%v",err)
	}
	log.Println("Start listen:"+ PORT)
	server.Serve(listener)

}

