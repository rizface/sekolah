package controller

import (
	"github.com/gorilla/mux"
	request2 "github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

type detailPegawai struct {
	service service.DetailPegawai
}

func NewDetailPegawai(service service.DetailPegawai) DetailPegawai {
	return &detailPegawai{
		service: service,
	}
}

func (d *detailPegawai) Post(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := request2.DetailPegawai{Nip: r.PostFormValue("nip")}
	result := d.service.Post(params["userId"],request)
	helper.Notif("Notification",result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}
