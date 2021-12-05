package controller

import (
	"github.com/gorilla/mux"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

type waliKelas struct {
	service service.WaliKelas
}

func NewWaliKelas(service service.WaliKelas) WaliKelas {
	return waliKelas{
		service: service,
	}
}

func (w2 waliKelas) GetWaliKelas(w http.ResponseWriter, r *http.Request) {
	tmp := helper.View("view/admin/kelas/wali_kelas.gohtml")
	withClass, wihtoutClass, class := w2.service.GetWaliKelas()
	err := tmp.Exec(w, r, "data_wali_kelas", map[string]interface{}{
		"withClass":    withClass,
		"withoutClass": wihtoutClass,
		"class":        class,
	})
	helper.PanicIfError(err)
}

func (w2 waliKelas) PostWaliKelas(w http.ResponseWriter, r *http.Request) {
	user := r.PostFormValue("idGuru")
	kelas := r.PostFormValue("idKelas")
	result := w2.service.PostKelas(user,kelas)
	helper.Notif("Nitification", result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}

func (w2 waliKelas) DeleteWaliKelas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	kelasId := params["kelasId"]
	result := w2.service.DeleteWaliKelas(userId,kelasId)
	helper.Notif("Notification",result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}

