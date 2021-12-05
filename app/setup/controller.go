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

func CrudAdminController() controller.AdminCrud {
	repo := repository.NewUser()
	service := service2.NewAdmin(helper.Connection(),valid,repo)
	controller := controller.NewAdmin(service)
	return controller
}

func CrudKelasController() controller.AdminCrud {
	repo := repository.NewKelas()
	service := service2.NewKelas(helper.Connection(),valid,repo)
	controller := controller.NewKelas(service)
	return controller
}

func CrudAssignStudentToClass() controller.Kelas {
	repo := repository.NewKelas()
	service := service2.NewKelas(helper.Connection(),valid,repo)
	controller := controller.NewAssignStudent(service)
	return controller
}

func CrudWalas() controller.WaliKelas {
	walasRepo := repository.NewWaliKelas()
	kelasRepo := repository.NewKelas()
	service := service2.NewWaliKelas(helper.Connection(),kelasRepo,walasRepo)
	controller := controller.NewWaliKelas(service)
	return controller
}


//func CrudGuruController() controller.AdminCrud {
//	repo := repository.NewUser()
//	service := service2.NewCrudGuru(helper.Connection(),valid,repo)
//	controller := controller.NewCrudGuru(service)
//	return controller
//}