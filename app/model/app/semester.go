package app

type Semester struct {
	Id string `gorm:"type:uuid;not null;default:uuid_generate_v4();primaryKey"`
	Semester string `gorm:"index:idx_semester,unique"`
}
