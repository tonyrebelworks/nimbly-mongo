package viewmodel

// JourneyPlanVM ...
type JourneyPlanVM struct {
	ID                    uint   ` json:"id"`
	Code                  string ` json:"code"`
	JourneyName           string ` json:"journeyName"`
	JourneySchedule       int    ` json:"journeySchedule"`
	DateCustom            []int  ` json:"dateCustom"`
	DaysOfWeek            []int  ` json:"daysOfWeek"`
	DateOfMonth           []int  ` json:"dateOfMonth"`
	Signatures            string ` json:"signatures"`
	RequireSelfie         bool   ` json:"requireSelfie"`
	Person                string ` json:"person"`
	StartTime             string ` json:"startTime"`
	EndTime               string ` json:"endTime"`
	IsDueToday            bool   ` json:"isDueToday"`
	IsDraft               bool   ` json:"isDraft"`
	IsMakeUp              bool   ` json:"isMakeUp"`
	TodayCompletedCount   int    ` json:"todayCompletedCount"`
	CompletedCount        int    ` json:"completedCount"`
	TodayScheduleCount    int    ` json:"todayScheduleCount"`
	IsCompletedToday      bool   ` json:"isCompletedToday"`
	IsCompletedThisPeriod bool   ` json:"isCompletedThisPeriod"`
	ScheduleCount         int    ` json:"scheduleCount"`
	IsScheduleThisPeriod  bool   ` json:"isScheduleThisPeriod"`
	CreatedAt             string ` json:"createdAt"`
	CreatedBy             string ` json:"createdBy"`
	UpdatedAt             string ` json:"updatedAt"`
	UpdatedBy             string ` json:"updatedBy"`
	// Sites                 []SitesVM           ` json:"sites"`
	Sites []string ` json:"sites"`
	// Questionnaires        []QuestionnairesVM  ` json:"questionnaires"`
	Questionnaires []string ` json:"questionnaires"`
	// AssignedAuditor []AssignedAuditorVM ` json:"assignedAuditor"`
	AssignedAuditor []string     ` json:"assignedAuditor"`
	EmailTargets    []string     ` json:"emailTargets"`
	Activity        []ActivityVM ` json:"activity"`
}

// SitesVM ...
type SitesVM struct {
	SiteID string `json:"siteID"`
}

// QuestionnairesVM ...
type QuestionnairesVM struct {
	QuestionnairesID string `json:"questionnairesID"`
	// EmailTargets          []string                `json:"emailTargets"`
	// EndTime               *string                 `json:"endTime"`
	// HasDeadline           bool                    `json:"hasDeadline"`
	// IsDoneToday           bool                    `json:"isDoneToday"`
	// IsQuestionnaireExists bool                    `json:"isQuestionnaireExists"`
	// IsScheduledToday      bool                    `json:"isScheduledToday"`
	// Key                   string                  `json:"key"`
	// QuestionnaireDetails  QuestionnairesDetailsVM `json:"questionnaireDetails"`
	// QuestionnaireTitle    string                  `json:"questionnaireTitle"`
	// ScheduledDates        map[string]interface{}  `json:"scheduledDates"`
	// // ScheduledDates   ScheduledDatesVM `json:"scheduledDates"`
	// SelfieSignatures []string `json:"selfieSignatures"`
	// Signatures       int      `json:"signatures"`
	// StartTime        *string  `json:"startTime"`
	// TitleLowercase   string   `json:"titleLowercase"`
	// TotalCompleted   int      `json:"totalCompleted"`
	// TotalScheduled   int      `json:"totalScheduled"`
	// UnfinishedDates  []string `json:"unfinishedDates"`
}

// QuestionnairesDetailsVM ...
type QuestionnairesDetailsVM struct {
	Key   string                       `json:"key"`
	Value QuestionnairesDetailsValueVM `json:"value"`
}

// ScheduledDatesVM ...
type ScheduledDatesVM struct {
	Date       string `json:"date"`
	IsComplete bool   `json:"isComplete"`
}

// QuestionnairesDetailsValueVM ...
type QuestionnairesDetailsValueVM struct {
	Disabled        bool                 `json:"disabled"`
	Latest          string               `json:"latest"`
	OrganizationKey string               `json:"organizationKey"`
	Questionnaire   QuestionnaireArrayVM `json:"questionnaire"`
	Tags            string               `json:"tags"`
	Title           string               `json:"title"`
	Versions        string               `json:"versions"`
}

// QuestionnaireArrayVM ...
type QuestionnaireArrayVM struct {
	DateCreated        string `json:"dateCreated"`
	DateUpdated        string `json:"dateUpdated"`
	Disabled           bool   `json:"disabled"`
	ModifiedBy         string `json:"modifiedBy"`
	QuestionnaireIndex string `json:"questionnaireIndex"`
	// Questions          []string `json:"questions"`
	Questions []QuestionVM `json:"questions"`
	Status    string       `json:"status"`
	Tags      string       `json:"tags"`
	Title     string       `json:"title"`
	Type      string       `json:"type"`
}

// QuestionVM ...
type QuestionVM struct {
	Answer         string                 `json:"answer"`
	AnswerRequired bool                   `json:"answerRequired"`
	Category       string                 `json:"category"`
	Comment        string                 `json:"comment"`
	Content        string                 `json:"content"`
	FlagLabel      map[string]interface{} `json:"flagLabel"`
	PhotoLimit     int                    `json:"photoLimit"`
	PhotoMinimum   int                    `json:"photoMinimum"`
	Reference      string                 `json:"reference"`
	Remedy         string                 `json:"remedy"`
	Score          int                    `json:"score"`
	ScoreWeight    int                    `json:"scoreWeight"`
	Sku            string                 `json:"sku"`
	Tags           map[string]interface{} `json:"tags"`
	Type           string                 `json:"type"`
	VideoLimit     int                    `json:"videoLimit"`
	VideoMinimum   int                    `json:"videoMinimum"`
}

// AssignedAuditorVM ...
type AssignedAuditorVM struct {
	UserID string `json:"userID"`
}

// EmailTargetsVM ...
type EmailTargetsVM struct {
	Email string `json:"email"`
}

// ActivityVM ...
type ActivityVM struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Datetime string `json:"datetime"`
}

// JourneyPlanMobileVM ...
type JourneyPlanMobileVM struct {
	// ID                    uint               ` json:"id"`
	Code                  string             ` json:"id"`
	Name                  string             ` json:"name"`
	StartTime             string             ` json:"startTime"`
	EndTime               string             ` json:"endTime"`
	Type                  string             ` json:"type"`
	Schedule              int                ` json:"schedule"`
	Priority              bool               ` json:"priority"`
	Language              string             ` json:"language"`
	Signatures            string             ` json:"signatures"`
	SelfieSignature       bool               ` json:"selfieSignature"`
	Person                string             ` json:"person"`
	Questionnaires        []QuestionnairesVM ` json:"questionnaires"`
	Sites                 []SitesVM          ` json:"site"`
	IsDueToday            bool               ` json:"isDueToday"`
	IsDraft               bool               ` json:"isDraft"`
	IsMakeUp              bool               ` json:"isMakeUp"`
	TodayCompletedCount   int                ` json:"todayCompletedCount"`
	CompletedCount        int                ` json:"completedCount"`
	TodayScheduleCount    int                ` json:"todayScheduleCount"`
	IsCompletedToday      bool               ` json:"isCompletedToday"`
	IsCompletedThisPeriod bool               ` json:"isCompletedThisPeriod"`
	ScheduleCount         int                ` json:"scheduleCount"`
	IsScheduleThisPeriod  bool               ` json:"isScheduleThisPeriod"`
	// CreatedAt             string             ` json:"createdAt"`
	// CreatedBy             string             ` json:"createdBy"`
	// UpdatedAt             string             ` json:"updatedAt"`
	// UpdatedBy             string             ` json:"updatedBy"`
}

// ReportJourneyPlanVM ...
type ReportJourneyPlanVM struct {
	ID              uint                ` json:"id"`
	Code            string              ` json:"code"`
	JourneyName     string              ` json:"journeyName"`
	JourneySchedule int                 ` json:"journeySchedule"`
	DateCustom      []int               ` json:"dateCustom"`
	DaysOfWeek      []int               ` json:"daysOfWeek"`
	DateOfMonth     []int               ` json:"dateOfMonth"`
	AssignedAuditor []AssignedAuditorVM ` json:"assignedAuditor"`
	Sites           []SitesVM           ` json:"sites"`
	Questionnaires  []QuestionnairesVM  ` json:"questionnaires"`
	// EmailTargets    []EmailTargetsVM    ` json:"emailTargets"`
	// Activity        []ActivityVM        ` json:"activity"`
	// Reports         []ReportsVM         ` json:"reports"`
	Reports         []ReportsVM         ` json:"reports"`
	TrackingTimeGPS []TrackingTimeGPSVM ` json:"trackingTimeGPS"`
	Signatures      string              ` json:"signatures"`
	StartJourney    string              ` json:"startJourney"`
	FinishJourney   string              ` json:"finishJourney"`
	CreatedAt       string              ` json:"createdAt"`
}

// TrackingTimeGPSVM ...
type TrackingTimeGPSVM struct {
	TrackingTime string `json:"trackingTime"`
	Lat          string `json:"lat"`
	Long         string `json:"long"`
}

// GetAllJourneyPlanMobileVM ...
type GetAllJourneyPlanMobileVM struct {
	Code     string ` json:"id"`
	Name     string ` json:"name"`
	Type     string ` json:"type"`
	Schedule int    ` json:"schedule"`
	Priority bool   ` json:"priority"`
	Language string ` json:"language"`
	// IsDueToday            bool   ` json:"isDueToday"`
	// IsDraft               bool   ` json:"isDraft"`
	// IsMakeUp              bool   ` json:"isMakeUp"`
	TodayCompletedCount int ` json:"todayCompletedCount"`
	CompletedCount      int ` json:"completedCount"`
	// TodayScheduleCount    int    ` json:"todayScheduleCount"`
	// IsCompletedToday      bool   ` json:"isCompletedToday"`
	// IsCompletedThisPeriod bool   ` json:"isCompletedThisPeriod"`
	// ScheduleCount         int    ` json:"scheduleCount"`
	// IsScheduleThisPeriod  bool   ` json:"isScheduleThisPeriod"`
}

// ReportsVM ...
type ReportsVM struct {
	URL string `json:"url"`
}
