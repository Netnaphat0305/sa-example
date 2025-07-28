package entity

import "gorm.io/gorm"

type EmploymentType struct {
	gorm.Model
	EmploymentTypeName string `gorm:"type:varchar(20);not null;unique" json:"employment_type_name"`
}
