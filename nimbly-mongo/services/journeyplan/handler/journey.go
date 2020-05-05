package handler

import (
	"chi-rest/services/journeyplan/request"
	"chi-rest/usecase"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/xid"
)

// GetAllJourney ...
// GetStringByInt example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {string} string	MsgSuccess
// @Router / [get]
func (h *Contract) GetAllJourney(w http.ResponseWriter, r *http.Request) {
	var (
		types string
		maxID int
		limit int
		err   error
	)

	types = "next"
	maxID = 0
	limit = 10000

	// types = r.URL.Query().Get("types")
	// if types != "prev" && types != "next" {
	// 	h.SendBadRequest(w, "Invalid type value")
	// 	return
	// }
	// maxID, err = strconv.Atoi(r.URL.Query().Get("max_id"))
	// if err != nil {
	// 	h.SendBadRequest(w, "Invalid last id value")
	// 	return
	// }
	// limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
	// if err != nil {
	// 	h.SendBadRequest(w, "Invalid limit value")
	// 	return
	// }

	if err != nil {
		fmt.Println(err)
		return
	}
	res, pagination, err := usecase.UC{h.App}.GetAllJourney(types, maxID, limit)

	h.SendSuccess(w, res, pagination)
	return
}

// GetDetailJourney ...
func (h *Contract) GetDetailJourney(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.GetDetailJourney(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// AddJourney ...
func (h *Contract) AddJourney(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddJourneyRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	// if err = h.Validate.Struct(req); err != nil {
	// 	h.SendRequestValidationError(w, err.(validator.ValidationErrors))
	// 	return
	// }

	code := xid.New().String()

	JourneyName := req.JourneyName
	JourneySchedule := req.JourneySchedule
	// Salesman := req.Salesman

	if len(req.AssignedAuditor) > 0 {

	}
	assignedAuditors := make([]string, 0)
	for _, aa := range req.AssignedAuditor {
		assignedAuditors = append(assignedAuditors, aa.UserID)
	}

	if len(req.Sites) > 0 {

	}
	sitess := make([]string, 0)
	for _, si := range req.Sites {
		sitess = append(sitess, si.SiteID)
	}

	if len(req.Questionnaires) > 0 {

	}
	questionnairess := make([]string, 0)
	for _, qu := range req.Questionnaires {
		questionnairess = append(questionnairess, qu.QuestionnaireID)
	}

	if len(req.EmailTo) > 0 {

	}
	emails := make([]string, 0)
	for _, em := range req.EmailTo {
		emails = append(emails, em.Email)
	}

	if len(req.DatesCustom) > 0 {

	}
	datesCustom := make([]string, 0)
	for _, dc := range req.DatesCustom {
		datesCustom = append(datesCustom, dc.DatesCustom)
	}

	if len(req.DaysOfWeek) > 0 {

	}
	daysOfWeek := make([]string, 0)
	for _, dow := range req.DaysOfWeek {
		daysOfWeek = append(daysOfWeek, dow.DaysOfWeek)
	}

	if len(req.DatesOfMonth) > 0 {

	}
	datesOfMonth := make([]string, 0)
	for _, dom := range req.DatesOfMonth {
		datesOfMonth = append(datesOfMonth, dom.DateOfMonth)
	}

	// Sites := req.Sites
	// Questionnaires := req.Questionnaires
	Signatures := req.Signatures
	RequireSelfie := req.RequireSelfie
	Person := req.Person
	// EmailTo := req.EmailTo
	// Activity := req.Activity
	// StartJourney := req.StartJourney
	// FinishJourney := req.FinishJourney

	// mdl := usecase.UC{h.App}.StoreJourney(code, JourneyName, JourneySchedule, Salesman, Sites, Questionnaires, Signatures, RequireSelfie, EmailTo, Activity, StartJourney, FinishJourney)

	lastID, err := usecase.UC{h.App}.StoreJourney(
		code,
		JourneyName,
		JourneySchedule,
		datesCustom,
		daysOfWeek,
		datesOfMonth,
		assignedAuditors,
		sitess,
		questionnairess,
		Signatures,
		RequireSelfie,
		Person,
		emails,
		// Activity,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, lastID)
	return
}

// UpdateJourney ...
func (h *Contract) UpdateJourney(w http.ResponseWriter, r *http.Request) {
	var err error

	req := request.UpdateJourneyRequest{}
	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	// if err = h.Handler.Validate.Struct(req); err != nil {
	// 	h.SendRequestValidationError(w, err.(validator.ValidationErrors))
	// 	return
	// }

	JourneyName := req.JourneyName
	JourneySchedule := req.JourneySchedule
	Salesman := req.Salesman
	Sites := req.Sites
	Questionnaires := req.Questionnaires
	Signatures := req.Signatures
	RequireSelfie := req.RequireSelfie
	EmailTo := req.EmailTo
	Activity := req.Activity

	code := chi.URLParam(r, "code")
	mdl := usecase.UC{h.App}

	_, err = mdl.UpdateJourney(
		code,
		JourneyName,
		JourneySchedule,
		Salesman,
		Sites,
		Questionnaires,
		Signatures,
		RequireSelfie,
		EmailTo,
		Activity,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, nil)
	return
}

// DeleteJourney ...
func (h *Contract) DeleteJourney(w http.ResponseWriter, r *http.Request) {

	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.DeleteJourney(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// UpdateTimeJourney ...
func (h *Contract) UpdateTimeJourney(w http.ResponseWriter, r *http.Request) {
	var err error

	req := request.UpdateTimeJourneyRequest{}
	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	JourneyID := req.JourneyID
	StartTime := req.StartTime
	EndTime := req.EndTime

	mdl := usecase.UC{h.App}

	_, err = mdl.UpdateTimeJourney(
		JourneyID,
		StartTime,
		EndTime)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, nil)
	return
}

// GetDetailJourneyMobile ...
func (h *Contract) GetDetailJourneyMobile(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.GetDetailJourneyMobile(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// GetReportJourney ...
func (h *Contract) GetReportJourney(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.GetReportJourney(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// AddTrackingTimeJourney ...
func (h *Contract) AddTrackingTimeJourney(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddTrackingTimeJourneyRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	// if err = h.Validate.Struct(req); err != nil {
	// 	h.SendRequestValidationError(w, err.(validator.ValidationErrors))
	// 	return
	// }

	JourneyID := req.JourneyID
	Latitude := req.Latitude
	Longitude := req.Longitude
	UserCode := "5qFKQb4kNJVFGsDBTp1NVrKojn12"

	lastID, err := usecase.UC{h.App}.AddTrackingTimeJourney(
		JourneyID,
		UserCode,
		Latitude,
		Longitude,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, lastID)
	return
}

// GetAllJourneyMobile ...
func (h *Contract) GetAllJourneyMobile(w http.ResponseWriter, r *http.Request) {
	res, err := usecase.UC{h.App}.GetAllJourneyMobile()
	if err != nil {
		fmt.Println(err)
		return
	}
	h.SendSuccess(w, res, nil)
	return
}

// AddURLFirebase ...
func (h *Contract) AddURLFirebase(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddURLFirebaseRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	// if err = h.Validate.Struct(req); err != nil {
	// 	h.SendRequestValidationError(w, err.(validator.ValidationErrors))
	// 	return
	// }

	URL := req.URL
	JourneyID := req.JourneyID

	lastID, err := usecase.UC{h.App}.AddURLFirebase(
		URL,
		JourneyID,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, lastID)
	return
}
