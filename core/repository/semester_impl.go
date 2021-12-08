package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type semester struct{}

func NewSemester() Semester {
	return semester{}
}

func (s semester) Get(db *gorm.DB) ([]response.Semester,error) {
	var semesters []response.Semester
	result :=  db.Find(&[]app.Semester{}).Scan(&semesters)
	return semesters,result.Error
}

