package entitypayloads

type JobApplicationPayload struct {
	JobApplicationId  int    `json:"job_application_id"`
	JobId             int    `json:"job_id"`
	UserId            int    `json:"user_id"`
	CoverLetter       string `json:"cover_letter"`
	ApplicationStatus string `json:"application_status"`
	ActiveStatus      bool   `json:"active_status"`
}

type JobApplicationRequest struct {
	JobId        int    `json:"job_id"`
	UserId       int    `json:"user_id"`
	CoverLetter  string `json:"cover_letter"`
	ActiveStatus bool   `json:"active_status"`
}

type JobApplicationUpdate struct {
	JobApplicationId  int    `json:"job_application_id"`
	ApplicationStatus string `json:"application_status"`
}
