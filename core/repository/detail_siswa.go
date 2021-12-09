package repository

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type DetailSiswa interface {
	GetByUserId(db *gorm.DB,userId string) (response.DetailSiswa,error)
	Post(db *gorm.DB, userId string, request request.DetailSiswa) (bool,error)
	Update(db *gorm.DB, userId string,request request.DetailSiswa) (bool,error)
}
