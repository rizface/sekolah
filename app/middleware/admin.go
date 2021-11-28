package middleware

import (
	"github.com/rizface/sekolah/helper"
	"net/http"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		s,_ := helper.Store.Get(request,"user-data")
		if s.Values["level"] != "admin" {
			http.Redirect(writer,request,request.Referer(),http.StatusSeeOther)
		} else {
			next.ServeHTTP(writer,request)
		}
	})
}
