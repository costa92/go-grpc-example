package main

import (
	"context"
	"fmt"
	proto "github.com/costa92/go-grpc-example/simple-tcp/server"
	"google.golang.org/grpc"
	"log"
)

const (
	PORT = ":8081"
)

func main() {
	conn,err := grpc.Dial(PORT,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err:%v",err)
	}
	defer conn.Close()

	client := proto.NewSearchServiceClient(conn)
	resp ,err := client.Search(context.Background(),&proto.SearchRequest{
		SearchName: "get test ",
	})

	if err != nil {
		log.Fatalf("client search err:%v",err)
	}

	fmt.Println(resp.GetResponse())
}
