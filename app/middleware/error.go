package middleware

import (
	"github.com/rizface/sekolah/helper"
	"net/http"
	"strings"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				error := err.(error)
				if strings.Contains(error.Error(),"crypto/bcrypt") {
					helper.Notif("Error", "Username / Password Salah")
				} else {
					helper.Notif("Error", error.Error())
				}
				http.Redirect(writer,request,request.Referer(),http.StatusSeeOther)
			}
		}()
		next.ServeHTTP(writer,request)
	})
}
