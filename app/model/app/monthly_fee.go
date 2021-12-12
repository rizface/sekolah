package app

import "time"

type MonthlyFee struct {
	Id        string `gorm:"type:uuid;not null;default:uuid_generate_v4();primaryKey"`
	StudentId string
	Student   User `gorm:"constraint:OnDelete:CASCADE"`
	OfficerId string
	Officer   User `gorm:"constraint:OnDelete:SET NULL"`
	Amount    int
	CreatedAt time.Time
}
