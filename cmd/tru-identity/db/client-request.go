package db

import (
	"database/sql"

	"github.com/truauth/truauth/pkg/pgdb"
	"github.com/truauth/truauth/pkg/grpc-identity"
)

// ClientDbRequest request made to connect to pgdb
type ClientDbRequest struct {
	PGCreds *pgdb.DbCreds
}

// FindByID finds a user identity given a username
func (req *ClientDbRequest) FindByID(ID string) (*grpcIdentity.ClientIdentity, error) {
	db := pgdb.CreateSession(req.PGCreds)

	if db == nil {
		return nil, sql.ErrConnDone
	}
	defer db.Close()

	sqlQuery := `Select * from client_information where id = $1`
	row := db.QueryRow(sqlQuery, ID)

	clientDetails := &grpcIdentity.ClientIdentity{}

	switch err := row.Scan(&clientDetails.ID, &clientDetails.Secret, &clientDetails.Name, &clientDetails.Redirect, &clientDetails.RelationID, &clientDetails.Callback); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		return clientDetails, nil 
	default:
		panic(err)
	}
}

// DirectCreate creates a client identity given the identity 
func (req *ClientDbRequest) DirectCreate(client *grpcIdentity.ClientIdentity) error {
	db := pgdb.CreateSession(req.PGCreds)
	if db == nil {
		return sql.ErrConnDone
	}
	defer db.Close()

	sqlQuery := `INSERT INTO client_information VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(sqlQuery, client.ID, client.Secret, client.Callback, client.Redirect, client.RelationID, client.Name)
	if err != nil {
		return err
	}
	return nil
}