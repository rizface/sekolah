package app

import "time"

type UserClass struct {
	Id string `gorm:"type:uuid;not null;default:uuid_generate_v4()"`
	UserId string
	User User `gorm:"constraint:OnDelete:CASCADE"`
	ClassId string
	Class Class `gorm:"constraint:OnUpdate:SET NULL"`
	CreatedAt time.Time
}