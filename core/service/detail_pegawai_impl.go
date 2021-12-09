package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
)

type detailPegawai struct {
	db *gorm.DB
	valid *validator.Validate
	repo repository.DetailPegawai
}

func NewDetailPegawai(db *gorm.DB, valid *validator.Validate, repo repository.DetailPegawai) DetailPegawai {
	return &detailPegawai{
		db:    db,
		valid: valid,
		repo:  repo,
	}
}

func (d *detailPegawai) Post(userId string, request request.DetailPegawai) string {
	var success bool
	var err error
	_,err = d.repo.GetByUserId(d.db,userId)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			success,err = d.repo.Post(d.db,userId,request)
			helper.PanicIfError(err)
		} else {
			fmt.Println(err.Error())
			panic(errors.New("proses gagal"))
		}
	} else {
		success,err = d.repo.Update(d.db,userId,request)
		helper.PanicIfError(err)
	}
	if success {
		return "data detail pegawai berhasil diupdate"
	}
	return "data detail pegawai gagal diupdate"
}

