package controller

import "net/http"

type Guru interface {
	GetMapel(w http.ResponseWriter,r *http.Request)
	GetKelas(w http.ResponseWriter, r *http.Request)
	GetSiswaKelas(w http.ResponseWriter, r *http.Request)
	PostGradePage(w http.ResponseWriter, r *http.Request)
	PostGrade(w http.ResponseWriter, r *http.Request)
	GetNilai(w http.ResponseWriter,r *http.Request)
	HapusNilai(w http.ResponseWriter, r *http.Request)
	UpdateNilaiPage(w http.ResponseWriter, r *http.Request)
	UpdateNilai(w http.ResponseWriter, r *http.Request)
}

