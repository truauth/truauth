package db

import (
	"database/sql"

	grpcIdentity "github.com/truauth/truauth/pkg/grpc-identity"
	"github.com/truauth/truauth/pkg/pgdb"
)

// UserDbRequest request made to connect to pgdb
type UserDbRequest struct {
	PGCreds *pgdb.DbCreds
}

// FindByUsername finds a user identity given a username
func (req *UserDbRequest) FindByUsername(username string) (*grpcIdentity.UnsafeUserIdentity, error) {
	db := pgdb.CreateSession(req.PGCreds)
	defer db.Close()

	sqlQuery := `Select * from user_credentials where username = $1`
	row := db.QueryRow(sqlQuery, username)

	model := &grpcIdentity.UnsafeUserIdentity{}

	err := row.Scan(&model.Username, &model.Password, &model.Email, &model.Type)

	if err != nil {
		return nil, err
	}

	return model, nil
}

// FindByEmail finds a user identity given a email
func (req *UserDbRequest) FindByEmail(email string) (*grpcIdentity.UnsafeUserIdentity, error) {
	db := pgdb.CreateSession(req.PGCreds)
	defer db.Close()

	sqlQuery := `Select * from user_credentials where email = $1`
	row := db.QueryRow(sqlQuery, email)

	model := &grpcIdentity.UnsafeUserIdentity{}

	err := row.Scan(&model.Username, &model.Password, &model.Type, &model.Email)

	if err != nil {
		return nil, err
	}

	return model, nil
}

// DirectCreate creates a user identity given the identity
func (req *UserDbRequest) DirectCreate(user *grpcIdentity.UnsafeUserIdentity) error {
	db := pgdb.CreateSession(req.PGCreds)
	if db == nil {
		return sql.ErrConnDone
	}
	defer db.Close()

	sqlQuery := `INSERT INTO user_credentials VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(sqlQuery, user.Username, user.Password, user.Email, user.Type)
	if err != nil {
		return err
	}

	return nil
}
