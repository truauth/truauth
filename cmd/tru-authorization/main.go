package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	grpcAuthorization "github.com/truauth/truauth/pkg/grpc-authorization"
	"github.com/truauth/truauth/pkg/pgdb"
)

func init() {
	_ = godotenv.Load("configuration.env", "default.env", "postgres.env")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3020"
	}

	tcpPort := flag.String("p", port, "specified port to start the gRPC server on")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", *tcpPort))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		PermitWithoutStream: true,
		MinTime:             1 * time.Minute,
	}))

	grpcAuthorization.RegisterAuthorizationServer(grpcServer, &ServiceRequest{
		PGCreds: &pgdb.DbCreds{
			Host:     os.Getenv("POSTGRES_HOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Dbname:   os.Getenv("POSTGRES_DBNAME"),
		},
	})

	fmt.Println(fmt.Sprintf("gRPC service started on port %s", *tcpPort))

	if serveErr := grpcServer.Serve(listener); err != nil {
		panic(serveErr)
	}
}
