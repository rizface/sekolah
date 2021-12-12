package service

import (
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
)

type semester struct {
	db *gorm.DB
	repo repository.Semester
}

func NewSemester(db *gorm.DB,repo repository.Semester) Semester{
	return &semester{db: db,repo: repo}
}

func (s *semester) GetSemester() []response.Semester {
	semester,err := s.repo.Get(s.db)
	helper.PanicIfError(err)
	return semester
}

