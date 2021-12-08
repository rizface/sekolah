package controller

import (
	"github.com/gorilla/mux"
	request2 "github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
	"strconv"
)

type nilaiSiswa struct {
	service service.NilaiSiswa
}

func NewNilaiSiswa(service service.NilaiSiswa) NilaiSiswa {
	return &nilaiSiswa{
		service: service,
	}
}

func (n *nilaiSiswa) PostForm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mapel, semester := n.service.GetData()
	tmp := helper.View("view/admin/kelas/nilai.gohtml")
	err := tmp.Exec(w, r, "input_nilai", map[string]interface{}{
		"mapel":    mapel,
		"semester": semester,
		"userId": params["userId"],
	})
	helper.PanicIfError(err)
}

func (n *nilaiSiswa) Post(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	grade,_ := strconv.ParseFloat(r.PostFormValue("nilai"),64)
	request := request2.Nilai{
		StudentId:  params["userId"],
		TeacherId:  r.PostFormValue("teacherId"),
		SemesterId: r.PostFormValue("semester"),
		SubjectId:  r.PostFormValue("mapel"),
		Grade:      grade,
	}
	result := n.service.Post(request)
	helper.Notif("Notification",result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}

func (n *nilaiSiswa) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gradeId := params["gradeId"]
	result := n.service.Delete(gradeId)
	helper.Notif("Notification",result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}

func (n *nilaiSiswa) UpdateForm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gradeId := params["gradeId"]
	grade,semester,mapel := n.service.GetNilaiById(gradeId)
	tmp := helper.View("view/admin/kelas/update_nilai.gohtml")
	err := tmp.Exec(w,r,"update_nilai",map[string]interface{} {
		"grade": grade,
		"semester": semester,
		"mapel": mapel,
	})
	helper.PanicIfError(err)
}

func (n *nilaiSiswa) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gradeId := params["gradeId"]
	grade,_ := strconv.ParseFloat(r.PostFormValue("nilai"),64)
	request := request2.Nilai{
		SemesterId: r.PostFormValue("semester"),
		SubjectId:  r.PostFormValue("mapel"),
		Grade:      grade,
	}
	result := n.service.Update(gradeId,request)
	helper.Notif("Notification",result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}
