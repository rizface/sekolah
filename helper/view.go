package helper

import "text/template"

var tmp,_  = template.ParseGlob("view/wrapper/*.gohtml")

func View(viewPath string) *template.Template {
	tmp.ParseGlob("view/nav/*.gohtml")
	tmp.ParseGlob("view/partials/*.gohtml")
	tmp,err := tmp.ParseFiles(viewPath)
	PanicIfError(err)
	return tmp
}