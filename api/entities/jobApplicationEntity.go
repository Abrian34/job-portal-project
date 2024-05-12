package entities

var CreateJobApplicationTabel = "job_application_tbl"

type JobApplication struct {
	JobApplicationId  int    `gorm:"column:job_application_id;size:30;not null;primaryKey" json:"job_application_id"`
	JobId             int    `gorm:"column:job_id;size:30;not null;" json:"job_id"`
	UserId            int    `gorm:"column:user_id;size:30;not null;" json:"user_id"`
	CoverLetter       string `gorm:"column:cover_letter;size:50;not null" json:"cover_letter"`
	ApplicationStatus string `gorm:"column:application_status;size:50;not null" json:"application_status"`
	ActiveStatus      bool   `gorm:"column:active_status;not null;default:true" json:"active_status"`
}

func (*JobApplication) TableName() string {
	return CreateJobApplicationTabel
}
