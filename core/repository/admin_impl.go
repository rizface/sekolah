package repository

import (
	"fmt"
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
)

type admin struct{}

func NewUser() AdminCrud {
	return &admin{}
}

func (a *admin) Get(db *gorm.DB, levelName string) ([]response.User,error) {
	fmt.Println(levelName, "ini level")
	var user []response.User
	var level app.Level
	result := db.Where("level = ?", levelName).First(&level)
	helper.PanicIfError(result.Error)
	result = db.Find(&[]app.User{}).Where("level_id = ?", level.Id).Scan(&user)
	return user,result.Error
}

func (a *admin) StudentDetail(db *gorm.DB,userId string) (response.DetailSiswa, error) {
	var detail response.DetailSiswa
	result := db.Where("user_id = ?",userId).First(&app.DetailStudent{}).Scan(&detail)
	return detail,result.Error
}

func (a *admin) EmployeeDetail(db *gorm.DB, userId string) (response.DetailPegawai, error) {
	var detail response.DetailPegawai
	result := db.Where("user_id = ?",userId).First(&app.DetailEmployee{}).Scan(&detail)
	return detail,result.Error
}

func (a *admin) TeacherDetail(userId string) {
	panic("implement me")
}

func (a *admin) Post(request request.User, levelName string,db *gorm.DB) (bool,error) {
	var level app.Level
	var gender app.Gender
	result := db.Where("level = ?", levelName).First(&level)
	result = db.Where("gender = ?", request.JenisKelamin).First(&gender)
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

func (a *admin) Delete(db *gorm.DB, adminId interface{}) (bool,error) {
	result := db.Where("id = ?",adminId).Delete(&app.User{})
	return result.RowsAffected > 0,result.Error
}

func (a *admin) GetById(db *gorm.DB,adminId interface{}) (response.User,error) {
	var admin response.User
	result := db.Where("id = ?", adminId).First(&app.User{}).Scan(&admin)
	return admin, result.Error
}

func (a *admin) Update(db *gorm.DB, admin response.User, request request.User) bool {
	result := db.Where("id = ?", admin.Id).Updates(&app.User{
		NamaDepan:    request.NamaDepan,
		NamaBelakang: request.NamaBelakang,
		Username:     request.Username,
		Password:     request.Password,
	})
	return result.RowsAffected > 0
}

