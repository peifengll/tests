package main

import (
	"context"
	pb "github.com/peifengll/tests/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	name := "world"
	addr := "127.0.0.1:50051"

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("connect error to ", addr)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalln("can not greet to ", addr, err)
	} else {
		log.Println("response from server: ", r.GetMessage())
	}
}
