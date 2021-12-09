package app

import "time"

type DetailStudent struct {
	Id string `gorm:"type:uuid;not null;default:uuid_generate_v4();primaryKey"`
	UserId string
	User User `gorm:"constraint:onDelete:CASCADE"`
	Nisn string `gorm:"index:idx_nisn,unique"`
	Nis string `gorm:"index:idx_nis,unique"`
	CreatedAt time.Time
}
