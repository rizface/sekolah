package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/app/request"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
)

type admin struct{
	db *gorm.DB
	validate *validator.Validate
	repo repository.AdminCrudAdmin
}

func NewAdmin(db *gorm.DB, validate *validator.Validate,repo repository.AdminCrudAdmin) AdminCrudAdmin {
	return admin{
		db:   db,
		validate: validate,
		repo: repo,
	}
}

func (a admin) Get() []response.Admin {
	result,err := a.repo.Get(a.db)
	helper.PanicIfError(err)
	return result
}

func (a admin) Post(request request.Admin) string {
	err := a.validate.Struct(request)
	helper.PanicIfError(err)

	request.Password = helper.GeneratePassword(request.Password)
	success,err := a.repo.Post(request,a.db)
	helper.PanicIfError(err)

	if success {
		return "admin dengan nama " + request.NamaDepan + " " + request.NamaBelakang + " berhasil ditambahkan"
	}
	return "admin gagal ditambahkan"
}

func (a admin) Delete(adminId interface{}) string {
	success,err := a.repo.Delete(a.db,adminId)
	helper.PanicIfError(err)
	if success {
		return "data admin berhasl dihapus"
	}
	return "data admin gagal dihapus"
}

func (a admin) GetById(adminId interface{}) response.Admin {
	admin,err := a.repo.GetById(a.db,adminId)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			panic(errors.New("data tidak ditemukan"))
		} else {
			panic(errors.New("proses gagal"))
		}
	}
	return admin
}

func (a admin) Update(adminId interface{}, request request.Admin) string {
	admin := a.GetById(adminId)
	if len(request.Password) > 0 {
		request.Password = helper.GeneratePassword(request.Password)
	}
	success := a.repo.Update(a.db,admin,request)
	if success {
		return "update data admin berhasil"
	}
	return "update data admin gagal"
}
