package main

import "github.com/truauth/truauth/pkg/pgdb"

// ServiceServer service request struct used to pass service information.
type ServiceServer struct {
	PGCreds *pgdb.DbCreds
}