package app

import "time"

type TeacherClass struct {
	Id string `gorm:"type:uuid;not null;default:uuid_generate_v4()"`
	UserId string
	User User `gorm:"constraint:OnDelete:CASCADE"`
	ClassId string
	Class Class `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
}
