package controller

import "net/http"

type NilaiSiswa interface {
	PostForm(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	UpdateForm(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
