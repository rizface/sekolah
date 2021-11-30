package controller

import (
	"fmt"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"html/template"
	"net/http"
)

type login struct {
	service service.Login
}

func NewLogin(service service.Login) Login {
	return login{service:service}
}

func (l login) LoginPage(w http.ResponseWriter, r *http.Request) {
	tmp,err := template.ParseGlob("view/partials/*.gohtml")
	helper.PanicIfError(err)
	tmp,err = tmp.ParseFiles("view/auth/login.gohtml")
	helper.PanicIfError(err)
	err = tmp.ExecuteTemplate(w,"login",nil)
	helper.PanicIfError(err)
}

func (l login) Login(w http.ResponseWriter, r *http.Request) {
 	loginPayload := request.Login{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}
	fmt.Println(loginPayload)
	user := l.service.Login(loginPayload)
	helper.CreateSession(w,r,user)
	http.Redirect(w,r,"/"+user.Level,http.StatusSeeOther)
}

