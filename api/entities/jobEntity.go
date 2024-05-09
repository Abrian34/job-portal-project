package entities

import "time"

var CreateJobTable = "job_tbl"

type Job struct {
	JobId          int       `gorm:"column:job_id;size:30;not null;primaryKey" json:"job_id"`
	JobCode        string    `gorm:"column:job_code;size:50;not null" json:"job_code"`
	EmployerId     int       `gorm:"column:employer_id;size:30;not null" json:"employer_id"`
	JobPostDate    time.Time `gorm:"column:job_post_date;not null;" json:"job_post_date"`
	CompanyId      int       `gorm:"column:company_id;size:30;not null" json:"company_id"`
	JobTitle       string    `gorm:"column:job_title;size:50;not null" json:"job_title"`
	JobDescription string    `gorm:"column:job_description;size:50;not null" json:"job_description"`
	JobLevel       string    `gorm:"column:job_level;size:50;not null" json:"job_level"`
	JobVacancy     string    `gorm:"column:job_vacancy;size:50;not null" json:"job_vacancy"`
	ActiveStatus   bool      `gorm:"column:active_status;size:50;not null" json:"active_status"`
}

func (*Job) TableName() string {
	return CreateJobTable
}
