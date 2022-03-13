package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/amillert/go-grpc-course/grpc/greetpb"
	"github.com/amillert/go-grpc-course/grpc/sumpb"
	"google.golang.org/grpc"
)

func callUnaryGreet(c greetpb.GreetServiceClient) {
	fmt.Println("Starting unary Greet RPC")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Albert",
			LastName:  "Millert",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func callUnarySum(c sumpb.SumServiceClient) {
	fmt.Println("Starting unary Sum RPC")

	req := &sumpb.SumRequest{
		Sum: &sumpb.Sum{
			FirstNumber:  10,
			SecondNumber: 2,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Result)
}

func callStreamingGreet(c greetpb.GreetServiceClient) {
	fmt.Println("Starting server streaming Greet Multi RPC")

	req := &greetpb.GreetMultiRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Albert",
			LastName:  "Millert",
		},
	}
	// var res greetpb.GreetService_GreetMultiClient
	res, err := c.GreetMulti(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling streaming Greet Multi RPC: %v", err)
	}

	for {
		msg, err := res.Recv()

		if err == io.EOF {
			log.Fatalf("Reached end of stream")
			break
		} else if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		} else {
			log.Printf("Response from Greet Multi: %v\n", msg.GetResult())
		}
	}
}

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer cc.Close()

	gc := greetpb.NewGreetServiceClient(cc)
	callUnaryGreet(gc)

	sc := sumpb.NewSumServiceClient(cc)
	callUnarySum(sc)

	callStreamingGreet(gc)
}
