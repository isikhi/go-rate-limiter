package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/isikhi/go-rate-limiter/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRateLimitClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CheckRateLimit(ctx, &pb.CheckRequest{ClientId: "TestClient1Min"})
	if err != nil {
		log.Fatalf("an error occurred: %v", err)
	}
	log.Printf("Remaining token count: : %d", r.GetRemainingTokens())
	log.Printf("Overall Response: : %v", r)
}
