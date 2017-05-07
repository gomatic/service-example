package main

import (
	"context"

	gw "github.com/gomatic/service-example/model"
	"github.com/gomatic/servicer"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

//
func run(ctx context.Context, settings servicer.Settings, rpc *grpc.Server) (*runtime.ServeMux, error) {

	gw.RegisterExampleServiceServer(rpc, newService())

	rpcHost := settings.Rpc.String()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := gw.RegisterExampleServiceHandlerFromEndpoint(ctx, mux, rpcHost, opts); err != nil {
		return nil, err
	}

	return mux, nil
}
