package users

import (
    	r "gopkg.in/dancannon/gorethink.v2"
	"github.com/jaconsta/users_ms/utils"
)

var tableName = "users"

/**
	Configuration
 */
// Create the users table
func New () (error){
	err := utils.CreateTable(tableName)
	return err
}

/**
	CRUD METHODS
 */

// Store an user
func createUser(user User) (*User, error){

	res, err := r.DB(utils.Database).Table(tableName).Insert(user).RunWrite(utils.DbSession)
	if err != nil {
		return nil, err
	}

	if user.ID == "" && len(res.GeneratedKeys) == 1 {
		user.ID = res.GeneratedKeys[0]
	}

	return nil, nil
}

// Get the user from it's email
func getUser(serEmail string) (*User, error) {
	return nil, nil
}
