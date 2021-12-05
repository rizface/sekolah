package controller

import "net/http"

type WaliKelas interface {
	GetWaliKelas(w http.ResponseWriter, r *http.Request)
	PostWaliKelas(w http.ResponseWriter, r *http.Request)
	DeleteWaliKelas(w http.ResponseWriter, r *http.Request)
}
