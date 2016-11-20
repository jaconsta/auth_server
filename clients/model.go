package clients

import (
	"errors"

	r "gopkg.in/dancannon/gorethink.v2"
	"github.com/jaconsta/users_ms/utils"
)


var tableName = "client_credentials"


// Create the clients table and create a default client.
func New() error {
	err := utils.CreateTable(tableName)
	if err != nil {
		return err
	}
	// Create a default client
	err = AddClient("client_one", "asdf1234")
	if err!= nil {
		return err
	}
	return err
}

/** CRUD METHODS **/

// Get a Client from it's clientId.
func findFromClientId (clientId string) ([]*client, error)  {
	var aClient []*client
	rows, err := r.DB(utils.Database).Table(tableName).Run(utils.DbSession)
	if err != nil {
		return nil, errors.New("Could not fetch the database.")
	}
	defer rows.Close()

	err = rows.All(&aClient)
	if err != nil {
		return nil, errors.New("Error parsing client response")
	}

	return aClient, nil
}

// Store the client.
func createClient (client client) error {

	_, err := r.DB(utils.Database).Table(tableName).Insert(client).Run(utils.DbSession)
	if err != nil {
		return errors.New("Could not create client.")
	}

	return err
}
