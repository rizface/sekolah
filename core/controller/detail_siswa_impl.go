package controller

import (
	"github.com/gorilla/mux"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

type detailSiswa struct {
	service service.DetailSiswa
}

func NewDetailSiswa(service service.DetailSiswa) DetailSiswa {
	return &detailSiswa{
		service: service,
	}
}

func (d *detailSiswa) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := request.DetailSiswa{
		Nisn: r.PostFormValue("nisn"),
		Nis:  r.PostFormValue("nis"),
	}
	result := d.service.Update(params["userId"],request)
	helper.Notif("Notification", result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}

