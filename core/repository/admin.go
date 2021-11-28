package repository

import (
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/app/request"
	"gorm.io/gorm"
)

type AdminCrudAdmin interface {
	Get(db *gorm.DB) ([]response.Admin,error)
	GetById(db *gorm.DB,adminId interface{}) (response.Admin,error)
	Post(request request.Admin, db *gorm.DB) (bool,error)
	Delete(db *gorm.DB,adminId interface{}) (bool,error)
	Update(db *gorm.DB, admin response.Admin, request request.Admin) bool
}
