package main

import (
	"context"
	"io"
	"log"
	pb "serverstreamingrpc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50060", grpc.WithInsecure())
	defer conn.Close()
	gc := pb.NewGreeterClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	m_obj := &pb.Request{Msg: "Request"}

	stream, err := gc.GreetManyTimes(context.Background(), m_obj)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		log.Printf("Response: %v", msg.Msg)
	}
}
