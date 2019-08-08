package db

import (
	"database/sql"

	"github.com/truauth/truauth/pkg/pgdb"
)

// AuthDbRequest request made to connect to authorization tabl
type AuthDbRequest struct {
	PGCreds *pgdb.DbCreds
}

// AuthorizationColumn authorization table response
type AuthorizationColumn struct {
	UserId            string
	AuthorizedClients string
}

// FindByUserID returns an authorized response
func (req *AuthDbRequest) FindByUserID(userID string) (*AuthorizationColumn, error) {
	db := pgdb.CreateSession(req.PGCreds)

	if db == nil {
		return nil, sql.ErrConnDone
	}
	defer db.Close()

	sqlQuery := `Select * from authorization_information where userId = $1`
	row := db.QueryRow(sqlQuery, userID)

	details := &AuthorizationColumn{}
	err := row.Scan(&details.UserId, &details.AuthorizedClients)
	if err != nil {
		return nil, err
	}

	return details, nil
}

// Create creates a authorization column
func (req *AuthDbRequest) Create(client *AuthorizationColumn) error {
	db := pgdb.CreateSession(req.PGCreds)
	if db == nil {
		return sql.ErrConnDone
	}
	defer db.Close()

	sqlQuery := `INSERT INTO authorization_information VALUES ($1, $2)`

	_, err := db.Exec(sqlQuery, client.UserId, client.AuthorizedClients)
	if err != nil {
		return err
	}

	return nil
}

// Update updates authorization column
func (req *AuthDbRequest) Update(client *AuthorizationColumn) error {
	db := pgdb.CreateSession(req.PGCreds)
	if db == nil {
		return sql.ErrConnDone
	}
	defer db.Close()

	sqlQuery := `UPDATE authorization_information SET authorizedclients = $1 WHERE userID = $2`
	_, err := db.Exec(sqlQuery, client.AuthorizedClients, client.UserId)
	if err != nil {
		return err
	}

	return nil
}
