package controller

import "net/http"

type Mapel interface {
	GetPengampu(w http.ResponseWriter, r *http.Request)
}
