package main

import (
	"context"
	"log"
	"time"

	"github.com/costa92/go-grpc-example/grpcpro/services"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}

	defer conn.Close()
	client := services.NewProdServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetProdStock(ctx, &services.ProdRequest{
		ProdId: 1,
	})

	if err != nil {
		log.Fatalf("could not greet:%v", err)
	}
	log.Printf("Greeting: %d", r.ProdStock)
}
