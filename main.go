package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaconsta/users_ms/utils"
	"github.com/jaconsta/users_ms/clients"
	"github.com/jaconsta/users_ms/users"
	//"github.com/jaconsta/users_ms/authentication"
)

type application struct {
	db utils.DB
}

func main() {
	// Test the cipher function
	utils.Cypter()

	/**
	Database access
	 */
	dbUrl := utils.EnvOrDefault("DATABASE_URL", "localhost:32769")
	database := utils.DB{Name: "my_db", Url: dbUrl}
	err := utils.ConnectDB(dbUrl)
	if err != nil {
		log.Panic(err)
	}
	// Ensure all tables are created at first
	users.New()
	clients.New()

	// Urls
	http.HandleFunc("/users", users.UsersController)

	// Start server
	port := utils.EnvOrDefault("PORT", "8080")
	serverUrl := fmt.Sprintf(":%s", port)
	log.Println("Server starting at ", serverUrl)
	http.ListenAndServe(serverUrl, nil)
}
