package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type detailSiswa struct{}

func NewDetailSiswa() DetailSiswa {
	return &detailSiswa{}
}

func (d *detailSiswa) GetByUserId(db *gorm.DB, userId string) (response.DetailSiswa, error) {
	var detail response.DetailSiswa
	result := db.Where("user_id = ?", userId).First(&app.DetailStudent{}).Scan(&detail)
	return detail,result.Error
}

func (d *detailSiswa) Update(db *gorm.DB, userId string, request request.DetailSiswa) (bool, error) {
	result := db.Where("user_id = ?", userId).Updates(&app.DetailStudent{
		Nisn:      request.Nisn,
		Nis:       request.Nis,
	})
	return result.RowsAffected > 0, result.Error
}

func (d *detailSiswa) Post(db *gorm.DB, userId string,request request.DetailSiswa) (bool, error) {
	result := db.Create(&app.DetailStudent{
		UserId:    userId,
		Nisn:      request.Nisn,
		Nis:       request.Nis,
	})
	return result.RowsAffected > 0, result.Error
}

