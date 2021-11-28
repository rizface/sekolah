package app

type Level struct {
	Id string `gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()"`
	Level string
}
