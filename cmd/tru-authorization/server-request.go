package main

import "github.com/truauth/truauth/pkg/pgdb"

// ServiceRequest service request struct used to pass service information.
type ServiceRequest struct {
	PGCreds *pgdb.DbCreds
}
