package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaconsta/users_ms/utils"
	"github.com/jaconsta/users_ms/users"
	//"github.com/jaconsta/users_ms/authentication"
)

func main() {
	// Test the cipher function
	utils.Cypter()

	/**
	Database access
	 */
	db, err := utils.Connect(utils.Database, utils.DbUrl)
	if err != nil {
		log.Panic(err)
	}

	env := &utils.Env{DbSession: db,}

	// Ensure all tables are created at first
	_ = utils.CreateTable(env.DbSession, utils.Database, "users")

	// Urls
	http.Handle("/users/", users.UsersIndex(env))

	// Start server
	port := utils.EnvOrDefault("PORT", "8080")
	serverUrl := fmt.Sprintf(":%s", port)
	log.Println("Server starting at ", serverUrl)
	http.ListenAndServe(serverUrl, nil)
}
