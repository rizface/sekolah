package repository

import (
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type WaliKelas interface {
	GetWalasWithClass(db *gorm.DB) ([]response.Walas,error)
	GetWalasWithoutClass(db *gorm.DB) ([]response.User,error)
	PostWalas(db *gorm.DB,userId string, kelasId string) (bool,error)
	DeleteWaliKelas(db *gorm.DB, userId string,kelasId string) (bool,error)
}
