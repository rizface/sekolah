package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	request2 "github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
	"strconv"
)

type guruMapel struct{
	service service.GuruMapel
	kelasService service.Kelas
	mapelService service.Mapel
	semesterService service.Semester
	nilaiService service.NilaiSiswa
	userService service.AdminCrud
	absenService service.Absen
}

func NewGuruMapel(service service.GuruMapel, kelasService service.Kelas, mapelService service.Mapel, semesterService service.Semester,nilaiService service.NilaiSiswa, userService service .AdminCrud, absenService service.Absen) Guru {
	return &guruMapel{
		service: service,
		kelasService: kelasService,
		mapelService: mapelService,
		semesterService: semesterService,
		nilaiService: nilaiService,
		userService: userService,
		absenService: absenService,
	}
}

func (g *guruMapel) GetMapel(w http.ResponseWriter, r *http.Request) {
	fmt.Println("masuk")
	s,_ := helper.Store.Get(r,"user-data")
	subjects := g.service.GetMapel(s.Values["user_id"])
	tmp := helper.ViewGuru("view/guru/mapel.gohtml")
	err := tmp.Exec(w,r,"data_mapel_guru", map[string]interface{} {
		"mapel":subjects,
	})
	helper.PanicIfError(err)
}

func (g *guruMapel) GetKelas(w http.ResponseWriter, r *http.Request) {
	kelas := g.kelasService.Get()
	tmp := helper.ViewGuru("view/guru/kelas.gohtml")
	err := tmp.Exec(w,r,"data_guru_kelas", map[string]interface{}{
		"classes":kelas,
	})
	helper.PanicIfError(err)
}

func (g *guruMapel) GetSiswaKelas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kelasId := params["kelasId"]
	students,_ := g.kelasService.GetStudent(kelasId)
	fmt.Println("masil")
	tmp := helper.ViewGuru("view/guru/siswa_kelas.gohtml")
	err := tmp.Exec(w,r,"data_guru_siswa_kelas", map[string]interface{} {
		"students": students,
		"kelasId": kelasId,
	})
	helper.PanicIfError(err)
}

func (g *guruMapel) PostGradePage(w http.ResponseWriter, r *http.Request) {
	s,_ := helper.Store.Get(r,"user-data")
	params := mux.Vars(r)
	mapel := g.mapelService.Get()
	semester := g.semesterService.GetSemester()
	tmp := helper.ViewGuru("view/guru/nilai.gohtml")
	err := tmp.Exec(w,r,"input_nilai",map[string]interface{} {
		"mapel": mapel,
		"semester": semester,
		"siswaId": params["siswaId"],
		"guruId": s.Values["user_id"],
		"kelasId": params["kelasId"],
	})
	helper.PanicIfError(err)
}

func (g *guruMapel) PostGrade(w http.ResponseWriter, r *http.Request) {
	s,_ := strconv.ParseFloat(r.PostFormValue("nilai"),64)
	request := request2.Nilai{
		StudentId:  r.PostFormValue("siswaId"),
		TeacherId:  r.PostFormValue("guruId"),
		SemesterId: r.PostFormValue("semester"),
		SubjectId:  r.PostFormValue("mapel"),
		Grade:      s,
	}
	result := g.nilaiService.Post(request)
	helper.Notif("Notification", result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}

func (g *guruMapel) GetNilai(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["siswaId"]
	semester := r.URL.Query().Get("semester")
	subject := r.URL.Query().Get("subject")

	grade := g.nilaiService.GetNilai(userId, semester, subject)
	semesters:= g.semesterService.GetSemester()
	subjects := g.mapelService.Get()

	tmp := helper.ViewGuru("view/guru/detail_nilai.gohtml")
	err := tmp.Exec(w, r, "detail_nilai", map[string]interface{}{
		"grades": grade,
		"semesters": semesters,
		"subjects": subjects,
	})
	helper.PanicIfError(err)
}

func (g *guruMapel) HapusNilai(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := g.nilaiService.Delete(params["nilaiId"])
	helper.Notif("Notification", result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}

func (g *guruMapel) UpdateNilaiPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	grade,semesters,subjects := g.nilaiService.GetNilaiById(params["nilaiId"])
	tmp := helper.ViewGuru("view/guru/update_nilai.gohtml")
	err := tmp.Exec(w, r, "update_nilai", map[string]interface{}{
		"grades": grade,
		"semesters": semesters,
		"subjects": subjects,
		"nilaiId": params["nilaiId"],
	})
	helper.PanicIfError(err)
}

func (g *guruMapel) UpdateNilai(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	s,_ := strconv.ParseFloat(r.PostFormValue("nilai"),64)
	request := request2.Nilai{
		SemesterId: r.PostFormValue("semester"),
		SubjectId:  r.PostFormValue("mapel"),
		Grade:      s,
	}
	result := g.nilaiService.Update(params["nilaiId"],request)
	helper.Notif("Notification", result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}

func (g *guruMapel) GetDetailSiswa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	siswaId := params["siswaId"]
	detail := g.userService.DetailById("murid",siswaId);
	general := g.userService.GetById(siswaId);
	absen := g.absenService.GetStudentAbsent(siswaId)
	fmt.Println(absen)
	tmp := helper.ViewGuru("view/guru/detail_siswa.gohtml")
	err := tmp.Exec(w,r,"detail_siswa",map[string]interface{} {
		"general": general,
		"detail": detail,
		"absen": absen,
	})
	helper.PanicIfError(err)
}

func (g *guruMapel) AbsenPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kelasId := params["kelasId"]
	students,_ := g.kelasService.GetStudent(kelasId)
	subjects := g.mapelService.Get()
	tmp := helper.ViewGuru("view/guru/absen.gohtml")
	err := tmp.Exec(w,r,"absen",map[string]interface{} {
		"students": students,
		"subjects": subjects,
		"kelasId": kelasId,
	})
	helper.PanicIfError(err)
}

func (g *guruMapel) Absen(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	helper.PanicIfError(err)
	result := g.absenService.Absen(r.PostFormValue("subjectId"),r.PostForm["absen"])
	helper.Notif("Notification",result)
	http.Redirect(w,r,r.Referer(),http.StatusSeeOther)
}
