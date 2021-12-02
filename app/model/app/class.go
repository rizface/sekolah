package app

import "time"

type Class struct {
	Id        string `gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()"`
	Tingkat   string `gorm:"not null"`
	Kelas     string `gorm:"index:idx_kelas,unique"`
	CreatedAt time.Time
}
