package pgdb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq" // comment justifing it.. lol
	"os"
)

// DbCreds struct for holding db creds
type DbCreds struct {
	Host 		string
	User 		string
	Password 	string
	Dbname		string
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
	file, _ := os.Open("./postgres_env.json")
	defer file.Close()

	decoder := json.NewDecoder(file)

	creds := &DbCreds{}
	err := decoder.Decode(creds)

	if err != nil {
		return nil
	}

	return creds
}