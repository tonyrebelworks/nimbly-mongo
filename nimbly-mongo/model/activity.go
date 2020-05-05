package model

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// ActivityEntity ...
type ActivityEntity struct {
	ID          uint           `db:"id" json:"id"`
	Code        string         `db:"code" json:"code"`
	UserID      string         `db:"user_code" json:"userID"`
	Username    string         `db:"username" json:"username"`
	JourneyCode string         `db:"journey_code" json:"journeyCode"`
	Datetime    sql.NullString `db:"created_at" json:"createdAt"`
}

type activityOp struct{}

// ActivityOp ...
var ActivityOp = &activityOp{}

// GetByJourneyCode ...
func (op *activityOp) GetByJourneyCode(db *sqlx.DB, code string) ([]ActivityEntity, error) {
	var err error

	res := []ActivityEntity{}
	err = db.Select(&res, "SELECT id, code, user_code, username, journey_code, created_at FROM activity WHERE journey_code = ? ", code)

	return res, err
}
