package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/app/request"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
)

type login struct {
	validator *validator.Validate
	repo      repository.Login
	db        *gorm.DB
}

func NewLogin(validate *validator.Validate, repo repository.Login, db *gorm.DB) Login {
	return login{
		validator: validate,
		repo:      repo,
		db:        db,
	}
}

func (l login) Login(login request.Login) response.Login {
	err := l.validator.Struct(login)
	helper.PanicIfError(err)
	user,err := l.repo.FindByUsername(login.Username,l.db)
	fmt.Println(user)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			panic(errors.New("data tidak ditemukan"))
		}
	}
	err = helper.ComparePassword(login.Password,user.Password)
	helper.PanicIfError(err)


	return user
}
