package middleware

import (
	"fmt"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		s,_ := helper.Store.Get(request,"user-data")
		level := s.Values["level"]
		if  level != "admin" {
			http.Redirect(writer,request,fmt.Sprintf("/%s",level),http.StatusSeeOther)
		} else {
			next.ServeHTTP(writer,request)
		}
	})
}
