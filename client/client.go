package main

import (
	"context"
	"fmt"
	"log"

	"github.com/amillert/go-grpc-course/grpc/greetpb"
	"github.com/amillert/go-grpc-course/grpc/sumpb"
	"google.golang.org/grpc"
)

func callUnaryGreet(c greetpb.GreetServiceClient) {
	fmt.Println("Starting unary greet RPC")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Albert",
			LastName:  "Millert",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Errorw while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func callUnarySum(c sumpb.SumServiceClient) {
	fmt.Println("Starting unary sum RPC")

	req := &sumpb.SumRequest{
		Sum: &sumpb.Sum{
			FirstNumber:  10,
			SecondNumber: 2,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Errorw while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Result)
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
}
