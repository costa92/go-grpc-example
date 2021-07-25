package main

import (
	"context"
	proto "github.com/costa92/go-grpc-example/simple-tcp/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const PORT = ":9001"

func main() {
	pathDir := "/Users/costalong/Documents/Code/Go/src/github.com/costa92/go-grpc-example/tls/config/"
	c, err := credentials.NewClientTLSFromFile(pathDir+"server.pem", "go-grpc-example")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}

	conn, err := grpc.Dial(PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := proto.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &proto.SearchRequest{
		SearchName: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())

}
