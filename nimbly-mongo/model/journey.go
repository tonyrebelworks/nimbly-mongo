package model

import (
	"database/sql"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

// JourneyEntity ...
type JourneyEntity struct {
	ID              uint           `db:"id" json:"id"`
	Code            string         `db:"code" json:"code"`
	JourneyName     string         `db:"journey_name" json:"journeyName"`
	JourneySchedule int            `db:"journey_schedule" json:"journeySchedule"`
	DatesCustom     sql.NullString `db:"dates_custom" json:"datesCustom"`
	DaysOfWeek      sql.NullString `db:"days_of_week" json:"daysOfWeek"`
	DatesOfMonth    sql.NullString `db:"dates_of_month" json:"datesOfMonth"`
	Salesman        string         `db:"salesman" json:"assignedAuditor"`
	// Salesman              []SalesmanEntity `db:"salesman" json:"assignedAuditor"`
	Sites                 string         `db:"sites" json:"sites"`
	Questionnaires        string         `db:"questionnaires" json:"questionnaires"`
	Signatures            string         `db:"signatures" json:"signatures"`
	RequireSelfie         bool           `db:"require_selfie" json:"requireSelfie"`
	Person                sql.NullString `db:"person" json:"person"`
	EmailTo               string         `db:"email_to" json:"emailTargets"`
	StartJourney          sql.NullString `db:"start_journey" json:"startTime"`
	FinishJourney         sql.NullString `db:"finish_journey" json:"endTime"`
	IsDueToday            sql.NullString `db:"is_due_today" json:"isDueToday"`
	IsDraft               sql.NullString `db:"is_draft" json:"isDraft"`
	IsMakeUp              sql.NullString `db:"is_makeup" json:"isMakeUp"`
	TodayCompletedCount   sql.NullString `db:"today_completed_count" json:"todayCompletedCount"`
	CompletedCount        sql.NullString `db:"completed_count" json:"completedCount"`
	TodayScheduleCount    sql.NullString `db:"today_schedule_count" json:"todayScheduleCount"`
	IsCompletedToday      sql.NullString `db:"is_completed_today" json:"isCompletedToday"`
	IsCompletedThisPeriod sql.NullString `db:"is_completed_this_period" json:"isCompletedThisPeriod"`
	ScheduleCount         sql.NullString `db:"schedule_count" json:"scheduleCount"`
	IsScheduleThisPeriod  sql.NullString `db:"is_schedule_this_period" json:"isScheduleThisPeriod"`
	CreatedAt             sql.NullString `db:"created_at" json:"createdAt"`
	CreatedBy             sql.NullString `db:"created_by" json:"createdBy"`
	UpdatedAt             sql.NullString `db:"updated_at" json:"updatedAt"`
	UpdatedBy             sql.NullString `db:"updated_by" json:"updatedBy"`
	DeletedAt             sql.NullString `db:"deleted_at" json:"deletedAt"`
}

type journeyOp struct{}

// JourneyOp ...
var JourneyOp = &journeyOp{}

// GetAll ...
func (op *journeyOp) GetAll(db *sqlx.DB, types string, maxID, limit int) ([]JourneyEntity, error) {
	var (
		err error
	)

	res := []JourneyEntity{}

	native := "SELECT id, code, journey_name, journey_schedule, dates_custom, days_of_week, dates_of_month, salesman, sites, questionnaires, signatures, require_selfie, person, email_to, start_journey, finish_journey, created_at, updated_at FROM journey_plan "

	if types == "next" {
		// maxID = 0 detect page 1
		if maxID == 0 {
			sql := native + " ORDER BY id DESC LIMIT ?"
			// fmt.Println(sql)

			err = db.Select(&res, sql, limit)

		} else {
			sql := native + " WHERE id > ? ORDER BY id DESC LIMIT ?"
			err = db.Select(&res, sql, maxID, limit)

		}
	} else {
		sql := native + " WHERE id > ? ORDER BY id DESC LIMIT ?"
		err = db.Select(&res, sql, maxID, limit)

	}

	// activeQ := "WHERE deleted_at IS NULL "
	// limitQ := "LIMIT 2"

	// err := db.Select(&res, "SELECT id, code, journey_name, journey_schedule, salesman, sites, questionnaires, signatures, require_selfie, person, email_to, start_journey, finish_journey, created_at, updated_at FROM journey_plan "+activeQ+limitQ)

	return res, err
}

// GetDetail ...
func (op *journeyOp) GetDetail(db *sqlx.DB, code string) (JourneyEntity, error) {
	var err error

	res := JourneyEntity{}
	err = db.Get(&res, "SELECT * FROM journey_plan WHERE code = ? LIMIT 1", code)

	return res, err
}

// Store ...
func (op *journeyOp) Store(
	db *sqlx.DB,
	code string,
	journeyName string,
	journeySchedule int64,
	datesCustom []string,
	daysOfWeek []string,
	datesOfMonth []string,
	salesman []string,
	sites []string,
	questionnaires []string,
	signatures int64,
	requireSelfie int64,
	person string,
	emailTo []string,
	// startJourney string,
	// finishJourney string,
	changedAt time.Time,

) (int64, error) {

	createdAt := changedAt.Format("2006-01-02 15:04:05")

	s := salesman
	salesmans := strings.Join(s, "|")

	si := sites
	sitess := strings.Join(si, "|")

	qu := questionnaires
	questionnairess := strings.Join(qu, "|")

	em := emailTo
	emailTos := strings.Join(em, "|")

	dc := datesCustom
	datesCustoms := strings.Join(dc, ",")

	dow := daysOfWeek
	daysOfWeeks := strings.Join(dow, ",")

	dom := datesOfMonth
	datesOfMonths := strings.Join(dom, ",")

	var sql = "INSERT INTO journey_plan (code, journey_name, journey_schedule, dates_custom, days_of_week, dates_of_month, salesman, sites, questionnaires, signatures, require_selfie,person, email_to,created_at) VALUES ( ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := db.Exec(sql, code, journeyName, journeySchedule, datesCustoms, daysOfWeeks, datesOfMonths, salesmans, sitess, questionnairess, signatures, requireSelfie, person, emailTos, createdAt)
	if err != nil {
		return 0, err
	}

	lID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lID, nil
}

// Update ...
func (op *journeyOp) Update(
	db *sqlx.DB,
	code string,
	journeyName string,
	journeySchedule int64,
	salesman string,
	sites string,
	questionnaires string,
	signatures int64,
	requireSelfie int64,
	emailTo string,
	activity string,
	// startJourney string,
	// finishJourney string,
	changedAt time.Time,

) (int64, error) {

	updatedAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "UPDATE journey_plan SET journey_name = ?,  journey_schedule = ?, salesman = ?, sites = ?, questionnaires = ?, signatures = ?, require_selfie = ?, email_to = ?, activity = ?, updated_at = ? WHERE code = ?"

	_, err := db.Exec(sql, journeyName, journeySchedule, salesman, sites, questionnaires, signatures, requireSelfie, emailTo, activity, updatedAt, code)
	if err != nil {
		return 0, err
	}

	return 0, err
}

// DeleteJourney ...
func (op *journeyOp) DeleteJourney(
	db *sqlx.DB,
	code string,
	changedAt time.Time,

) ([]*JourneyEntity, error) {

	deletedAt := changedAt.Format("2006-01-02 15:04:05")

	r := []*JourneyEntity{}
	sql := "UPDATE journey_plan SET deleted_at = ? WHERE code = ? "

	_, err := db.Exec(sql, deletedAt, code)
	return r, err
}

// UpdateTimeJourney ...
func (op *journeyOp) UpdateTimeJourney(
	db *sqlx.DB,
	code string,
	startTime string,
	endTime string,
	changedAt time.Time,
) ([]*JourneyEntity, error) {

	updatedAt := changedAt.Format("2006-01-02 15:04:05")
	r := []*JourneyEntity{}

	if startTime != "" {
		sql := "UPDATE journey_plan SET start_journey = ?, updated_at = ? WHERE code = ? "

		_, err := db.Exec(sql, startTime, updatedAt, code)
		return r, err
	}
	if endTime != "" {
		sql := "UPDATE journey_plan SET finish_journey = ?, updated_at = ? WHERE code = ? "

		_, err := db.Exec(sql, endTime, updatedAt, code)
		return r, err
	}
	sql := "UPDATE journey_plan SET updated_at = ? WHERE code = ? "

	_, err := db.Exec(sql, updatedAt, code)
	return r, err
}
