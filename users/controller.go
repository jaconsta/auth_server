package users

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)


func UsersController (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUserdata(w, r)
	case http.MethodPost:
		addUser(w, r)
	default:
		http.Error(w, http.StatusText(405), 405)
	}
	return
}

func getUserdata(w http.ResponseWriter, r *http.Request){
	queries  := r.URL.Query()
	userEmail := queries.Get("email")

	user, err := getUser(userEmail)
	if err != nil {
		http.Error(w, "Bad username.", http.StatusBadRequest)
		return
	}

	res, _ := json.Marshal(user)
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}

func addUser( w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad body.", http.StatusBadRequest)
		return
	}

	// Serialize body
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil{
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	_, err = createUser(user)
	if err != nil{
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal("User created.")
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}