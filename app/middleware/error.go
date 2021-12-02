package middleware

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rizface/sekolah/app/exception"
	"github.com/rizface/sekolah/helper"
	logger "github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strings"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				error := err.(error)
				logger.WithError(error)
				if strings.Contains(error.Error(),"crypto/bcrypt") {
					helper.Notif("Error", "Username / Password Salah")
				} else{
					errField,ok := error.(validator.ValidationErrors)
					if len(errField) > 0 && ok {
						msg := fmt.Sprintf("%s %s", errField[0].Field(),exception.Exception[errField[0].Tag()])
						helper.Notif("Error", msg)
					} else {
						helper.Notif("Error",error.Error())
					}
				}
				http.Redirect(writer,request,request.Referer(),http.StatusSeeOther)
			}
		}()
		next.ServeHTTP(writer,request)
	})
}
