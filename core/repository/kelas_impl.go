package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type kelas struct{}

func NewKelas() Kelas {
	return &kelas{}
}

func (k *kelas) Get(db *gorm.DB) ([]response.Kelas, error) {
	var classes []response.Kelas
	result := db.Find(&[]app.Class{}).Scan(&classes)
	return classes, result.Error
}

func (k *kelas) GetById(kelasId string,db *gorm.DB) (response.Kelas,error) {
	var class response.Kelas
	result := db.Where("id = ?",kelasId).First(&app.Class{}).Scan(&class)
	return class,result.Error
}

func (k *kelas) Post(request request.Kelas, db *gorm.DB) bool {
	result := db.Create(&app.Class{
		Tingkat: request.Tingkat,
		Kelas:   request.Kelas,
	})
	return result.RowsAffected > 0
}

func (k *kelas) Update(kelas response.Kelas, request request.Kelas, db *gorm.DB) error {
	result := db.Where("id = ?",kelas.Id).Updates(&app.Class{
		Tingkat:   request.Tingkat,
		Kelas:     request.Kelas,
	})
	return result.Error
}

func (k *kelas) Delete(kelasId string, db *gorm.DB) bool {
	result := db.Where("id = ?",kelasId).Delete(&app.Class{})
	return result.RowsAffected > 0
}
