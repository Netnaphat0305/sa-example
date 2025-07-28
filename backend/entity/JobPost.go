package entity

import (
	"time"

	"gorm.io/gorm"
)

type StatusEnum string 

const (
	Open   StatusEnum = "Open"
	Close  StatusEnum = "Close"
)
type JobPost struct {
	gorm.Model
	Title			string			`gorm:"type:varchar(100);not null" json:"title"`	
	Description		string			`gorm:"type:text;not null" json:"description"`
	Deadline		time.Time		`gorm:"type:date;not null" json:"deadline"`
	Status			StatusEnum		`gorm:"type:enum('Open','Close');not null" json:"status"`
	Employer		*Employer		`gorm:"foreignKey: employer_id" json:"employer"`
	JobCategory		*JobCategory	`gorm:"foreignKey: jobcategory_id" json:"jobcategory"`
	Location		*Location		`gorm:"foreignKey: location_id" json:"location"`
	EmploymentType	*EmploymentType	`gorm:"foreignKey: employer_id" json:"employment_type"`
}