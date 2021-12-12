package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

func NewAssignStudent(kelas2 service.Kelas, nilaiService service.NilaiSiswa, userService service.AdminCrud) Kelas {
	return &kelas{
		service:       kelas2,
		nilaiService:  nilaiService,
		userService:   userService,
	}
}

func (k *kelas) GetStudents(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	students, withoutClass := k.service.GetStudent(params["kelasId"])
	tmp := helper.View("view/admin/kelas/siswa.gohtml")
	err := tmp.Exec(w, r, "data_siswa_kelas", map[string]interface{}{
		"data":    students,
		"noclass": withoutClass,
		"kelasId": params["kelasId"],
	})
	helper.PanicIfError(err)
}

func (k *kelas) AddStudents(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kelasId := params["kelasId"]
	userId := params["userId"]
	result := k.service.AddStudent(kelasId, userId)
	helper.Notif("Notification", result)
	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

func (k *kelas) DeleteStudents(W http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kelasId := params["kelasId"]
	userId := params["userId"]
	result := k.service.DeleteStudent(kelasId, userId)
	helper.Notif("Notication", result)
	http.Redirect(W, r, r.Referer(), http.StatusSeeOther)
}

func (k *kelas) DetailStudents(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	semester := r.URL.Query().Get("semester")
	grade := k.nilaiService.GetNilaiBySemester(userId, semester)
	data := k.userService.GetById(userId)
	detail := k.userService.DetailById("murid", userId)
	fmt.Println(detail)
	tmp := helper.View("view/admin/siswa/detail_siswa.gohtml")
	err := tmp.Exec(w, r, "detail_siswa", map[string]interface{}{
		"grades": grade,
		"data":   data,
		"detail": detail,
	})
	helper.PanicIfError(err)
}