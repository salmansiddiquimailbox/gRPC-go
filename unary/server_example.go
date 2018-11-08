package main

import (
	pb "grpcexample"
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Println("Info: Greet RPC called")
	value := in.Msg
	log.Println("Recieved: ", value)
	res := &pb.Response{Msg: "Hello Back!"}
	return res, nil
}

func main() {
	log.Println("Starting gRPC Server-->")
	lis, err := net.Listen("tcp", ":50060")
	if err != nil {
		log.Fatalf("Error while listening: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	er := s.Serve(lis)
	if er != nil {
		log.Fatalf("Error while serving: %v", er)
	}
}
