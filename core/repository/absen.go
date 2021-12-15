package repository

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type Absen interface {
	GetStudentAbsent(db *gorm.DB, siswaId string) ([]response.Absen,error)
	Absen(db *gorm.DB,request request.Absen) error
}
