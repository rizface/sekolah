package repository

import (
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type GuruMapel interface{
	GetMapel(db *gorm.DB, guruId interface{}) ([]response.Mapel,error)
}