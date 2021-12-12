package repository

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type NilaiSiswa interface {
	Post(db *gorm.DB,request request.Nilai) (bool,error)
	GetAllGrade(db *gorm.DB,userId string)([]response.Nilai,error)
	GetGradeBySemester(db *gorm.DB, userId string, semester string)([]response.Nilai,error)
	GetGradeBySubject(db *gorm.DB, userId string, subject string) ([]response.Nilai,error)
	GetGradeSubjectSemester(db *gorm.DB, userId string, semester string, subjectId string) ([]response.Nilai,error)
	GetGradeById(db *gorm.DB, gradeId string) (response.Nilai,error)
	Delete(db *gorm.DB, gradeId string) (bool,error)
	Update(db *gorm.DB,gradeId string, request request.Nilai) (bool,error)
}
