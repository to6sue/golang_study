package client

/*
import (
	pb "app/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

package main

import (
pb "proto"
"context"
"google.golang.org/grpc"
"log"
"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewConfigStoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &pb.ConfigRequest{Profile: "dev"})
	if err != nil {
		log.Fatalf("could not request: %v", err)
	}

	r2, err := c.Get(ctx, &pb.ConfigRequest{Profile: "prod"})
	if err != nil {
		log.Fatalf("could not request: %v", err)
	}

	log.Printf("Config: %v", r)
	log.Printf("Config: %v", r2)
}
*/
