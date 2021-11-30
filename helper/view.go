package helper

import (
	"net/http"
	"text/template"
)

var tmp,_  = template.ParseGlob("view/wrapper/*.gohtml")

type ExecView struct {
	tmp *template.Template
}

func(e *ExecView) Exec(w http.ResponseWriter,r *http.Request,name string, data map[string]interface{}) error {
	s,_ := Store.Get(r,"user-data")
	if data == nil {
		data = make(map[string]interface{})
	}
	data["user"] = s.Values
	err := e.tmp.ExecuteTemplate(w,name,data)
	return err
}

func View(viewPath string)  ExecView {
	var admin ExecView
	tmp.ParseGlob("view/nav/*.gohtml")
	tmp.ParseGlob("view/partials/*.gohtml")
	tmp,err := tmp.ParseFiles(viewPath)
	PanicIfError(err)
	admin.tmp = tmp
	return admin
}