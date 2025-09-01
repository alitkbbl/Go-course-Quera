package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title       string `gorm:"size:100;not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Level       string `gorm:"size:50;default:'mid'" json:"level"`
	SalaryRange string `gorm:"size:100" json:"salary_range"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`
	Users       []User `gorm:"foreignKey:JobID" json:"users,omitempty"`
}

type User struct {
	gorm.Model
	FirstName string    `gorm:"size:50;not null" json:"first_name"`
	LastName  string    `gorm:"size:50;not null" json:"last_name"`
	Email     string    `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Age       int       `gorm:"check:age > 0" json:"age"`
	Salary    float64   `gorm:"default:0" json:"salary"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	HireDate  time.Time `json:"hire_date"`
	JobID     uint      `json:"job_id"`
	Job       Job       `gorm:"foreignKey:JobID" json:"job"`
}
