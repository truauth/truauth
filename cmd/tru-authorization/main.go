package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"

	grpcAuthorization "github.com/truauth/truauth/pkg/grpc-authorization"
	"github.com/truauth/truauth/pkg/pgdb"
)

func main() {
	port := flag.String("p", "3020", "specified port to start the gRPC server on")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	grpcAuthorization.RegisterAuthorizationServer(grpcServer, &ServiceRequest{
		PGCreds: pgdb.FetchCreds(),
	})

	fmt.Println(fmt.Sprintf("gRPC service started on port %s", *port))

	if serveErr := grpcServer.Serve(listener); err != nil {
		panic(serveErr)
	}
}
