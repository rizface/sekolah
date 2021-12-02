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

type kelas struct {
	db *gorm.DB
	validate *validator.Validate
	repo repository.Kelas
}

func NewKelas(db *gorm.DB, validate *validator.Validate, repo repository.Kelas) Kelas{
	return &kelas{
		db:       db,
		validate: validate,
		repo:     repo,
	}
}

func (k *kelas) Get() []response.Kelas {
	classes,err := k.repo.Get(k.db)
	helper.PanicIfError(err)
	return classes
}

func (k *kelas) GetById(kelasId string) response.Kelas {
	class,err := k.repo.GetById(kelasId,k.db)
	if err != nil {
		fmt.Println(err.Error())
		if errors.Is(err,gorm.ErrRecordNotFound) {
			panic(errors.New("kelas tidak ditemukan"))
		}
		panic(errors.New("proses gagal"))
	}
	return class
}

func (k *kelas) Post(request request.Kelas) string {
	err := k.validate.Struct(request)
	helper.PanicIfError(err)
	success := k.repo.Post(request,k.db)
	if success {
		return fmt.Sprintf("%s %s - %s %s","kelas",request.Tingkat,request.Kelas,"berhasil ditambahkan")
	}
	return fmt.Sprintf("%s %s - %s %s","kelas",request.Tingkat,request.Kelas,"gagal ditambahkan")
}

func (k *kelas) Update(kelasId string, request request.Kelas) string {
	class := k.GetById(kelasId)
	k.repo.Update(class,request,k.db)
	return "update kelas berhasil dilakukan"
}

func (k *kelas) Delete(kelasId string) string {
	success := k.repo.Delete(kelasId, k.db)
	if success {
		return "kelas berhasil dihapus"
	}
	return "kelas gagal dihapus"
}

