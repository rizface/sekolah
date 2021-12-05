package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type waliKelas struct{}

func NewWaliKelas() WaliKelas{
	return waliKelas{}
}

func (w waliKelas) GetWalasWithClass(db *gorm.DB) ([]response.Walas,error) {
	var level app.Level
	var walas []response.Walas
	db.Where("level = ?","guru").First(&level)
	result := db.Table("teacher_classes").Joins("INNER JOIN users ON users.id = teacher_classes.user_id INNER JOIN classes ON classes.id = teacher_classes.class_id").Select("users.id,users.nama_depan,users.nama_belakang,classes.id AS kelas_id,classes.tingkat,classes,kelas").Where("users.level_id = ?",level.Id).Scan(&walas)
	return walas,result.Error
}

func (w waliKelas) GetWalasWithoutClass(db *gorm.DB) ([]response.User,error) {
	var level app.Level
	var teacher []response.User
	db.Where("level = ?","guru").First(&level)
	result := db.Where("level_id = ? AND teacher_classes.class_id IS NULL",level.Id).Select("users.id,users.nama_depan,users.nama_belakang").Joins("LEFT JOIN teacher_classes ON teacher_classes.user_id = users.id").Find(&teacher)
	return teacher,result.Error
}

func (w waliKelas) PostWalas(db *gorm.DB, userId string, kelasId string) (bool, error) {
	result := db.Create(&app.TeacherClass{
		UserId:    userId,
		ClassId:   kelasId,
	})
	return result.RowsAffected > 0, result.Error
}

func (w waliKelas) DeleteWaliKelas(db *gorm.DB, userId string,kelasId string) (bool, error) {
	result := db.Where("class_id = ? AND user_id = ?", kelasId,userId).Delete(&app.TeacherClass{})
	return result.RowsAffected > 0, result.Error
}
