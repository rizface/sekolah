package app

type TeacherSubject struct {
	Id        string `gorm:"type:uuid;not null;default:uuid_generate_v4();primaryKey"`
	UserId    string
	User      User `gorm:"constraint:onDelete:CASCADE"`
	SubjectId string
	Subject   Subject `gorm:"constraint:onDelete:CASCADE"`
}
