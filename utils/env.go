package utils

import (
	r "gopkg.in/dancannon/gorethink.v2"
)

// Config (manual)
const Database = "my_db"
var DbUrl = EnvOrDefault("DATABASE_URL", "127.0.0.1:32769")

// Env
type Env struct {
	DbSession	*r.Session
}