package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employer struct {
	gorm.Model
	CompanyName   string    `gorm:"type:varchar(255);not null" json:"company_name"`
	ContactPerson string    `gorm:"type:varchar(255);not null" json:"contact_person"`
	PostDate      time.Time `gorm:"type:date;not null" json:"post_date"`
	Phone         string    `gorm:"type:varchar(10);not null" json:"phone"`
	Address       string    `gorm:"type:text;not null" json:"address"`
}
