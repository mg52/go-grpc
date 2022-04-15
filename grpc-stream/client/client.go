package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/mg52/go-grpc/grpc-stream/greetpb"
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

	option := flag.Int64("o", 0, "0 is for server streaming, 1 is for client streaming")
	fn := flag.String("f", "unnamed", "Firstname")
	ln := flag.String("l", "unnamed", "Lastname")
	flag.Parse()

	firstname = *fn
	lastname = *ln

	c := greetpb.NewGreetServiceClient(cc)
	if *option == 0 {
		doServerStreaming(c)
	} else if *option == 1 {
		doClientStreaming(c)
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	var reqList []*greetpb.GreetManyTimesRequest

	req1 := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "ahmet",
			LastName:  "xxxx",
		},
	}
	reqList = append(reqList, req1)
	req2 := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "mehmet",
			LastName:  "yyyy",
		},
	}
	reqList = append(reqList, req2)
	req3 := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "kamil",
			LastName:  "mmmm",
		},
	}
	reqList = append(reqList, req3)
	req4 := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "ayse",
			LastName:  "qqqq",
		},
	}
	reqList = append(reqList, req4)
	req5 := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "mustafa",
			LastName:  "gggg",
		},
	}
	reqList = append(reqList, req5)

	stream, err := c.GreetAll(context.Background())
	if err != nil {
		log.Fatalf("%v.GreetAll(_) = _, %v", c, err)
	}
	for _, req := range reqList {
		if err := stream.Send(req); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, req, err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}

	fmt.Printf("Response from GreetManyTimes:\n")
	fmt.Printf("%v\n", reply.GetResult())
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
