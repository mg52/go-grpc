package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/mg52/go-grpc/server-stream/greetpb"
	"google.golang.org/grpc"
)

var firstname string
var lastname string

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	fn := flag.String("f", "unnamed", "The name")
	ln := flag.String("l", "unnamed", "The name")
	flag.Parse()

	firstname = *fn
	lastname = *ln

	c := greetpb.NewGreetServiceClient(cc)

	doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: firstname,
			LastName:  lastname,
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		fmt.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}
}
