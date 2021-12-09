package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type detailPegawai struct{}

func NewDetailPegawai() DetailPegawai {
	return detailPegawai{}
}

func (d detailPegawai) GetByUserId(db *gorm.DB, userId string) (response.DetailPegawai,error) {
	var detail response.DetailPegawai
	result := db.Where("user_id = ?",userId).First(&app.DetailEmployee{}).Scan(&detail)
	return detail,result.Error
}

func (d detailPegawai) Post(db *gorm.DB, userId string, request request.DetailPegawai) (bool, error) {
	result := db.Create(&app.DetailEmployee{
		UserId:    userId,
		Nip:       request.Nip,
	})
	return result.RowsAffected > 0, result.Error
}

func (d detailPegawai) Update(db *gorm.DB, userId string, request request.DetailPegawai) (bool, error) {
	result := db.Where("user_id = ?",userId).Updates(&app.DetailEmployee{
		Nip:       request.Nip,
	})
	return result.RowsAffected > 0, result.Error
}

