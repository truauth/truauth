package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcIdentity "github.com/truauth/truauth/pkg/grpc-identity"
	"github.com/truauth/truauth/pkg/pgdb"
)

func main() {
	listener, err := net.Listen("tcp", ":4840")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	grpcIdentity.RegisterIdentityServer(grpcServer, &ServiceServer{
		PGCreds: pgdb.FetchCreds(),
	})
	reflection.Register(grpcServer)

	fmt.Println("server started on port 4840")
	if serveErr := grpcServer.Serve(listener); err != nil {
		panic(serveErr)
	}
}
