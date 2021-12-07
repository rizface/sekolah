package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	request2 "github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

type mapel struct {
	service service.Mapel
}

func NewMapel(service service.Mapel) AdminCrud {
	return &mapel{
		service: service,
	}
}

func NewPengampu(service service.Mapel) Mapel {
	return &mapel{
		service: service,
	}
}

func (m *mapel) Get(w http.ResponseWriter, r *http.Request) {
	mapel := m.service.Get()
	fmt.Println(mapel)
	tmp := helper.View("view/admin/mapel/index.gohtml")
	err := tmp.Exec(w, r, "data_mapel_index", map[string]interface{}{
		"mapel": mapel,
	})
	helper.PanicIfError(err)
}

func (m *mapel) PostPage(w http.ResponseWriter, r *http.Request) {
	tmp := helper.View("view/admin/mapel/tambah.gohtml")
	err := tmp.Exec(w, r, "tambah_mapel", nil)
	helper.PanicIfError(err)
}

func (m *mapel) Post(w http.ResponseWriter, r *http.Request) {
	request := request2.Mapel{Mapel: r.PostFormValue("mapel")}
	result := m.service.Post(request)
	helper.Notif("Notification", result)
	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

func (m *mapel) UpdatePage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mapel := m.service.GetById(params["mapelId"])
	tmp := helper.View("view/admin/mapel/update.gohtml")
	err := tmp.Exec(w, r, "update_mapel", map[string]interface{}{
		"mapel": mapel,
	})
	helper.PanicIfError(err)
}

func (m *mapel) Update(w http.ResponseWriter, r *http.Request) {
	request := request2.Mapel{
		Mapel: r.PostFormValue("mapel"),
	}
	params := mux.Vars(r)
	mapelId := params["mapelId"]
	result := m.service.Update(mapelId, request)
	helper.Notif("Notification", result)
	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

func (m *mapel) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mapelId := params["mapelId"]
	result := m.service.Delete(mapelId)
	helper.Notif("Nitification", result)
	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

// Pengampu

func (m *mapel) GetPengampu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mapelId := params["mapelId"]
	pengampu, guru := m.service.GetPengampu(mapelId)
	tmp := helper.View("view/admin/mapel/pengampu.gohtml")
	err := tmp.Exec(w, r, "data_pengampu", map[string]interface{}{
		"pengampu": pengampu,
		"guru":     guru,
		"mapelId": mapelId,
	})
	helper.PanicIfError(err)
}

func (m *mapel) DeletePengampu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mapelId := params["mapelId"]
	userId := params["userId"]
	result := m.service.DeletePengampu(mapelId,userId)
	helper.Notif("Notification",result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}

func (m *mapel) PostPengampu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mapelId := params["mapelId"]
	userId := params["userId"]
	result := m.service.PostPengampu(mapelId,userId)
	helper.Notif("Notification",result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}


