package users

import (
	"github.com/jaconsta/users_ms/utils"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type User struct {
	ID string `gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
	Email string `json:"email" gorethink:"email"`
	Password string `json:"password" gorethink:"password"`
	Location string `json:"location" gorethink:"location"`
	BirthDate string `json:"birth_date" gorethink:"birth_date"`
}

func UsersIndex(env *utils.Env) http.Handler {

	listUsers := func(w http.ResponseWriter, r *http.Request) {
		users, err := utils.FetchAll(env.DbSession, "users")
		if err != nil{
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		res, _ := json.Marshal(users)
		fmt.Println("Not working")
		w.Header().Set("Content-type", "application/json")
		w.Write(res)
	}

	createUsers := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad body.", http.StatusBadRequest)
			return
		}

		// Serialize
		var user *User
		err = json.Unmarshal(body, &user)
		if err != nil{
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		// Persist
		var userId string
		userId, err = utils.Create(env.DbSession, user, "users")
		user.ID = userId
		// Resolve
		res, _ := json.Marshal(user)
		w.Header().Set("Content-type", "application/json")
		w.Write(res)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("*** New request ***")
		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			listUsers(w, r)
		case http.MethodPost:
			createUsers(w, r)
		case http.MethodPut:
			createUsers(w, r)
		default:
			http.Error(w, http.StatusText(405), 405)
		}
		return
	})
}