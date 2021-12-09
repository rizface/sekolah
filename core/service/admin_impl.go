package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
)

type admin struct{
	db *gorm.DB
	validate *validator.Validate
	repo repository.AdminCrud
}

func NewAdmin(db *gorm.DB, validate *validator.Validate,repo repository.AdminCrud) AdminCrud {
	return &admin{
		db:   db,
		validate: validate,
		repo: repo,
	}
}

func (a *admin) Get(level string) []response.User {
	result,err := a.repo.Get(a.db,level)
	helper.PanicIfError(err)
	return result
}

func (a *admin) Post(request request.User, level string) string {
	err := a.validate.Struct(request)
	helper.PanicIfError(err)

	request.Password = helper.GeneratePassword(request.Password)
	success,err := a.repo.Post(request,level,a.db)
	helper.PanicIfError(err)

	if success {
		return "data dengan nama " + request.NamaDepan + " " + request.NamaBelakang + " berhasil ditambahkan"
	}
	return "data gagal ditambahkan"
}

func (a *admin) Delete(adminId interface{}) string {
	success,err := a.repo.Delete(a.db,adminId)
	helper.PanicIfError(err)
	if success {
		return "data berhasl dihapus"
	}
	return "data gagal dihapus"
}

func (a *admin) GetById(userId interface{}) response.User {
	admin,err := a.repo.GetById(a.db,userId)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			panic(errors.New("data tidak ditemukan"))
		} else {
			panic(errors.New("proses gagal"))
		}
	}
	return admin
}

func checkErrorType(err error ) {
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) == false {
			helper.PanicIfError(err)
		}
	}
}

func (a *admin) DetailById(level string, userId string) interface{} {
	var result interface{}
	var err error
	if level == "murid" {
		result,err = a.repo.StudentDetail(a.db,userId)
		checkErrorType(err)
	} else {
		result,err = a.repo.EmployeeDetail(a.db,userId)
		checkErrorType(err)
	}
	return result
}

func (a *admin) Update(adminId interface{}, request request.User) string {
	admin := a.GetById(adminId)
	if len(request.Password) > 0 {
		request.Password = helper.GeneratePassword(request.Password)
	}
	success := a.repo.Update(a.db,admin,request)
	if success {
		return "update data berhasil"
	}
	return "update data gagal"
}
