package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type kelas struct{}

func NewKelas() Kelas {
	return &kelas{}
}

func (k *kelas) Get(db *gorm.DB) ([]response.Kelas, error) {
	var classes []response.Kelas
	result := db.Find(&[]app.Class{}).Scan(&classes)
	return classes, result.Error
}

func (k *kelas) GetWithoutWalas(db *gorm.DB) ([]response.Kelas, error) {
	var classes []response.Kelas
	result := db.Where("teacher_classes.class_id IS NULL").Joins("LEFT JOIN teacher_classes ON teacher_classes.class_id = classes.id").Find(&[]app.Class{}).Scan(&classes)
	return classes,result.Error
}

func (k *kelas) GetById(kelasId string,db *gorm.DB) (response.Kelas,error) {
	var class response.Kelas
	result := db.Where("id = ?",kelasId).First(&app.Class{}).Scan(&class)
	return class,result.Error
}

func (k *kelas) GetStudents(kelasId string, db *gorm.DB) ([]response.User, error) {
	var students []response.User
	result := db.Table("user_classes").Where("class_id",kelasId).Joins("INNER JOIN users ON users.id = user_classes.user_id").Select("users.id,users.nama_depan,users.nama_belakang").Scan(&students)
	return students,result.Error
}

func (k *kelas) GetStudentsWithoutClass(db *gorm.DB) ([]response.User, error) {
	var level app.Level
	var students []response.User
	db.Where("level = ?", "murid").First(&level)
	result := db.Where("users.level_id = ? AND user_classes.class_id IS NULL",level.Id).Joins("LEFT JOIN user_classes ON user_classes.user_id = users.id").Select("users.id,users.nama_depan,users.nama_belakang").Find(&[]app.User{}).Scan(&students)
	return students,result.Error
}

func (k *kelas) Post(request request.Kelas, db *gorm.DB) bool {
	result := db.Create(&app.Class{
		Tingkat: request.Tingkat,
		Kelas:   request.Kelas,
	})
	return result.RowsAffected > 0
}

func (k *kelas) AddStudent(kelasId string, userId string, db *gorm.DB) (bool,error) {
	result := db.Create(&app.UserClass{
		UserId:    userId,
		ClassId:   kelasId,
	})
	return result.RowsAffected > 0 ,result.Error
}

func (k *kelas) Update(kelas response.Kelas, request request.Kelas, db *gorm.DB) error {
	result := db.Where("id = ?",kelas.Id).Updates(&app.Class{
		Tingkat:   request.Tingkat,
		Kelas:     request.Kelas,
	})
	return result.Error
}

func (k *kelas) Delete(kelasId string, db *gorm.DB) bool {
	result := db.Where("id = ?",kelasId).Delete(&app.Class{})
	return result.RowsAffected > 0
}

func (k *kelas) DeleteStudent(kelasId string, userId string, db *gorm.DB) (bool, error) {
	result := db.Where("class_id = ? AND user_id = ?", kelasId,userId).Delete(&app.UserClass{})
	return result.RowsAffected > 0, result.Error
}
