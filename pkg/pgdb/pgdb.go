package pgdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // comment justifing it.. lol
	"github.com/truauth/truauth/pkg/settings"
)

// DbCreds struct for holding db creds
type DbCreds struct {
	Host     string
	User     string
	Password string
	Dbname   string
}

// CreateSession creates a pdb session
func CreateSession(creds *DbCreds) *sql.DB {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", creds.Host, creds.User, creds.Password, creds.Dbname)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return db
}

// FetchCreds fetches the db credentials from local en
func FetchCreds() *DbCreds {
	creds := &DbCreds{}

	settings.Init(creds, "postgres_env")

	return creds
}
