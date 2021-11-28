package app

type Gender struct {
	Id string `gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()"`
	Gender string
}
