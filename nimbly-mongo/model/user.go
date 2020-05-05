package model

import (
	"github.com/jmoiron/sqlx"
)

// User ...
type User struct {
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

type userOp struct{}

// UserOp ...
var UserOp = &userOp{}

// Get ...
func (op *userOp) Get(db *sqlx.DB) ([]*User, error) {
	r := []*User{}
	err := db.Select(&r, "SELECT name, email FROM borrowers LIMIT 10")

	return r, err
}
