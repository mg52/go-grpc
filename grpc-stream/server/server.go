package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/mg52/go-grpc/grpc-stream/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) GreetAll(stream greetpb.GreetService_GreetAllServer) error {
	var output string
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.GreetManyTimesResponse{
				Result: output,
			})
		}
		if err != nil {
			return err
		}
		fmt.Printf("GreetAll function was invoked with %s %s.\n", point.GetGreeting().GetFirstName(), point.GetGreeting().GetLastName())
		output = output + fmt.Sprintf("%s %s\n", point.GetGreeting().GetFirstName(), point.GetGreeting().GetLastName())
	}
}
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v.\n", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " " + lastName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		log.Printf("Sent: %v", res)

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Print("Server started\n")
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
