package user_model

type User struct {
	ID       int
	Email    string
	Roles    []Role
	PassHash []byte
}
