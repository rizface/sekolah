package setup

import (
	"github.com/gorilla/mux"
	"github.com/rizface/sekolah/app/middleware"
)

var R = mux.NewRouter()

func init() {
	R.Use(middleware.ErrorHandler)
	R.Use(middleware.Extract)
}

