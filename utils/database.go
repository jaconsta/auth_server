package utils

import (
	"log"
	"errors"

	r "gopkg.in/dancannon/gorethink.v2"
)

// Connect to the database
func Connect(dbName string, url string) (*r.Session, error) {
	conn, err := r.Connect(r.ConnectOpts{
		Address: url,
		Database: dbName,
	})
	if err != nil {
		log.Fatal(err)
	}
	// Ensure a database is created
	r.DBCreate(dbName).Exec(conn)

	return conn, nil
}

// Disconnect from the database
func Disconnect(session *r.Session) {
	session.Close()
}

// Check if a collection exists in the database and create it.
func CreateTable(session *r.Session, database string, table string) error{
	var row bool
	res, _ := r.DB(database).TableList().Contains(table).Run(session)
	for res.Next(&row){
	}
	defer res.Close()

	if row == false {
		err := r.DB(database).TableCreate(table).Exec(session)
		if err != nil {
			return errors.New("Models. Could not create table.")
		}
	}
	return nil
}

// Insert a new document in the database.
func Create(session *r.Session, i interface{}, table string) (string, error) {
	result, err := r.Table(table).Insert(i).RunWrite(session)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Could not create records.")
	}

	// fmt.Println("*** Document inserted ***")
	// printObj(result)
	return result.GeneratedKeys[0], nil
}

// Get all records in table
func FetchAll(session *r.Session, table string) (interface{}, error){
	rows, err := r.Table(table).Run(session)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Could not fetch records.")
	}

	// Should use reflection to work
	// https://play.golang.org/p/kkBu56JYb8
	return rows, nil
}

/*
func printObj(v interface{}) {
    vBytes, _ := json.Marshal(v)
    fmt.Println(string(vBytes))
}
*/
