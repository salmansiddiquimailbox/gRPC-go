package main

import (
	pb "clientstreaming"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

//stream pb.ServiceName_RPCNameServer
func (*server) Greet(stream pb.Greeter_GreetServer) error {
	log.Println("Greet Called")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF recieved!")
			return stream.SendAndClose(&pb.Response{
				Msg: result,
			})
		}

		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		result += "Hello" + " "
		log.Println("Received: ", req.Msg)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50060")
	if err != nil {
		log.Fatalf("%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("Info: Starting gRPC server")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
