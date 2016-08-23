package utils

import (
	"log"
	"errors"

	r "gopkg.in/dancannon/gorethink.v2"
)

var Database = "my_db"

type DB struct {
	*r.Session
}

// Database connection session handler
// Also Global access for other packages requesting DB access.
var ( DbSession *r.Session )

func ConnectDB(url string) (error){
	var err error
	DbSession, err = r.Connect(r.ConnectOpts{
		Address: url,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	r.DBCreate(Database).Exec(DbSession)
	return nil
}

func CreateTable(tableName string) (error) {
	// Connection from context
	err := r.DB(Database).TableCreate(tableName).Exec(DbSession)
	if err != nil {
		return errors.New("Models. Could not create table.")
	}
	return nil
}

