package utils

import (
	r "gopkg.in/dancannon/gorethink.v2"
)

type Env struct {
	session	*r.Session
}