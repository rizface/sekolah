package repository

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type Mapel interface {
	Get(db *gorm.DB) ([]response.Mapel,error)
	GetById(db *gorm.DB, mapelId string) (response.Mapel,error)
	GetPengampu(db *gorm.DB, mapelId string) ([]response.User,error)
	Post(db *gorm.DB,request request.Mapel) (bool,error)
	Delete(db *gorm.DB, mapelId string) (bool,error)
	Update(db *gorm.DB, mapelId string, request request.Mapel) (bool,error)
}
