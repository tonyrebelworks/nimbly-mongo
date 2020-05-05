package model

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// ReportFirebaseEntity ...
type ReportFirebaseEntity struct {
	ID          uint           `db:"id" json:"id"`
	URL         string         `db:"url" json:"url"`
	JourneyCode string         `db:"journey_code" json:"journeyCode"`
	CreatedAt   sql.NullString `db:"created_at" json:"createdAt"`
}

type reportFirebaseOp struct{}

// ReportFirebaseOp ...
var ReportFirebaseOp = &reportFirebaseOp{}

// GetByJourneyCode ...
func (op *reportFirebaseOp) GetByJourneyCode(db *sqlx.DB, code string) ([]ReportFirebaseEntity, error) {
	var err error

	res := []ReportFirebaseEntity{}
	err = db.Select(&res, "SELECT id, url, journey_code, created_at FROM report_firebase WHERE journey_code = ? ", code)

	return res, err
}

// Store ...
func (op *reportFirebaseOp) Store(
	db *sqlx.DB,
	url string,
	journeyCode string,
	changedAt time.Time,

) (int64, error) {

	createdAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "INSERT INTO report_firebase (url,journey_code, created_at) VALUES ( ?,?, ?)"
	res, err := db.Exec(sql, url, journeyCode, createdAt)
	if err != nil {
		return 0, err
	}

	lID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lID, nil
}
