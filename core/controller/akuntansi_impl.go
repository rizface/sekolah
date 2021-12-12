package controller

import (
	"github.com/gorilla/mux"
	request2 "github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

type akuntansi struct {
	userService service.AdminCrud
	sppService service.Spp
}

func NewAkuntansi(userService service.AdminCrud,sppService service.Spp) Akuntansi {
	return &akuntansi{
		userService: userService,
		sppService:  sppService,
	}
}

func (a *akuntansi) GetStudets(w http.ResponseWriter, r *http.Request) {
	students := a.sppService.GetStudents()
	tmp := helper.View("view/admin/akuntansi/siswa.gohtml")
	err := tmp.Exec(w,r,"spp_siswa",map[string]interface{} {
		"students": students,
	})
	helper.PanicIfError(err)
}

func (a *akuntansi) GetDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := a.sppService.GetDetailSpp(params["userId"])
	tmp := helper.View("view/admin/akuntansi/detail_spp.gohtml")
	err := tmp.Exec(w,r,"detail_spp_siswa",map[string]interface{} {
		"details":result,
	})
	helper.PanicIfError(err)
}

func (a *akuntansi) PostPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := a.userService.GetById(params["userId"])
	tmp := helper.View("view/admin/akuntansi/bayar_spp.gohtml")
	err := tmp.Exec(w,r,"bayar_spp_siswa",map[string]interface{} {
		"data":data,
	})
	helper.PanicIfError(err)
}

func (a *akuntansi) Post(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := request2.Spp{
		Nominal:    r.PostFormValue("nominal"),
		StudentId:  params["userId"],
		EmployeeId: r.PostFormValue("employee_id"),
	}
	result := a.sppService.PostSpp(request)
	helper.Notif("Notification", result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}
