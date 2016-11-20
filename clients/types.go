package clients

type client struct {
	Id string `json:"-id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
	ClientId string `json:"client_id" gorethink:"client_id"`
	Secret string `json:"client_secret" gorethink:"client_secret"`
}
