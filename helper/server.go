package helper

import (
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer(port string,r *mux.Router) {
	SetStaticFile(r)
	server := http.Server{
		Addr: port,
		Handler: r,
	}
	err := server.ListenAndServe()
	PanicIfError(err)
}
