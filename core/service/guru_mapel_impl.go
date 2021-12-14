package service

import (
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
)

type guruMapel struct {
	db *gorm.DB
	repo repository.GuruMapel
}

func NewGuruMapel(db *gorm.DB, repo repository.GuruMapel) GuruMapel {
	return &guruMapel{
		db:   db,
		repo: repo,
	}
}

func (g *guruMapel) GetMapel(guruId interface{}) []response.Mapel {
	result,err := g.repo.GetMapel(g.db,guruId)
	helper.PanicIfError(err)
	return result
}
