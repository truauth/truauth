package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"net"
	"os"
	"time"

	grpcIdentity "github.com/truauth/truauth/pkg/grpc-identity"
	"github.com/truauth/truauth/pkg/pgdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func init() {
	_ = godotenv.Load("configuration.env", "default.env", "postgres.env") // lets condense to one file, or a handful of files .
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3005"
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

	grpcIdentity.RegisterIdentityServer(grpcServer, &ServiceRequest{
		PGCreds: &pgdb.DbCreds{
			Host:     os.Getenv("POSTGRES_HOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Dbname:   os.Getenv("POSTGRES_DBNAME"),
		},
	})
	reflection.Register(grpcServer)

	fmt.Println(fmt.Sprintf("gRPC service started on port %s", *tcpPort))
	if serveErr := grpcServer.Serve(listener); err != nil {
		panic(serveErr)
	}

	grpcServer.Serve(listener)
}
