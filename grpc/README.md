# gRPC使用指南
gRPC是一个由google开源的RPC框架，它支持多种语言之间的相互调用，依赖开源项目`google/protobuf`作为序列化工具。

相比于几年前的gRPC，新版gRPC的使用方式发生了一些改变，编译`.proto`文件时所用的命令已经与之前不同，而且产生的`.go`文件按照数据结构序列化和使用接口分为两个文件。
## 准备环境

### 安装Golang
安装[最新版Golang](https://golang.org/dl/)。

### 安装Protocol Buffers编译器
`protoc`是一个二进制文件，用来编译`.proto`文件，输出指定语言的源码。进入[`protoc`发布页面](https://github.com/google/protobuf/releases)，按照自己的平台下载相应压缩包，比如64位linux就下载[`protoc-3.13.0-linux-x86_64.zip`](https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip)，解压后把二进制文件放到`$GOPATH/bin`下。
```
$ protoc --version
libprotoc 3.13.0
```

### 安装Protocol Buffers GO插件
`protoc-gen-go`和`protoc-gen-go-grpc`两个二进制文件会被`protoc`调用，前者用于从`.proto`文件中编译出Go语言数据结构，这些数据结构带有序列化函数，后者用于产生Go语言编程接口。
```
go get google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### 添加Golang依赖
执行前确保项目中包含了`go.mod`文件。
```
go get google.golang.org/grpc
```

## 编写一个Demo
新建一个名为grpc的Golang项目，结构如下：
```
grpc
├── client
│   └── main.go
├── proto
│   ├── grpc.pb.go
│   └── grpc.proto
└── server
    └── main.go
```

### 编写.proto文件
`grpc/proto/grpc.proto`
```
ssyntax = "proto3";

package proto;

option go_package = "github.com/sycki/examples/grpc";

// 定义一个RPC Server
// 包含两个函数可供Client调用
service Greeter {
  rpc AddUser (UserRequest) returns (UserResponse);
  rpc GetUser (UserRequest) returns (UserResponse);
}
// 调用时的数据格式
message UserRequest {
	string Name = 1;
	int32 age = 2;
	string address = 3;
}
// 返回值的数据格式
message UserResponse {
	string Name = 1;
	int32 age = 2;
	string address = 3;
}
```

### 生成Golang代码
```
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
proto/grpc.proto
```
执行成功会在proto目录下生成两个`.go`源文件，一个包含了数据结构，一个包含调用接口。

### 编写RPC服务端
`grpc/server/main.go`
```
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
```

### 编写RPC客户端
`grpc/client/main.go`
```
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
```

### 启动服务端
```
go run server/main.go
```

### 启动客户端
```
go run clinet/main.go

2017/12/31 17:46:02 success AddUser => Name:jack a year older, Age:19
```

关于
---

__作者__：张佳军

__阅读__：100

__点赞__：3

__创建__：2017-12-31
