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
	nilaiService := service2.NewNilaiSiswa(helper.Connection(),valid,repository.NewSemester(),repository.NewMapel(),repository.NewNilaiSiswa())
	service := service2.NewKelas(helper.Connection(),valid,repo)
	controller := controller.NewAssignStudent(service,nilaiService)
	return controller
}

func CrudWalas() controller.WaliKelas {
	walasRepo := repository.NewWaliKelas()
	kelasRepo := repository.NewKelas()
	service := service2.NewWaliKelas(helper.Connection(),kelasRepo,walasRepo)
	controller := controller.NewWaliKelas(service)
	return controller
}

func CrudMapel() controller.AdminCrud {
	repo := repository.NewMapel()
	service := service2.NewMapel(helper.Connection(),valid,repo)
	controller := controller.NewMapel(service)
	return controller
}

func CrudPengampu() controller.Mapel {
	repo := repository.NewMapel()
	service := service2.NewMapel(helper.Connection(),valid,repo)
	controller := controller.NewPengampu(service)
	return controller
}

func CrudNilaisiswa() controller.NilaiSiswa {
	repo := repository.NewNilaiSiswa()
	semesterRepo := repository.NewSemester()
	subjectRepo := repository.NewMapel()
	service := service2.NewNilaiSiswa(helper.Connection(),valid,semesterRepo,subjectRepo,repo)
	controller := controller.NewNilaiSiswa(service)
	return controller
}

func CrudDetailSiswa() controller.DetailSiswa {
	repo := repository.NewDetailSiswa()
	service := service2.NewDetailSiswa(helper.Connection(),valid,repo)
	controller := controller.NewDetailSiswa(service)
	return controller
}

//func CrudGuruController() controller.AdminCrud {
//	repo := repository.NewUser()
//	service := service2.NewCrudGuru(helper.Connection(),valid,repo)
//	controller := controller.NewCrudGuru(service)
//	return controller
//}