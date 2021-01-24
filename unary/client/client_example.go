package main

import (
	pb "github.com/salmansiddiquimailbox/gRPC-go.git/unary/exampleProto"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	log.Println("starting gRPC client ")
	conn, err := grpc.Dial("localhost:50060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial: \n%v", err)
	}
	defer conn.Close()

	gc := pb.NewGreeterClient(conn)
	//Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	m_obj := &pb.Request{Msg: "This is a request"}
	rec, er := gc.Greet(ctx, m_obj)
	if er != nil {
		log.Fatalf("Could'nt receive: %v", er)
	}

	log.Println("Recieved: ", rec.Msg)
}
