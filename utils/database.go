package utils

import (
	"log"
	"errors"

	r "gopkg.in/dancannon/gorethink.v2"
)

var Database = "my_db"

type DB struct {
	Name string
	Url string
	session	*r.Session
}

// Connect to the database.
func (db *DB) Connect() {
	conn, err := r.Connect(r.ConnectOpts{
		Address: db.Url,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Persist the connection.
	db.session = conn

	// Ensure a database is created
	r.DBCreate(db.Name).Exec(db.session)
}

// Close the session with the database.
func (db *DB) Disconnect() {
	db.session.Close()
}

// Creates a table in the database
func (db *DB) CreateTable(tableName string) error {
	var row bool
	res, _ := r.DB(Database).TableList().Contains(tableName).Run(db.session)
	for res.Next(&row){
	}
	defer res.Close()

	if row == false {
		err := r.DB(db.Name).TableCreate(tableName).Exec(db.session)
		if err != nil {
			return errors.New("Models. Could not create table.")
		}
	}
	return nil
}

// Interface to create documents in the database.
func (db *DB) Create(i interface{}, table string) (interface{}, error) {
	//r.Table("test").Insert(doc, r.InsertOpts{
	//    Conflict: func(id, oldDoc, newDoc r.Term) interface{} {
	//	return newDoc.Merge(map[string]interface{}{
	//	    "count": oldDoc.Add(newDoc.Field("count")),
	//	})
	//    },
	//})
	_, err := r.DB(db.Name).Table(table).Insert(i).Run(db.session)
	if err != nil {
		return nil, errors.New("Could not create records.")
	}
	return nil, nil
}

func (db *DB) FindById() {

}

// OLD. Not needed
// Database connection session handler
// Also Global access for other packages requesting DB access.
var ( DbSession *r.Session )

// Instantiate a new database object.
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

// OLD
func CreateTable(tableName string) (error) {
	var row bool
	res, _ := r.DB(Database).TableList().Contains(tableName).Run(DbSession)
	for res.Next(&row){
	}

	if row == false{
		err := r.DB(Database).TableCreate(tableName).Exec(DbSession)
		if err != nil {
			return errors.New("Models. Could not create table.")
		}
	}
	return nil
}



