package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sycki/examples/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 实现grpc/proto包中的GreeterServer接口
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) AddUser(ctx context.Context, user *pb.UserRequest) (*pb.UserResponse, error) {
	log.Println("add user:", user.Name)
	return &pb.UserResponse{
		Name: user.Name + " a year older",
		Age:  user.Age + 1,
	}, nil
}

func (s *server) GetUser(ctx context.Context, user *pb.UserRequest) (*pb.UserResponse, error) {
	log.Println("get user:", user.Name)
	return &pb.UserResponse{
		Name: user.Name + " be get",
	}, nil
}

func main() {
	// 监听一个端口
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个RPC服务端并启动
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: ", err.Error())
	}
}
