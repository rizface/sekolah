package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	request2 "github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

type kelas struct {
	service service.Kelas
}

func NewKelas(service service.Kelas) AdminCrud {
	return &kelas{
		service: service,
	}
}

func (k *kelas) Get(w http.ResponseWriter, r *http.Request) {
	classes := k.service.Get()
	tmp := helper.View("view/admin/kelas/index.gohtml")
	err := tmp.Exec(w, r, "index_data_kelas", map[string]interface{}{
		"classes": classes,
	})
	helper.PanicIfError(err)
}

func (k *kelas) PostPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("masuk")
	tmp := helper.View("view/admin/kelas/tambah.gohtml")
	err := tmp.Exec(w, r, "tambah_data_kelas", nil)
	helper.PanicIfError(err)
}

func (k *kelas) Post(w http.ResponseWriter, r *http.Request) {
	request := request2.Kelas{
		Tingkat: r.PostFormValue("tingkat"),
		Kelas:   r.PostFormValue("kelas"),
	}
	result := k.service.Post(request)
	helper.Notif("Notification", result)
	http.Redirect(w, r, "/admin/data-kelas", http.StatusSeeOther)
}

func (k *kelas) UpdatePage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	class := k.service.GetById(params["kelasId"])
	tmp := helper.View("view/admin/kelas/update.gohtml")
	err := tmp.Exec(w,r,"update_data_kelas", map[string]interface{} {
		"class":class,
	})
	helper.PanicIfError(err)
}

func (k *kelas) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := request2.Kelas{
		Tingkat: r.PostFormValue("tingkat"),
		Kelas:   r.PostFormValue("kelas"),
	}
	result := k.service.Update(params["kelasId"],request)
	helper.Notif("Notification",result)
	http.Redirect(w,r,"/admin/data-kelas",http.StatusSeeOther)
}

func (k *kelas) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := k.service.Delete(params["kelasId"])
	helper.Notif("Nitification",result)
	http.Redirect(w,r,"/admin/data-kelas",http.StatusSeeOther)
}
