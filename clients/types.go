package clients

type client struct {
	Id string `json:"-id"`
	ClientId string `json:"client_id" gorethink: "client_id"`
	Secret string `json:"client_secret" gorethink: "client_secret"`
}

var TableName = "client_credentials"
