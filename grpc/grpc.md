# GRPC

## 官方文档

http://doc.oschina.net/grpc

## 快速开始

### 安装

####  protobuf

[下载地址](https://github.com/protocolbuffers/protobuf/releases)

> 解压下载的压缩包，并将其移动到 `$GOBIN` 目录，查看 `$GOBIN` 目录

**验证**

```shell
protoc --version
```

#### protoc-gen-go-grpc

```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 创建项目

```
mkdir -p $GOPATH/src/protobuf-go && cd $GOPATH/src/protobuf-go && go mod init
mkdir pb
```

### protol文件

> server.proto

```protobuf
syntax = "proto3"; //proto3 协议
option go_package = "./pb";

message Empty{}

message HelloResponse{
	string hello = 1;
}

message RegisterRequest{
	string name = 1;
	string password = 2;
}

message RegisterResponse{
	string uid = 1;
}

service Server{
	rpc Hello(Empty) returns(HelloResponse);
	rpc Register(RegisterRequest) returns(RegisterResponse);
}

```

### 命令生成

> 生成的文件路径有问题 需要手动移动一下

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --go-grpc_opt=require_unimplemented_servers=false \
    pb/server.proto 
```

### main文件

```go
package main

import (
	"fmt"

	"github.com/xiongwei/protobuf-go/pb"
)

func main() {
	NewRpcMessage()
}

func NewRpcMessage() {
	msg := pb.RegisterRequest{
		Name: "xw",
	}
	fmt.Println(msg.GetName())
}

```

### gomod

```
go mod tidy
```

### 运行

```
go run main.go
```

## 脚本

> 需手动创建对应的文件夹

```shell
echo "生成 rpc server 代码"

OUT=../server/pb

mkdir -p ${OUT}
protoc \
--go_out=${OUT} \
--go-grpc_out=${OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
*.proto

mv ${OUT}/pb/*  ${OUT}/
rm -rf ${OUT}/pb


echo "生成 rpc client 代码"

OUT=../client/pb

mkdir -p ${OUT}
protoc \
--go_out=${OUT} \
--go-grpc_out=${OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
*.proto

mv ${OUT}/pb/*  ${OUT}/
rm -rf ${OUT}/pb

```



## grpc服务端

server.go

```go
package main

import (
	"context"
	"fmt"
	"study/protobuf-go/server/pb"
)

type Server struct {
}

func (s Server) Hello(ctx context.Context, request *pb.Empty) (*pb.HelloResponse, error) {
	resp := pb.HelloResponse{Hello: "hello client."}
	return &resp, nil
}

func (s Server) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	resp := pb.RegisterResponse{}
	resp.Uid = fmt.Sprintf("%s.%s", request.GetName(), request.GetPassword())
	return &resp, nil
}

```



main.go

```go
package main

import (
	"fmt"
	"log"
	"net"

	"study/protobuf-go/server/pb"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))

	if err != nil {
		log.Fatalf("启动grpc server失败")
		return
	}

	grpcServer := grpc.NewServer()

	pb.RegisterServerServer(grpcServer, Server{})

	log.Println("service start")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动grpc server失败")
	}
}
```

## grpc客户端

main.go

```go
package main

import (
	"context"
	"fmt"
	"log"
	"study/protobuf-go/client/pb"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	ServerClient := pb.NewServerClient(conn)

	helloRespone, err := ServerClient.Hello(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(helloRespone, err)

	registerResponse, err := ServerClient.Register(context.Background(), &pb.RegisterRequest{Name: "chihuo", Password: "123456"})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(registerResponse, err)

}
```







