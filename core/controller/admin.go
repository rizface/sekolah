package controller

import "net/http"

type AdminCrud interface {
	Get(w http.ResponseWriter, r *http.Request)
	PostPage(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	UpdatePage(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

//type AdminCrudGuru interface {
//	Get(w http.ResponseWriter, r *http.Request)
//	PostPage(w http.ResponseWriter, r *http.Request)
//	Post(w http.ResponseWriter, r *http.Request)
//	UpdatePage(w http.ResponseWriter, r *http.Request)
//	Update(w http.ResponseWriter, r *http.Request)
//	Delete(w http.ResponseWriter, r *http.Request)
//}
