package app

import "time"

type AbsenStudent struct {
	Id string `gorm:"type:uuid;not null;default:uuid_generate_v4();primaryKey"`
	SubjectId string
	Subject Subject `gorm:"constraint:onDelete:CASCADE"`
	StudentId string
	Student User `gorm:"constraint:onDelete:CASCADE"`
	Status string
	CreatedAt time.Time
}
