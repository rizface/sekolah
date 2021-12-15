package service

import "github.com/rizface/sekolah/app/model/response"

type Absen interface{
	GetStudentAbsent(siswaId string) []response.Absen
	Absen(subjectId string, absen []string) string
}
