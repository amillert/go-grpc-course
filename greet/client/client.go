package main

import (
	"context"
	"fmt"
	"log"

	"github.com/amillert/go-grpc-course/greet/grpc/greetpb"
	"google.golang.org/grpc"
)

func callUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting unary RPC")

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

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	callUnary(c)
}
