package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
)

type spp struct {
	db *gorm.DB
	valid *validator.Validate
	repo repository.Spp
}

func NewSpp(db *gorm.DB, valid *validator.Validate, repo repository.Spp) Spp {
	return &spp{
		db:    db,
		valid: valid,
		repo:  repo,
	}
}

func (s *spp) GetStudents() []response.SppStudents {
	result,err := s.repo.GetStudents(s.db)
	helper.PanicIfError(err)
	return result
}

func (s *spp) GetDetailSpp(userId string) []response.DetailSpp {
	result,err := s.repo.GetDetailSpp(s.db,userId)
	helper.PanicIfError(err)
	return result
}

func (s *spp) PostSpp(request request.Spp) string {
	var success bool
	var err error
	err = s.valid.Struct(request)
	helper.PanicIfError(err)
	student,err := s.repo.GetByUserId(s.db,request.StudentId)
	fmt.Println(err)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			success,err = s.repo.PostSpp(s.db,request)
		} else {
			fmt.Println(err.Error())
			panic(errors.New("proses gagal"))
		}
	} else {
		fmt.Println("student : ",student.CreatedAt)
		request.CreatedAt = student.CreatedAt.AddDate(0,1,0)
		fmt.Println("masuk sini : ", request)
		success,err = s.repo.PostSpp(helper.Connection(),request)
		helper.PanicIfError(err)
	}
	if success {
		return "pembayaran spp berhasil"
	}
	return "pembayaran spp gagal"
}

