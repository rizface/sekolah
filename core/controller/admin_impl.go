package controller

import (
	"github.com/gorilla/mux"
	request2 "github.com/rizface/sekolah/app/request"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

type admin struct {
	service service.AdminCrudAdmin
}

func NewAdmin(service service.AdminCrudAdmin) AdminCrudAdmin {
	return admin{
		service: service,
	}
}

func (a admin) Get(w http.ResponseWriter, r *http.Request) {
	tmp := helper.View("view/admin/admin/index.gohtml")
	admin := a.service.Get()
	err := tmp.ExecuteTemplate(w, "data_admin_index", map[string]interface{}{
		"admin": admin,
	})
	helper.PanicIfError(err)
}

func (a admin) PostPage(w http.ResponseWriter, r *http.Request) {
	tmp := helper.View("view/admin/admin/tambah.gohtml")
	err := tmp.ExecuteTemplate(w,"tambah_data_admin",nil)
	helper.PanicIfError(err)
}

func (a admin) Post(w http.ResponseWriter, r *http.Request) {
	request := request2.Admin{
		NamaDepan:    r.PostFormValue("nama_depan"),
		NamaBelakang: r.PostFormValue("nama_belakang"),
		Username:     r.PostFormValue("username"),
		Password:     r.PostFormValue("password"),
	}
	result := a.service.Post(request)
	helper.Notif("Notification", result)
	http.Redirect(w,r,"/admin/data-admin",http.StatusSeeOther)
}

func (a admin) UpdatePage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	admin := a.service.GetById(params["adminId"])
	tmp := helper.View("view/admin/admin/update.gohtml")
	tmp.ExecuteTemplate(w,"update_data_admin",map[string]interface{} {
		"admin": admin,
	})
}

func (a admin) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := request2.Admin{
		NamaDepan:    r.PostFormValue("nama_depan"),
		NamaBelakang: r.PostFormValue("nama_belakang"),
		Username:     r.PostFormValue("username"),
		Password:     r.PostFormValue("password"),
	}
	result := a.service.Update(params["adminId"],request)
	helper.Notif("Notification", result)
	http.Redirect(w,r,"/admin/data-admin/update/"+params["adminId"],http.StatusSeeOther)
}

func (a admin) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adminId := params["adminId"]
	result := a.service.Delete(adminId)
	helper.Notif("Notification",result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}
