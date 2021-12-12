package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type Spp interface{
	GetStudents(db *gorm.DB) ([]response.SppStudents,error)
	GetByUserId(db *gorm.DB,userId string) (app.MonthlyFee,error)
	GetDetailSpp(db *gorm.DB, userId string) ([]response.DetailSpp,error)
	PostSpp(db *gorm.DB,request request.Spp) (bool,error)
}