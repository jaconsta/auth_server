package users

type User struct {
	ID string `gorethink:"id,omitempty"`
	FirstName string `json:"first_name" gorethink:"first_name"`
	LastName string `json:"last_name" gorethink:"last_name"`
	Email string `json:"email" gorethink:"email"`
}