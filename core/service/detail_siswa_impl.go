package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
)

type detailSiswa struct {
	db *gorm.DB
	valid *validator.Validate
	repo repository.DetailSiswa
}

func NewDetailSiswa(db *gorm.DB, valid *validator.Validate, repo repository.DetailSiswa) DetailSiswa {
	return &detailSiswa{
		db:    db,
		valid: valid,
		repo:  repo,
	}
}

func (d *detailSiswa) Update(userId string,request request.DetailSiswa) string {
	var success bool
	var err error
	err = d.valid.Struct(request)
	helper.PanicIfError(err)
	_,err = d.repo.GetByUserId(d.db,userId)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			success,err = d.repo.Post(d.db,userId,request)
			helper.PanicIfError(err)
		} else {
			panic(errors.New("proses gagal"))
		}
	} else {
		success,err = d.repo.Update(d.db,userId,request)
		helper.PanicIfError(err)
	}
	if success {
		return "detail siswa berhasil diupdate"
	}
	return "detail siswa gagal diupdate"
}



