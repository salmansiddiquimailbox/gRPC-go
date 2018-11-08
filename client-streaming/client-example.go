package main

import (
	pb "clientstreaming"
	"context"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50060", grpc.WithInsecure())
	defer conn.Close()
	gc := pb.NewGreeterClient(conn)

	stream, err := gc.Greet(context.Background()) //Only context is needed
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	i := 1
	for {
		msg := "Dummy" + strconv.Itoa(i)
		req := &pb.Request{Msg: msg}
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		time.Sleep(time.Second)
		i += 1
		if i == 6 {
			resp, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("Error! %v", err)
			}
			log.Printf("Received: %v", resp.Msg)
			break
		}

	}
}
