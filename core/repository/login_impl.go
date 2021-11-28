package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type login struct{}

func NewLogin() Login{
	return login{}
}

func (l login) FindByUsername(username string, db *gorm.DB) (response.Login,error) {
	var user response.Login
	result := db.Joins("INNER JOIN levels ON levels.id = users.level_id").Select("users.id,users.nama_depan,users.nama_belakang,levels.level,users.username,users.password").Where("username = ?", username).First(&app.User{}).Scan(&user)
	return user,result.Error
}

