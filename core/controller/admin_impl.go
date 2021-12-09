package controller

import (
	"github.com/gorilla/mux"
	request2 "github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

type admin struct {
	service service.AdminCrud
}

func NewAdmin(service service.AdminCrud) AdminCrud {
	return &admin{
		service: service,
	}
}

func (a *admin) Get(w http.ResponseWriter, r *http.Request) {
	tmp := helper.View(helper.GetHeader(r, "X-INDEX"))
	user := a.service.Get(helper.GetHeader(r, "X-ROLE"))
	err := tmp.Exec(w, r, helper.GetHeader(r, "X-INDEX-VIEW-NAME"), map[string]interface{}{
		"data": user,
	})
	helper.PanicIfError(err)
}

func (a *admin) PostPage(w http.ResponseWriter, r *http.Request) {
	tmp := helper.View(helper.GetHeader(r, "X-POST-PAGE"))
	err := tmp.Exec(w, r, helper.GetHeader(r, "X-POST-PAGE-VIEW-NAME"), nil)
	helper.PanicIfError(err)
}

func (a *admin) Post(w http.ResponseWriter, r *http.Request) {
	request := request2.User{
		NamaDepan:    r.PostFormValue("nama_depan"),
		NamaBelakang: r.PostFormValue("nama_belakang"),
		Username:     r.PostFormValue("username"),
		Password:     r.PostFormValue("password"),
	}
	result := a.service.Post(request, helper.GetHeader(r, "X-ROLE"))
	helper.Notif("Notification", result)
	http.Redirect(w, r, helper.GetHeader(r, "X-INDEX-PATH"), http.StatusSeeOther)
}

func (a *admin) UpdatePage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := a.service.GetById(params["userId"])
	detail := a.service.DetailById(helper.GetHeader(r, "X-ROLE"), params["userId"])
	tmp := helper.View(helper.GetHeader(r, "X-UPDATE-PAGE"))
	err := tmp.Exec(w, r, helper.GetHeader(r, "X-UPDATE-PAGE-VIEW-NAME"), map[string]interface{}{
		"data":   user,
		"detail": detail,
	})
	helper.PanicIfError(err)
}

func (a *admin) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := request2.User{
		NamaDepan:    r.PostFormValue("nama_depan"),
		NamaBelakang: r.PostFormValue("nama_belakang"),
		Username:     r.PostFormValue("username"),
		Password:     r.PostFormValue("password"),
	}
	userId := params["userId"]
	result := a.service.Update(userId, request)
	helper.Notif("Notification", result)
	redirect := helper.GetHeader(r, "X-UPDATE-REDIRECT") + userId
	http.Redirect(w, r, redirect, http.StatusSeeOther)
}

func (a *admin) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	result := a.service.Delete(userId)
	helper.Notif("Notification", result)
	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}
