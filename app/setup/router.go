package setup

import (
	"github.com/rizface/sekolah/app/middleware"
	"github.com/rizface/sekolah/route"
	"net/http"
)

func Login() {
	r := R.NewRoute().Subrouter()
	controller := LoginController()

	r.HandleFunc(route.LOGIN, controller.LoginPage).Methods(http.MethodGet)
	r.HandleFunc(route.LOGIN, controller.Login).Methods(http.MethodPost)
}

func Admin() {
	r := R.NewRoute().Subrouter()
	r.Use(middleware.Admin)

	// CRUD GURU
	crudAdmin := CrudAdminController()
	r.HandleFunc(route.ADMIN_DASHBOARD, crudAdmin.Get).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_ADMIN,crudAdmin.PostPage).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_ADMIN, crudAdmin.Post).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_ADMIN,crudAdmin.Delete).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_ADMIN,crudAdmin.UpdatePage).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_ADMIN,crudAdmin.Update).Methods(http.MethodPost)

	// CRUD GURU
	//crudGuru := CrudGuruController()
	r.HandleFunc(route.GURU_DASHBOARD, crudAdmin.Get).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_GURU,crudAdmin.PostPage).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_GURU,crudAdmin.Post).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_GURU,crudAdmin.Delete).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_GURU,crudAdmin.UpdatePage).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_GURU,crudAdmin.Update).Methods(http.MethodPost)
}