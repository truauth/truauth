package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcIdentity "github.com/truauth/truauth/pkg/grpc-identity"
	"github.com/truauth/truauth/pkg/pgdb"
)

func main() {
	tcpPort := flag.String("p", "3005", "specified port to start the gRPC server on")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", *tcpPort))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	grpcIdentity.RegisterIdentityServer(grpcServer, &ServiceServer{
		PGCreds: pgdb.FetchCreds(),
	})
	reflection.Register(grpcServer)

	fmt.Println(fmt.Sprintf("gRPC service started on port %s", *tcpPort))
	if serveErr := grpcServer.Serve(listener); err != nil {
		panic(serveErr)
	}

	grpcServer.Serve(listener)
}
