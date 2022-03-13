package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/amillert/go-grpc-course/grpc/greetpb"
	"github.com/amillert/go-grpc-course/grpc/sumpb"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet invoked with %v\n", req)

	firstName := req.GetGreeting().GetFirstName()
	res := &greetpb.GreetResponse{Result: "Hello " + firstName}

	return res, nil
}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("Sum invoked with %v\n", req)

	firstNum, secondNum := req.GetSum().GetFirstNumber(), req.GetSum().GetSecondNumber()
	res := &sumpb.SumResponse{Result: firstNum + secondNum}

	return res, nil
}

func main() {
	l, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listed %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
