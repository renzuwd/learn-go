package main

import (
	"context"
	"fmt"
	"learn-go/grpc/client/pb"
	"log"

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
