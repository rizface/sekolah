package controller

import "net/http"

type DetailPegawai interface {
	Post(w http.ResponseWriter, r *http.Request)
}
