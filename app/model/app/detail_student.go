package app

import "time"

type DetailStudent struct {
	Id string `gorm:"type:uuid;not null;default:uuid_generate_v4();primaryKey"`
	UserId string
	User User
	Nisn string
	Nis string
	CreatedAt time.Time
}
