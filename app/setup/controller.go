package setup

import (
	"github.com/go-playground/validator/v10"
	"github.com/rizface/sekolah/core/controller"
	"github.com/rizface/sekolah/core/repository"
	service2 "github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
)

var valid = validator.New()

func LoginController() controller.Login {
	repo := repository.NewLogin()
	service := service2.NewLogin(valid,repo,helper.Connection())
	controller := controller.NewLogin(service)
	return controller
}

func CrudAdminController() controller.AdminCrudAdmin {
	repo := repository.NewAdmin()
	service := service2.NewAdmin(helper.Connection(),valid,repo)
	controller := controller.NewAdmin(service)
	return controller
}
