package controller

import (
	"github.com/rizface/sekolah/helper"
	"net/http"
)

type dashboard struct{}

func NewDashboard() Dashboard {
	return dashboard{}
}

func (d dashboard) DashboardAdmin(w http.ResponseWriter, r *http.Request) {
	tmp := helper.View("view/admin/dashboard.gohtml")
	err := tmp.Exec(w,r,"admin",nil)
	helper.PanicIfError(err)
}
