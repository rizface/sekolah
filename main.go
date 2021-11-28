package main

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/setup"
	"github.com/rizface/sekolah/helper"
)

func init() {
	db := helper.Connection()
	err := db.AutoMigrate(&app.Level{},&app.User{})
	helper.PanicIfError(err)
}

func main() {
	setup.Login()
	setup.Admin()
	helper.StartServer(":8080",setup.R)
}
