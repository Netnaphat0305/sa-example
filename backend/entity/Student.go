package entity

import (
	"time"

	"gorm.io/gorm"
)

// GenderEnum กำหนดค่า enum สำหรับ Gender
type GenderEnum string

const (
	Male   GenderEnum = "Male"
	Female GenderEnum = "Female"
	Other  GenderEnum = "Other"
)

type Student struct {
	gorm.Model // ประกอบด้วย ID, CreatedAt, UpdatedAt, DeletedAt
	Password  string     `gorm:"type:varchar(255);not null"`
	Email     string     `gorm:"type:varchar(255);not null"`
	FirstName string     `gorm:"type:varchar(100);not null"`
	LastName  string     `gorm:"type:varchar(100);not null"`
	Birthday  time.Time  `gorm:"type:date;not null"`
	Gender    GenderEnum `gorm:"type:enum('Male','Female','Other');not null"`
	Age       int        `gorm:"not null"`
	GPA       float32    `gorm:"not null"`
	Year      int        `gorm:"not null"`
	Faculty   string     `gorm:"type:varchar(255);not null"`
	Phone     string     `gorm:"type:varchar(20);not null"`
}
