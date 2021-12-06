package app

import "time"

type Subject struct {
	Id string `gorm:"type:uuid;not null;default:uuid_generate_v4();primaryKey"`
	Subject string `gorm:"index:idx_subject,unique"`
	CreatedAt time.Time
}
