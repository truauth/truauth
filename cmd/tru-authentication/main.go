package main

import (
	"github.com/gin-gonic/gin"

	"github.com/truauth/truauth/cmd/tru-authentication/identity-grpc"
	"github.com/truauth/truauth/pkg/web-serve"
)

type request struct {
	SitePages webserve.Files
}

func main() {
	identityGRPCclient := identitygrpc.CreateGRPCClient("locahost:4820") // todo, arg this	
}