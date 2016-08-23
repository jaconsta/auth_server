package clients

import (
	"errors"

	r "gopkg.in/dancannon/gorethink.v2"
	"github.com/jaconsta/users_ms/utils"
)


var tableName = "client_credentials"

func findFromClientId (clientId string) ([]*client, error)  {
	var aClient *client
	rows, err := r.DB(utils.Database).Table(tableName).Run(utils.DbSession)
	if err != nil {
		return errors.New("Could not fetch the database.")
	}

	err = rows.All(&aClient)

	return aClient, nil
}
