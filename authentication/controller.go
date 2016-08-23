package authentication

import (
	"golang.org/x/net/context"
	"net/http"

	"gopkg.in/dgrijalva/jwt-go.v3"
	"time"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var mySigningKey = []byte("This_is_top_secrep")
var tokenLife = int64(3600000) // One hour in millis

type authClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type responseToken struct {
	Token string `json:"token"`
}

type authBody struct {
	// Identifies the authentication operation to perform
	grantType string `json:"grant_type"`
	// User username when grant_type == password
	username string `json:"username"`
	// User password when grant_type == password
	password string `json:"password"`
	// Refresh token after expiring and when grant_type == refresh_token
	refreshToken string `json:"refresh_token"`
}
func makeTimestamp() int64 {
    return time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
}

func AuthController(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		authenticateUser(ctx, w, r)
	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
	return
}

// Performs a JWT authentication mechanism.
func authenticateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	// Authenticate client
	clientId, clientSecret, ok := r.BasicAuth()
	if ok != false {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	fmt.Println("Client credentials")
	fmt.Println(clientId)
	fmt.Println(clientSecret)

	// Parse the body and validate the authentication type
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad body.", http.StatusBadRequest)
		return
	}
	var authCredentials authBody

	switch authCredentials.grantType {
	case "password":
		fmt.Println("Password authentication type.")
	case "refresh_token":
		fmt.Println("Renew a poor old token.")
	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
	// return
	err = json.Unmarshal(body, &authCredentials)
	if err != nil{
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	claim := authClaim{
		"my_username",
		jwt.StandardClaims{
			ExpiresAt: makeTimestamp()+tokenLife,
			Issuer: "thisServer",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedJWT, err := token.SignedString(mySigningKey)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(responseToken{Token: signedJWT})
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}