package entity

import "gorm.io/gorm"

type JobCategory struct {
	gorm.Model
	JobCategoryName	string		`gorm:"type:varchar(50);not null" json:"jobcategory_name"`
}