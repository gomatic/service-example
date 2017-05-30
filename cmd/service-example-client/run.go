package main

import (
	"log"
	"os"
	"strings"

	pb "github.com/gomatic/service-example/model"
	"github.com/gomatic/servicer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//
func run(settings servicer.Settings) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(settings.Rpc.String(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewExampleServiceClient(conn)

	r, err := c.Echo(context.Background(), &pb.ExampleRequest{Message: strings.Join(os.Args[1:], " ")})
	if err != nil {
		log.Fatalf("service failed: %v", err)
	}
	log.Printf("Response: %#v", r)

	return nil
}
