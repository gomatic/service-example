package main

import (
	"fmt"
	"log"

	gw "github.com/gomatic/service-example/model"
	"golang.org/x/net/context"
)

type service struct{}

func newService() gw.ExampleServiceServer {
	return new(service)
}

func (s *service) Echo(ctx context.Context, msg *gw.ExampleRequest) (*gw.ExampleResponse, error) {
	log.Print(msg)
	return &gw.ExampleResponse{Message: fmt.Sprintf("responding to: %s", msg.Message)}, nil
}
