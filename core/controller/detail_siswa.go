package controller

import "net/http"

type DetailSiswa interface {
	Update(w http.ResponseWriter, r *http.Request)
}
