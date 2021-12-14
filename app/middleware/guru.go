package middleware

import (
	"github.com/rizface/sekolah/helper"
	"net/http"
)

func Guru(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		s,_ := helper.Store.Get(request,"user-data")
		level := s.Values["level"]
		if level == "guru" {
			next.ServeHTTP(writer,request)
		} else {
			http.Redirect(writer,request,request.Referer(),http.StatusSeeOther)
		}
	})
}
