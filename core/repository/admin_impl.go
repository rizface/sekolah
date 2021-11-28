package repository

import (
	"fmt"
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/app/request"
	"gorm.io/gorm"
)

type admin struct{}

func NewAdmin() AdminCrudAdmin{
	return admin{}
}

func (a admin) Get(db *gorm.DB) ([]response.Admin,error) {
	var result []response.Admin
	db.Find(&[]app.User{}).Scan(&result)
	return result,db.Error
}

func (a admin) Post(request request.Admin, db *gorm.DB) (bool,error) {
	var level app.Level
	var gender app.Gender
	result := db.Where("level = ?", "admin").First(&level)
	fmt.Println(result.Error)
	result = db.Where("gender = ?", "laki-laki").First(&gender)
	fmt.Println(result.Error)
	result = db.Create(&app.User{
		NamaDepan:    request.NamaDepan,
		NamaBelakang: request.NamaBelakang,
		Username:     request.Username,
		Password:     request.Password,
		LevelId:      level.Id,
		GenderId:     gender.Id,
	})
	return result.RowsAffected > 0,result.Error
}

func (a admin) Delete(db *gorm.DB, adminId interface{}) (bool,error) {
	result := db.Where("id = ?",adminId).Delete(&app.User{})
	return result.RowsAffected > 0,result.Error
}

func (a admin) GetById(db *gorm.DB,adminId interface{}) (response.Admin,error) {
	var admin response.Admin
	result := db.Where("id = ?", adminId).First(&app.User{}).Scan(&admin)
	return admin, result.Error
}

func (a admin) Update(db *gorm.DB, admin response.Admin, request request.Admin) bool {
	result := db.Where("id = ?", admin.Id).Updates(&app.User{
		NamaDepan:    request.NamaDepan,
		NamaBelakang: request.NamaBelakang,
		Username:     request.Username,
		Password:     request.Password,
	})
	return result.RowsAffected > 0
}

