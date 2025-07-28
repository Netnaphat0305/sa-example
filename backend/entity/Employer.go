package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employer struct {
	gorm.Model
	Company_Name    string    `gorm:"type:varchar(255);not null" json:"company_name"`
	Contact_Person  string    `gorm:"type:varchar(255);not null" json:"contact_person"`
	Post_Date       time.Time `gorm:"type:date;not null" json:"post_date"`
	Phone			string	  `gorm:"type:varchar(10);not null" json:"phone"`
	Address			string	  `gorm:"type:varchar(255);not null" json:"address"`
}
