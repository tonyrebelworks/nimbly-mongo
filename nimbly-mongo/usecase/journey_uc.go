package usecase

import (
	"chi-rest/model"
	"chi-rest/usecase/viewmodel"
	"log"
	"strconv"
	"strings"
	"time"
)

// GetAllJourney ...
func (uc UC) GetAllJourney(types string, maxID, limit int) ([]map[string]interface{}, viewmodel.SimplePaginationVM, error) {
	var (
		pagination viewmodel.SimplePaginationVM
	)

	data, err := model.JourneyOp.GetAll(uc.DB, types, maxID, limit)

	if len(data) > 0 {
		firstRecord := data[0]
		firstID := int(firstRecord.ID)
		lastRecord := data[len(data)-1]
		lastID := int(lastRecord.ID)
		pagination = SimplePaginationRes(types, maxID, firstID, lastID, limit)
	}

	if err != nil {
		return nil, pagination, err
	}

	resMap := make([]map[string]interface{}, 0)
	for _, a := range data {
		dataActivity, err := model.ActivityOp.GetByJourneyCode(uc.DB, a.Code)
		if err != nil {
			return nil, pagination, err
		}

		sitesRes := make([]map[string]interface{}, 0)
		site := a.Sites
		arrSites := strings.Split(site, "|")
		for i := range arrSites {
			sitesRes = append(sitesRes, map[string]interface{}{
				"siteID": arrSites[i],
			})
		}

		questionnairesRes := make([]map[string]interface{}, 0)
		question := a.Questionnaires
		arrQuestion := strings.Split(question, "|")
		for i := range arrQuestion {
			questionnairesRes = append(questionnairesRes, map[string]interface{}{
				"questionnaireID": arrQuestion[i],
			})
		}

		// emailRes := make([]map[string]interface{}, 0)
		email := a.EmailTo
		arrEmail := strings.Split(email, "|")
		// for i := range arrEmail {
		// 	emailRes = append(emailRes, map[string]interface{}{
		// 		"email": arrEmail[i],
		// 	})
		// }

		assignedAuditorRes := make([]map[string]interface{}, 0)
		assignAud := a.Salesman
		arrAssignAud := strings.Split(assignAud, "|")
		for i := range arrAssignAud {
			assignedAuditorRes = append(assignedAuditorRes, map[string]interface{}{
				"userID": arrAssignAud[i],
			})
		}

		activityRes := []viewmodel.ActivityVM{}
		for _, a := range dataActivity {
			tempRes := viewmodel.ActivityVM{
				UserID:   a.UserID,
				Username: a.Username,
				Datetime: a.Datetime.String,
			}
			activityRes = append(activityRes, tempRes)
		}

		// datesCustom := strings.Split(a.DatesCustom.String, ",")
		// daysOfWeek := strings.Split(a.DaysOfWeek.String, ",")
		// datesOfMonth := strings.Split(a.DatesOfMonth.String, ",")
		datesCustom := a.DatesCustom.String
		daysOfWeek := a.DaysOfWeek.String
		datesOfMonth := a.DatesOfMonth.String

		tmpDC := strings.Split(datesCustom, ",")
		datesCustomToInt := make([]int, 0, len(tmpDC))
		if datesCustom != "" {
			for _, raw := range tmpDC {
				v, err := strconv.Atoi(raw)
				if err != nil {
					log.Print(err)
					continue
				}
				datesCustomToInt = append(datesCustomToInt, v)
			}
		}

		tmpDow := strings.Split(daysOfWeek, ",")
		daysOfWeekToInt := make([]int, 0, len(tmpDow))
		if daysOfWeek != "" {
			for _, raw := range tmpDow {
				v, err := strconv.Atoi(raw)
				if err != nil {
					log.Print(err)
					continue
				}
				daysOfWeekToInt = append(daysOfWeekToInt, v)
			}
		}
		tmpDom := strings.Split(datesOfMonth, ",")
		datesOfMonthToInt := make([]int, 0, len(tmpDom))
		if datesOfMonth != "" {
			for _, raw := range tmpDom {
				v, err := strconv.Atoi(raw)
				if err != nil {
					log.Print(err)
					continue
				}
				datesOfMonthToInt = append(datesOfMonthToInt, v)
			}
		}

		// fmt.Println(values)
		resMap = append(resMap, map[string]interface{}{
			"id":                    a.ID,
			"code":                  a.Code,
			"journeyName":           a.JourneyName,
			"journeySchedule":       a.JourneySchedule,
			"datesCustom":           datesCustomToInt,
			"daysOfWeek":            daysOfWeekToInt,
			"datesOfMonth":          datesOfMonthToInt,
			"activity":              activityRes,
			"signatures":            a.Signatures,
			"requireSelfie":         a.RequireSelfie,
			"person":                a.Person.String,
			"startTime":             a.StartJourney.String,
			"endTime":               a.FinishJourney.String,
			"isDueToday":            true,
			"isDraft":               false,
			"isMakeUp":              false,
			"todayCompletedCount":   0,
			"completedCount":        0,
			"todayScheduleCount":    1,
			"isCompletedToday":      false,
			"isCompletedThisPeriod": false,
			"scheduleCount":         7,
			"isScheduleThisPeriod":  true,
			"createdAt":             a.CreatedAt.String,
			"createdBy":             a.CreatedBy.String,
			"updatedAt":             a.UpdatedAt.String,
			"updatedBy":             a.UpdatedBy.String,
			"sites":                 sitesRes,
			"questionnaires":        questionnairesRes,
			"emailTargets":          arrEmail,
			"assignedAuditor":       assignedAuditorRes,
		})
	}

	return resMap, pagination, err
}

// GetDetailJourney ...
func (uc UC) GetDetailJourney(code string) (viewmodel.JourneyPlanVM, error) {
	// url := "https://api.hellonimbly.com/v1.0/questionnaires"
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// kunci := ""
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("authorization", kunci)
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // log.Printf("body = %v", string(body))
	// type Data struct {
	// 	Title string
	// }
	// type Summary struct {
	// 	Message string
	// 	Data    []Data
	// }

	// var summary = new(Summary)
	// err3 := json.Unmarshal(body, &summary)
	// if err3 != nil {
	// 	fmt.Println("whoops:", err3)
	// 	//outputs: whoops: <nil>
	// }
	// // fmt.Println(summary.Message)
	// fmt.Println(summary.Data)

	data, err := model.JourneyOp.GetDetail(uc.DB, code)
	if err != nil {
		return viewmodel.JourneyPlanVM{}, err
	}

	// sitesRes := make([]viewmodel.SitesVM, 0)
	site := data.Sites
	arrSites := strings.Split(site, "|")
	// for i := range arrSites {
	// 	sitesRes = append(sitesRes, viewmodel.SitesVM{
	// 		SiteID: arrSites[i],
	// 	})
	// }

	// questionnairesRes := make([]viewmodel.QuestionnairesVM, 0)
	questionnaires := data.Questionnaires
	arrQuestionnaires := strings.Split(questionnaires, "|")
	// for i := range arrQuestionnaires {
	// 	questionnairesRes = append(questionnairesRes, viewmodel.QuestionnairesVM{
	// 		QuestionnairesID: arrQuestionnaires[i],
	// 	})
	// }

	// assignedAuditorRes := make([]viewmodel.AssignedAuditorVM, 0)
	assignAud := data.Salesman
	arrAssignAud := strings.Split(assignAud, "|")
	// for i := range arrAssignAud {
	// 	assignedAuditorRes = append(assignedAuditorRes, viewmodel.AssignedAuditorVM{
	// 		UserID: arrAssignAud[i],
	// 	})
	// }

	// emailRes := make([]viewmodel.EmailTargetsVM, 0)
	email := data.EmailTo
	arrEmail := strings.Split(email, "|")
	// for i := range arrEmail {
	// 	emailRes = append(emailRes, viewmodel.EmailTargetsVM{
	// 		Email: arrEmail[i],
	// 	})
	// }

	dataActivity, err := model.ActivityOp.GetByJourneyCode(uc.DB, code)
	if err != nil {
		return viewmodel.JourneyPlanVM{}, err
	}

	activityRes := []viewmodel.ActivityVM{}
	for _, a := range dataActivity {
		tempRes := viewmodel.ActivityVM{
			UserID:   a.UserID,
			Username: a.Username,
			Datetime: a.Datetime.String,
		}
		activityRes = append(activityRes, tempRes)

	}

	datesCustom := data.DatesCustom.String
	daysOfWeek := data.DaysOfWeek.String
	datesOfMonth := data.DatesOfMonth.String

	tmpDC := strings.Split(datesCustom, ",")
	datesCustomToInt := make([]int, 0, len(tmpDC))
	if datesCustom != "" {
		for _, raw := range tmpDC {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			datesCustomToInt = append(datesCustomToInt, v)
		}
	}

	tmpDow := strings.Split(daysOfWeek, ",")
	daysOfWeekToInt := make([]int, 0, len(tmpDow))
	if daysOfWeek != "" {
		for _, raw := range tmpDow {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			daysOfWeekToInt = append(daysOfWeekToInt, v)
		}
	}
	tmpDom := strings.Split(datesOfMonth, ",")
	datesOfMonthToInt := make([]int, 0, len(tmpDom))
	if datesOfMonth != "" {
		for _, raw := range tmpDom {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			datesOfMonthToInt = append(datesOfMonthToInt, v)
		}
	}

	res := viewmodel.JourneyPlanVM{
		ID:                    data.ID,
		Code:                  data.Code,
		JourneyName:           data.JourneyName,
		JourneySchedule:       data.JourneySchedule,
		DateCustom:            datesCustomToInt,
		DaysOfWeek:            daysOfWeekToInt,
		DateOfMonth:           datesOfMonthToInt,
		AssignedAuditor:       arrAssignAud,
		Sites:                 arrSites,
		EmailTargets:          arrEmail,
		Questionnaires:        arrQuestionnaires,
		Activity:              activityRes,
		Signatures:            data.Signatures,
		RequireSelfie:         data.RequireSelfie,
		Person:                data.Person.String,
		CreatedAt:             data.CreatedAt.String,
		UpdatedAt:             data.UpdatedAt.String,
		StartTime:             data.StartJourney.String,
		EndTime:               data.FinishJourney.String,
		IsDueToday:            true,
		IsDraft:               false,
		IsMakeUp:              false,
		TodayCompletedCount:   0,
		CompletedCount:        0,
		TodayScheduleCount:    1,
		IsCompletedToday:      false,
		IsCompletedThisPeriod: false,
		ScheduleCount:         7,
		IsScheduleThisPeriod:  true,
	}

	return res, err
}

// StoreJourney ...
func (uc UC) StoreJourney(
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

) (int64, error) {

	dt, err := model.JourneyOp.Store(uc.DB, code, journeyName, journeySchedule, datesCustom, daysOfWeek, datesOfMonth, salesman, sites, questionnaires, signatures, requireSelfie, person, emailTo, time.Now().UTC())
	return dt, err
}

// UpdateJourney ...
func (uc UC) UpdateJourney(
	code string,
	journeyName string,
	journeySchedule int64,
	// datesCustom string,
	// daysOfWeek string,
	// datesOfMonth string,
	salesman string,
	sites string,
	questionnaires string,
	signatures int64,
	requireSelfie int64,
	emailTo string,
	activity string,

) (int64, error) {
	dt, err := model.JourneyOp.Update(uc.DB, code, journeyName, journeySchedule, salesman, sites, questionnaires, signatures, requireSelfie, emailTo, activity, time.Now().UTC())
	return dt, err
}

// DeleteJourney ...
func (uc UC) DeleteJourney(code string) ([]*model.JourneyEntity, error) {

	dt, err := model.JourneyOp.DeleteJourney(uc.DB, code, time.Now().UTC())
	return dt, err
}

// UpdateTimeJourney ...
func (uc UC) UpdateTimeJourney(
	JourneyID string,
	StartTime string,
	EndTime string,
) ([]*model.JourneyEntity, error) {

	dt, err := model.JourneyOp.UpdateTimeJourney(uc.DB, JourneyID, StartTime, EndTime, time.Now().UTC())
	return dt, err
}

// GetDetailJourneyMobile ...
func (uc UC) GetDetailJourneyMobile(code string) (viewmodel.JourneyPlanMobileVM, error) {
	data, err := model.JourneyOp.GetDetail(uc.DB, code)
	if err != nil {
		return viewmodel.JourneyPlanMobileVM{}, err
	}

	sitesRes := make([]viewmodel.SitesVM, 0)
	site := data.Sites
	arrSites := strings.Split(site, "|")
	for i := range arrSites {
		sitesRes = append(sitesRes, viewmodel.SitesVM{
			SiteID: arrSites[i],
		})
	}

	questionnairesRes := make([]viewmodel.QuestionnairesVM, 0)
	questionnaires := data.Questionnaires
	arrQuestionnaires := strings.Split(questionnaires, "|")
	for i := range arrQuestionnaires {
		questionnairesRes = append(questionnairesRes, viewmodel.QuestionnairesVM{
			QuestionnairesID: arrQuestionnaires[i],
		})
	}

	assignedAuditorRes := make([]viewmodel.AssignedAuditorVM, 0)
	assignAud := data.Salesman
	arrAssignAud := strings.Split(assignAud, "|")
	for i := range arrAssignAud {
		assignedAuditorRes = append(assignedAuditorRes, viewmodel.AssignedAuditorVM{
			UserID: arrAssignAud[i],
		})
	}

	emailRes := make([]viewmodel.EmailTargetsVM, 0)
	email := data.EmailTo
	arrEmail := strings.Split(email, "|")
	for i := range arrEmail {
		emailRes = append(emailRes, viewmodel.EmailTargetsVM{
			Email: arrEmail[i],
		})
	}

	dataActivity, err := model.ActivityOp.GetByJourneyCode(uc.DB, code)
	if err != nil {
		return viewmodel.JourneyPlanMobileVM{}, err
	}

	activityRes := []viewmodel.ActivityVM{}
	for _, a := range dataActivity {
		tempRes := viewmodel.ActivityVM{
			UserID:   a.UserID,
			Username: a.Username,
			Datetime: a.Datetime.String,
		}
		activityRes = append(activityRes, tempRes)

	}

	res := viewmodel.JourneyPlanMobileVM{
		// ID:                    data.ID,
		Code:            data.Code,
		Name:            data.JourneyName,
		StartTime:       data.StartJourney.String,
		EndTime:         data.FinishJourney.String,
		Type:            "basic",
		Schedule:        data.JourneySchedule,
		Language:        "en",
		Signatures:      data.Signatures,
		SelfieSignature: data.RequireSelfie,
		Person:          data.Person.String,
		Questionnaires:  questionnairesRes,
		Sites:           sitesRes,
		// CreatedAt:             data.CreatedAt.String,
		// UpdatedAt:             data.UpdatedAt.String,
		IsDueToday:            true,
		IsDraft:               false,
		IsMakeUp:              false,
		TodayCompletedCount:   0,
		CompletedCount:        0,
		TodayScheduleCount:    1,
		IsCompletedToday:      false,
		IsCompletedThisPeriod: false,
		ScheduleCount:         7,
		IsScheduleThisPeriod:  true,
	}

	return res, err
}

// GetReportJourney ...
func (uc UC) GetReportJourney(code string) (viewmodel.ReportJourneyPlanVM, error) {
	data, err := model.JourneyOp.GetDetail(uc.DB, code)
	if err != nil {
		return viewmodel.ReportJourneyPlanVM{}, err
	}

	sitesRes := make([]viewmodel.SitesVM, 0)
	site := data.Sites
	arrSites := strings.Split(site, "|")
	for i := range arrSites {
		sitesRes = append(sitesRes, viewmodel.SitesVM{
			SiteID: arrSites[i],
		})
	}

	questionnairesRes := make([]viewmodel.QuestionnairesVM, 0)
	questionnaires := data.Questionnaires
	arrQuestionnaires := strings.Split(questionnaires, "|")
	for i := range arrQuestionnaires {
		questionnairesRes = append(questionnairesRes, viewmodel.QuestionnairesVM{
			QuestionnairesID: arrQuestionnaires[i],
		})
	}

	dataRep, err := model.ReportFirebaseOp.GetByJourneyCode(uc.DB, code)
	if err != nil {
		return viewmodel.ReportJourneyPlanVM{}, err
	}

	reportsRes := []viewmodel.ReportsVM{}
	for _, a := range dataRep {
		tempRes := viewmodel.ReportsVM{
			URL: a.URL,
		}
		reportsRes = append(reportsRes, tempRes)

	}

	// reportsRes := make([]viewmodel.ReportsVM, 0)
	// report := data.Reports
	// arrReports := strings.Split(report, "|")
	// for i := range arrReports {
	// 	sitesRes = append(reportsRes, viewmodel.ReportsVM{
	// 		URL: arrReports[i],
	// 	})
	// }

	assignedAuditorRes := make([]viewmodel.AssignedAuditorVM, 0)
	assignAud := data.Salesman
	arrAssignAud := strings.Split(assignAud, "|")
	for i := range arrAssignAud {
		assignedAuditorRes = append(assignedAuditorRes, viewmodel.AssignedAuditorVM{
			UserID: arrAssignAud[i],
		})
	}

	dataTraTi, err := model.TrackingTimeOp.GetByJourneyCode(uc.DB, "bq7e2l5hipgeufbrju5g0", "5qFKQb4kNJVFGsDBTp1NVrKojn12")
	if err != nil {
		return viewmodel.ReportJourneyPlanVM{}, err
	}

	// traTiRes := make([]viewmodel.TrackingTimeGPSVM, 0)
	// traTi := dataTraTi
	// arrTraTi := strings.Split(traTi, "|")
	// for i := range arrTraTi {
	// 	traTiRes = append(traTiRes, viewmodel.TrackingTimeGPSVM{
	// 		TrackingTime: arrTraTi[i],
	// 		Coordinates:  arrTraTi[i],
	// 	})
	// }

	traTiRes := []viewmodel.TrackingTimeGPSVM{}
	for _, a := range dataTraTi {
		tempRes := viewmodel.TrackingTimeGPSVM{
			TrackingTime: a.CreatedAt.String,
			Lat:          a.Latitude,
			Long:         a.Longitude,
		}
		traTiRes = append(traTiRes, tempRes)

	}

	datesCustom := data.DatesCustom.String
	daysOfWeek := data.DaysOfWeek.String
	datesOfMonth := data.DatesOfMonth.String

	tmpDC := strings.Split(datesCustom, ",")
	datesCustomToInt := make([]int, 0, len(tmpDC))
	if datesCustom != "" {
		for _, raw := range tmpDC {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			datesCustomToInt = append(datesCustomToInt, v)
		}
	}

	tmpDow := strings.Split(daysOfWeek, ",")
	daysOfWeekToInt := make([]int, 0, len(tmpDow))
	if daysOfWeek != "" {
		for _, raw := range tmpDow {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			daysOfWeekToInt = append(daysOfWeekToInt, v)
		}
	}
	tmpDom := strings.Split(datesOfMonth, ",")
	datesOfMonthToInt := make([]int, 0, len(tmpDom))
	if datesOfMonth != "" {
		for _, raw := range tmpDom {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			datesOfMonthToInt = append(datesOfMonthToInt, v)
		}
	}

	res := viewmodel.ReportJourneyPlanVM{
		ID:              data.ID,
		Code:            data.Code,
		JourneyName:     data.JourneyName,
		JourneySchedule: data.JourneySchedule,
		DateCustom:      datesCustomToInt,
		DaysOfWeek:      daysOfWeekToInt,
		DateOfMonth:     datesOfMonthToInt,
		AssignedAuditor: assignedAuditorRes,
		Sites:           sitesRes,
		Questionnaires:  questionnairesRes,
		Reports:         reportsRes,
		Signatures:      data.Signatures,
		StartJourney:    data.StartJourney.String,
		FinishJourney:   data.FinishJourney.String,
		CreatedAt:       data.CreatedAt.String,
		// TrackingTimeGPS: traTiRes,
	}

	return res, err
}

// AddTrackingTimeJourney ...
func (uc UC) AddTrackingTimeJourney(
	journeyCode string,
	userCode string,
	latitude string,
	longitude string,

) (int64, error) {

	dt, err := model.TrackingTimeOp.Store(uc.DB, journeyCode, userCode, latitude, longitude, time.Now().UTC())
	return dt, err
}

// GetAllJourneyMobile ...
func (uc UC) GetAllJourneyMobile() ([]viewmodel.GetAllJourneyPlanMobileVM, error) {
	data, err := model.JourneyOp.GetAll(uc.DB, "next", 0, 10000)
	if err != nil {
		return nil, err
	}

	resMap := make([]viewmodel.GetAllJourneyPlanMobileVM, 0)
	for _, a := range data {

		resMap = append(resMap, viewmodel.GetAllJourneyPlanMobileVM{
			Code:     a.Code,
			Name:     a.JourneyName,
			Schedule: a.JourneySchedule,
			Type:     "basic",
			Priority: true,
			Language: "en",
			// IsDueToday:            true,
			// IsDraft:               false,
			// IsMakeUp:              false,
			TodayCompletedCount: 0,
			CompletedCount:      0,
			// TodayScheduleCount:    1,
			// IsCompletedToday:      false,
			// IsCompletedThisPeriod: false,
			// ScheduleCount:         7,
			// IsScheduleThisPeriod:  true,
		})
	}

	return resMap, err
}

// AddURLFirebase ...
func (uc UC) AddURLFirebase(
	url string,
	journeyID string,

) (int64, error) {

	dt, err := model.ReportFirebaseOp.Store(uc.DB, url, journeyID, time.Now().UTC())
	return dt, err
}
