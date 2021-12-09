package repository

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type DetailPegawai interface {
	GetByUserId(db *gorm.DB, userId string) (response.DetailPegawai,error)
	Post(db *gorm.DB, userId string, request request.DetailPegawai) (bool,error)
	Update(db *gorm.DB, userId string, request request.DetailPegawai) (bool,error)
}
