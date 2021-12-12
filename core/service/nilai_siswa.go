package service

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
)

type NilaiSiswa interface {
	GetData() ([]response.Mapel,[]response.Semester)
	GetNilai(userId string,semester string, subject string) []response.Nilai
	GetNilaiById(gradeId string) (response.Nilai,[]response.Semester,[]response.Mapel)
	Post(request request.Nilai) string
	Update(gradeId string, request request.Nilai) string
	Delete(gradeId string) string
}
