package main

import (
	"context"
	helloword "github.com/costa92/go-grpc-example/helloword/services/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:1234",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()

	c := helloword.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hello(ctx, &helloword.String{Value: "test"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", r.String())
}
