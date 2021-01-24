package main

import (
	pb "github.com/salmansiddiquimailbox/gRPC-go.git/server-streaming/exampleProto"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetManyTimes(in *pb.Request, stream pb.Greeter_GreetManyTimesServer) error {
	message := in.Msg
	log.Printf("Recieved: %v\n", message)

	for i := 0; i < 10; i++ {
		result := "Hello" + strconv.Itoa(i)
		res := &pb.Response{
			Msg: result,
		}
		stream.Send(res)
		time.Sleep(time.Second)
	}
	return nil
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
