package middleware

import (
	"github.com/rizface/sekolah/helper"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		s,_ := helper.Store.Get(request,"user-data")
		level := s.Values["level"];
		if level == nil {
			http.Redirect(writer,request,"/",http.StatusSeeOther)
		}
		next.ServeHTTP(writer,request)
	})
}
