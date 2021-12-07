package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type mapel struct{}

func NewMapel() Mapel {
	return &mapel{}
}

func (m *mapel) Get(db *gorm.DB) ([]response.Mapel,error) {
	var mapel []response.Mapel
	result := db.Raw("SELECT id,subject,(SELECT COUNT(*) FROM teacher_subjects WHERE subject_id = subjects.id) AS jumlah_guru FROM subjects").Scan(&mapel)
	return mapel,result.Error
}

func (m *mapel) GetById(db *gorm.DB, mapelId string) (response.Mapel, error) {
	var mapel response.Mapel
	result := db.Where("id = ?",mapelId).First(&app.Subject{}).Scan(&mapel)
	return mapel,result.Error
}

func (m *mapel) GetPengampu(db *gorm.DB, mapelId string) ([]response.User, error) {
	var teacher []response.User
	result := db.Select("users.id,users.nama_depan,users.nama_belakang").Where("teacher_subjects.subject_id = ?",mapelId).Joins("INNER JOIN users ON users.id = teacher_subjects.user_id").Find(&[]app.TeacherSubject{}).Scan(&teacher)
	return teacher,result.Error
}

func (m *mapel) GetTeacher(db *gorm.DB, mapelId string) ([]response.User, error) {
	var teacher []response.User
	var level app.Level
	db.Where("level = ?", "guru").First(&level)
	result := db.Select("users.id,users.nama_depan,users.nama_belakang").Where("teacher_subjects.subject_id != ? OR teacher_subjects.subject_id IS NULL AND users.level_id = ?",mapelId,level.Id).Joins("LEFT JOIN teacher_subjects ON teacher_subjects.user_id = users.id ").Find(&[]app.User{}).Scan(&teacher)
	return teacher,result.Error
}

func (m *mapel) Post(db *gorm.DB, request request.Mapel) (bool, error) {
	result := db.Create(&app.Subject{
		Subject:   request.Mapel,
	})
	return result.RowsAffected > 0, result.Error
}

func (m *mapel) PostPengampu(db *gorm.DB, mapelId string, userId string) (bool, error) {
	result := db.Create(&app.TeacherSubject{
		UserId:    userId,
		SubjectId: mapelId,
	})
	return result.RowsAffected > 0, result.Error
}

func (m *mapel) Delete(db *gorm.DB, mapelId string) (bool, error) {
	result := db.Where("id = ?",mapelId).Delete(&app.Subject{})
	return result.RowsAffected > 0,result.Error
}

func (m *mapel) DeletePengampu(db *gorm.DB, mapelId string, userId string) (bool, error) {
	result := db.Where("subject_id = ? AND user_id = ?",mapelId,userId).Delete(&app.TeacherSubject{})
	return result.RowsAffected > 0, result.Error
}

func (m *mapel) Update(db *gorm.DB, mapelId string, request request.Mapel) (bool, error) {
	result := db.Where("id = ?",mapelId).Updates(&app.Subject{
		Subject:   request.Mapel,
	})
	return result.RowsAffected > 0, result.Error
}
