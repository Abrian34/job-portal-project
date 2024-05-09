package entitypayloads

import "time"

type JobPayload struct {
	JobId          int       `json:"job_id" parent_entity:"job_tbl" main_table:"job_tbl"`
	JobCode        string    `json:"job_code" parent_entity:"job_tbl"`
	EmployerId     int       `json:"employer_id"`
	JobPostDate    time.Time `json:"job_post_date"`
	CompanyId      int       `json:"company_id"`
	JobTitle       string    `json:"job_title"`
	JobDescription string    `json:"job_description"`
	JobLevel       string    `json:"job_level"`
	JobVacancy     string    `json:"job_vacancy"`
	ActiveStatus   string    `json:"active_status"`
}
