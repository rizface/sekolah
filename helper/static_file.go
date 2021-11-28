package helper

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetStaticFile(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
}