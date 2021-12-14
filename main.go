package main

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/setup"
	"github.com/rizface/sekolah/helper"
)

func init() {
	db := helper.Connection()
	err := db.AutoMigrate(&app.Level{},&app.User{},&app.Class{},&app.UserClass{},&app.TeacherClass{},&app.Subject{},&app.TeacherSubject{},&app.Semester{},&app.Grade{},&app.DetailStudent{},&app.DetailEmployee{},&app.MonthlyFee{})
	helper.PanicIfError(err)
}

func main() {
	setup.Login()
	setup.Admin()
	setup.Guru()
	helper.StartServer(":8080",setup.R)
}
