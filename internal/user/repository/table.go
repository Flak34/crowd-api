package user_repository

type PersonTable struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	PassHash []byte `db:"pass_hash"`
}

type RoleTable struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
