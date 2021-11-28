package repository

import (
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type Login interface {
	FindByUsername(username string, db *gorm.DB) (response.Login,error)
}
