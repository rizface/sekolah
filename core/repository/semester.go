package repository

import (
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type Semester interface {
	Get(db *gorm.DB) ([]response.Semester,error)
}
