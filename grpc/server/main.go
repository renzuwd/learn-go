package main

import (
	"fmt"
	"log"
	"net"

	"learn-go/grpc/server/pb"

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
