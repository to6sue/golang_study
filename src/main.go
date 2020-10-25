package main

import (
	pb "app/proto"
	"log"

	"context"
	"google.golang.org/grpc"
	"net"
	"fmt"
)

/*// server는 protobuf에서 정의된 함수의 인자로서 사용된다.
type server struct{}


// ProtoBuf의 IDL에 정의되어 있는 함수
// 함수의 인자와 리턴 값인 HelloRequest, HelloReply, 그리고 아래의 함수들은 모두
// protoc에서 생성된 skeleton 코드를 그대로 사용한다.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, Again!" + in.Name}, nil
}*/

type server struct {
	pb.UnimplementedConfigStoreServer
}

func (s *server) Get(ctx context.Context, in *pb.ConfigRequest) (*pb.ConfigResponse, error) {
	log.Printf("Received profile: %v", in.GetProfile())
	if in.GetProfile() == "dev" {
		return &pb.ConfigResponse{JsonConfig: `"{"main":"http://dev.com"}"`}, nil
	} else {
		return &pb.ConfigResponse{JsonConfig: `"{"main":"http://prod.com"}"`}, nil
	}

}

const (
	port = ":50051"
)

func main() {
	/*
	fmt.Println("Get Database Config")
	config := config.GetConfig()

	fmt.Println("Run Server (port 8081)")
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8081")
*/


	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("GRPC Server Listening")
	s := grpc.NewServer()
	pb.RegisterConfigStoreServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		fmt.Println("GRPC Server Fail")
	} else {
		fmt.Println("GRPC Server Listening")
	}

}