package main

import (
	"context"
	pb "github.com/peifengll/tests/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("www.baidu.com")
	return &pb.HelloReply{Message: "66666 " + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalln("failed to listen", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("server listening at:", listen.Addr())

	if err := s.Serve(listen); err != nil {
		log.Fatalln("failed to  server ", err)
	}
}
