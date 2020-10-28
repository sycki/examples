package main

import (
	"context"
	"log"

	pb "github.com/sycki/examples/grpc/proto"

	"google.golang.org/grpc"
)

var address = "localhost:8081"

func main() {
	// 连接到端口
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("can not connect to server:", address)
	}
	defer conn.Close()

	// 创建RPC客户端
	client := pb.NewGreeterClient(conn)

	// 调用服务器中的AddUser方法，并得到返回值
	result, err := client.AddUser(context.Background(), &pb.UserRequest{
		Name: "jack",
		Age:  18,
	})
	if err != nil {
		log.Println("failed AddUser")
		return
	}

	log.Printf("success AddUser => Name:%v, Age:%v\n", result.Name, result.Age)
}
