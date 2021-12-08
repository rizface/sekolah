package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
)

type nilaiSiswa struct {
	db           *gorm.DB
	valid        *validator.Validate
	semesterRepo repository.Semester
	subjectRepo  repository.Mapel
	repo         repository.NilaiSiswa
}

func NewNilaiSiswa(db *gorm.DB, valid *validator.Validate, semesterRepo repository.Semester, subjectRepo repository.Mapel, repo repository.NilaiSiswa) NilaiSiswa {
	return &nilaiSiswa{
		db:           db,
		valid:        valid,
		semesterRepo: semesterRepo,
		subjectRepo:  subjectRepo,
		repo:         repo,
	}
}

func (n *nilaiSiswa) Post(request request.Nilai) string {
	err := n.valid.Struct(request)
	helper.PanicIfError(err)
	success,err := n.repo.Post(n.db,request)
	helper.PanicIfError(err)
	if success {
		return "nilai berhasil diinput"
	}
	return "nilai gagal diinput"
}

func (n *nilaiSiswa) GetData() ([]response.Mapel, []response.Semester) {
	semester, err := n.semesterRepo.Get(n.db)
	helper.PanicIfError(err)
	mapel, err := n.subjectRepo.Get(helper.Connection())
	helper.PanicIfError(err)
	return mapel, semester
}

func (n *nilaiSiswa) GetNilaiBySemester(userId string,semester string) []response.Nilai {
	var result []response.Nilai
	var err error
	if len(semester) == 0 {
		result,err = n.repo.GetAllGrade(n.db,userId)
		helper.PanicIfError(err)
	} else {
		fmt.Println(semester, "ini lho")
		result,err = n.repo.GetGradeBySemester(n.db,userId,semester)
		helper.PanicIfError(err)
	}
	return result
}

func (n *nilaiSiswa) GetNilaiById(gradeId string) (response.Nilai, []response.Semester, []response.Mapel) {
	mapel,semester := n.GetData()
	grade,err := n.repo.GetGradeById(n.db,gradeId)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			panic(errors.New("data tidak ditemukan"))
		}
		panic(errors.New("proses gagal"))
	}
	return grade,semester,mapel
}

func (n *nilaiSiswa) Delete(gradeId string) string {
	success,err := n.repo.Delete(n.db,gradeId)
	helper.PanicIfError(err)
	if success {
		return "nilai berhasil dihapus"
	}
	return "nilai gagal dihapus"
}

func (n *nilaiSiswa) Update(gradeId string, request request.Nilai) string {
	success,err := n.repo.Update(n.db,gradeId,request)
	helper.PanicIfError(err)
	if success {
		return "nilai berhasil diupdate"
	}
	return "nilai gagal diupdate"
}
