package repository

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type Kelas interface {
	Get(db *gorm.DB) ([]response.Kelas,error)
	GetById(kelasId string, db *gorm.DB) (response.Kelas,error)
	Post(request request.Kelas, db *gorm.DB) bool
	Update(kelas response.Kelas, request request.Kelas, db *gorm.DB) error
	Delete(kelasId string, db *gorm.DB) bool
}
