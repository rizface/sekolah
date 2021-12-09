package repository

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type AdminCrud interface {
	Get(db *gorm.DB, levelName string) ([]response.User,error)
	GetById(db *gorm.DB,adminId interface{}) (response.User,error)
	StudentDetail(db *gorm.DB,userId string) (response.DetailSiswa,error)
	EmployeeDetail(db *gorm.DB,userId string) (response.DetailPegawai,error)
	TeacherDetail(userId string)
	Post(request request.User, levelName string,db *gorm.DB) (bool,error)
	Delete(db *gorm.DB,adminId interface{}) (bool,error)
	Update(db *gorm.DB, admin response.User, request request.User) bool
}