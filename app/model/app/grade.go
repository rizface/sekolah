package app

import "time"

type Grade struct {
	Id string `gorm:"type:uuid;not null;default:uuid_generate_v4();primaryKey;"`
	StudentId string
	Student User `gorm:"constraint:onDelete:CASCADE"`
	TeacherId string
	Teacher User `gorm:"constraint:onDelete:SET NULL"`
	SemesterId string
	Semester Semester
	SubjectId string
	Subject Subject `gorm:"constraint:onDelete:CASCADE"`
	Grade float64
	CreatedAt time.Time
}
