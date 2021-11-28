package app

import "time"

type User struct {
	Id           string `gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()"`
	NamaDepan    string `gorm:"check:nama_depan_checker, LENGTH(nama_depan) > 0"`
	NamaBelakang string
	Username     string `gorm:"index:idx_username,unique"`
	Password     string `gorm:"check:password_checker,LENGTH(password) > 5"`
	LevelId      string
	Level        Level `gorm:"constraint:onDelete:SET NULL"`
	GenderId     string
	Gender       Gender `gorm:"constraint:onDelete:SET NULL"`
	CreatedAt    time.Time
}
