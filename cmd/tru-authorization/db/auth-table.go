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
	ID                string
	AuthorizedClients []string
}

// FindByUserID returns an authorized response
func (req *AuthDbRequest) FindByUserID(ID string) (*AuthorizationColumn, error) {
	db := pgdb.CreateSession(req.PGCreds)

	if db == nil {
		return nil, sql.ErrConnDone
	}
	defer db.Close()

	sqlQuery := `Select * from authorization_information where id = $1`
	row := db.QueryRow(sqlQuery, ID)

	details := &AuthorizationColumn{}

	switch err := row.Scan(&details.ID, &details.AuthorizedClients); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		return details, nil
	default:
		panic(err)
	}
}

// Create creates a authorization column
func (req *AuthDbRequest) Create(client *AuthorizationColumn) error {
	db := pgdb.CreateSession(req.PGCreds)
	if db == nil {
		return sql.ErrConnDone
	}
	defer db.Close()

	sqlQuery := `INSERT INTO authorization_information where id = $1`

	_, err := db.Exec(sqlQuery, client.ID, client.AuthorizedClients)
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

	sqlQuery := `UPDATE authorization_information SET authorizedClients = $1 WHERE id = $2`
	_, err := db.Exec(sqlQuery, client.AuthorizedClients, client.ID)
	if err != nil {
		return err
	}

	return nil
}
