package controller

import "net/http"

type Kelas interface {
	GetStudents(w http.ResponseWriter, r *http.Request)
	DetailStudents(w http.ResponseWriter, r *http.Request)
	AddStudents(w http.ResponseWriter, r *http.Request)
	DeleteStudents(W http.ResponseWriter, r *http.Request)
}
