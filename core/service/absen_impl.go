package service

import (
	request2 "github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
	"strings"
)

type absen struct{
	db *gorm.DB
	repo repository.Absen
}

func NewAbsen(db *gorm.DB, repo repository.Absen) Absen {
	return &absen{
		db:   db,
		repo: repo,
	}
}

func (a *absen) GetStudentAbsent(siswaId string) []response.Absen {
	result,err := a.repo.GetStudentAbsent(a.db,siswaId)
	helper.PanicIfError(err)
	return result
}

func (a *absen) Absen(subjectId string,absen []string) string {
	err := a.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range absen {
			items := strings.Split(v,"_")
			request := request2.Absen{
				Id:        items[0],
				SubjectId: subjectId,
				Key:       items[1],
			}
			err := a.repo.Absen(tx,request)
			if err != nil {
				return err
			}
		}
		return nil
	})
	helper.PanicIfError(err)
	return "absen telah dilakukan"
}
