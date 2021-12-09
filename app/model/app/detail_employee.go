package app

import "time"

type DetailEmployee struct {
	Id        string `gorm:"type:uuid;not null;default:uuid_generate_v4();primaryKey"`
	UserId    string
	User      User `gorm:"constraint:onDelete:CASCADE"`
	Nip       string `gorm:"index:idx_nip,unique"`
	CreatedAt time.Time
}
